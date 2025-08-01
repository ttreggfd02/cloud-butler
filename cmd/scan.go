package cmd

import (
	"context"
	"fmt"
	"os"
	"sync" // <-- 匯入 sync 套件
	"text/tabwriter"

	"github.com/ttreggfd02/cloud-butler/internal/scanner"
	"github.com/ttreggfd02/cloud-butler/pkg/aws_session"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var executeMode bool

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "掃描雲端資源以發現問題",
	Long:  `Scan 指令會根據設定對指定的雲端平台執行一系列的檢查。`,
	Run: func(cmd *cobra.Command, args []string) {
		if executeMode {
			fmt.Println("######################################################")
			fmt.Println("###               -- EXECUTE MODE --               ###")
			fmt.Println("###  接下來將執行真實的刪除動作，請謹慎操作！  ###")
			fmt.Println("######################################################")
		} else {
			fmt.Println("--- Dry Run Mode --- (僅報告，不執行任何動作)")
		}

		fmt.Println("\n--- 初始化 AWS 連線 ---")
		awsCfg, err := aws_session.NewConfig(context.TODO())
		if err != nil {
			fmt.Fprintf(os.Stderr, "建立 AWS session 失敗: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("AWS 連線成功！")

		// --- 使用 WaitGroup 實現併發掃描 ---
		var wg sync.WaitGroup
		fmt.Println("\n--- 開始併發執行所有掃描任務 ---")

		// 任務一：掃描未掛載磁碟區
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("-> (開始) 掃描未掛載磁碟區...")
			unattachedVolumes, err := scanner.ScanUnattachedVolumes(context.TODO(), awsCfg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "[錯誤] 掃描未掛載磁碟區失敗: %v\n", err)
				return
			}
			// 處理掃描結果...
			if len(unattachedVolumes) > 0 {
				if executeMode {
					// ... execute 邏輯 ...
				} else {
					// ... dry-run 報告邏輯 ...
					fmt.Println("\n[報告] 發現未掛載的 EBS 磁碟區！")
					w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
					fmt.Fprintln(w, "Volume ID\tSize (GiB)\tCreated")
					fmt.Fprintln(w, "----------\t----------\t----------")
					for _, vol := range unattachedVolumes {
						fmt.Fprintf(w, "%s\t%d\t%s\n", vol.VolumeID, vol.Size, vol.Created)
					}
					w.Flush()
				}
			} else {
				fmt.Println("-> (完成) 未發現任何未掛載的 EBS 磁碟區。")
			}
		}()

		// 任務二：掃描過舊快照
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("-> (開始) 掃描過舊快照...")
			minDaysOld := viper.GetInt("scanners.ebs.minDaysOldForSnapshots")
			oldSnapshots, err := scanner.ScanOldSnapshots(context.TODO(), awsCfg, minDaysOld)
			if err != nil {
				fmt.Fprintf(os.Stderr, "[錯誤] 掃描過舊快照失敗: %v\n", err)
				return
			}
			// 處理掃描結果...
			if len(oldSnapshots) > 0 {
				if executeMode {
					// ... execute 邏輯 ...
				} else {
					// ... dry-run 報告邏輯 ...
					fmt.Println("\n[報告] 發現過舊的 EBS 快照！")
					w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
					fmt.Fprintln(w, "Snapshot ID\tVolume ID\tSize (GiB)\tCreated")
					fmt.Fprintln(w, "------------\t----------\t----------\t----------")
					for _, snap := range oldSnapshots {
						fmt.Fprintf(w, "%s\t%s\t%d\t%s\n", snap.SnapshotID, snap.VolumeID, snap.VolumeSize, snap.StartTime)
					}
					w.Flush()
				}
			} else {
				fmt.Println("-> (完成) 未發現任何過舊的 EBS 快照。")
			}
		}()

		// 任務三：掃描 S3 公開儲存桶
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("-> (開始) 掃描 S3 公開儲存桶...")
			publicBuckets, err := scanner.ScanPublicS3Buckets(context.TODO(), awsCfg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "[錯誤] 掃描 S3 公開儲存桶失敗: %v\n", err)
				return
			}
			// 處理掃描結果...
			if len(publicBuckets) > 0 {
				// ... execute 邏輯 (目前 S3 掃描器沒有 execute 動作) ...
				// ... dry-run 報告邏輯 ...
				fmt.Println("\n[報告] 發現有風險的 S3 儲存桶！")
				w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
				fmt.Fprintln(w, "Bucket Name\tRegion\tIssue")
				fmt.Fprintln(w, "------------\t----------\t----------")
				for _, bucket := range publicBuckets {
					fmt.Fprintf(w, "%s\t%s\t%s\n", bucket.BucketName, bucket.Region, bucket.Issue)
				}
				w.Flush()
			} else {
				fmt.Println("-> (完成) 未發現任何有公開存取風險的 S3 儲存桶。")
			}
		}()

		// 等待所有 Goroutine 完成
		wg.Wait()
		fmt.Println("\n--- 所有掃描任務均已執行完畢 ---")
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().BoolVar(&executeMode, "execute", false, "執行實際的清理動作，而非僅顯示報告")
}
