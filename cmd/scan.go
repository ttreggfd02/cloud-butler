// cmd/scan.go (完整替換)

package cmd

import (
	"context"
	"fmt"
	"os"
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
		// --- 根據模式顯示警告 ---
		if executeMode {
			fmt.Println("######################################################")
			fmt.Println("###               -- EXECUTE MODE --               ###")
			fmt.Println("###  接下來將執行真實的刪除動作，請謹慎操作！  ###")
			fmt.Println("######################################################")
		} else {
			fmt.Println("--- Dry Run Mode --- (僅報告，不執行任何動作)")
		}

		// --- 初始化 ---
		fmt.Println("\n--- 初始化 AWS 連線 ---")
		awsCfg, err := aws_session.NewConfig(context.TODO())
		if err != nil {
			fmt.Fprintf(os.Stderr, "建立 AWS session 失敗: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("AWS 連線成功！")

		// --- 執行 EBS 未掛載磁碟區掃描 ---
		fmt.Println("\n--- 開始執行 EBS 未掛載磁碟區掃描 ---")
		unattachedVolumes, err := scanner.ScanUnattachedVolumes(context.TODO(), awsCfg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "掃描未掛載磁碟區失敗: %v\n", err)
			os.Exit(1)
		}
		if len(unattachedVolumes) > 0 {
			if executeMode {
				fmt.Println("\n[執行模式] 開始清理未掛載的磁碟區：")
				for _, vol := range unattachedVolumes {
					err := scanner.DeleteUnattachedVolume(context.TODO(), awsCfg, vol.VolumeID)
					if err != nil {
						fmt.Fprintf(os.Stderr, "==> [錯誤] 刪除 Volume ID %s 失敗: %v\n", vol.VolumeID, err)
					} else {
						fmt.Printf("==> [成功] 已刪除 Volume ID: %s (大小: %d GiB)\n", vol.VolumeID, vol.Size)
					}
				}
			} else {
				// Dry Run 模式下的報告
				fmt.Println("\n[預演模式] 掃描結果：發現以下未掛載的 EBS 磁碟區！")
				w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
				fmt.Fprintln(w, "Volume ID\tSize (GiB)\tCreated")
				fmt.Fprintln(w, "----------\t----------\t----------")
				for _, vol := range unattachedVolumes {
					fmt.Fprintf(w, "%s\t%d\t%s\n", vol.VolumeID, vol.Size, vol.Created)
				}
				w.Flush()
			}
		} else {
			fmt.Println("\n掃描結果：太棒了！沒有發現任何未掛載的 EBS 磁碟區。")
		}

		// --- 執行 EBS 過舊快照掃描 ---
		fmt.Println("\n--- 開始執行 EBS 過舊快照掃描 ---")
		minDaysOld := viper.GetInt("scanners.ebs.minDaysOldForSnapshots")
		fmt.Printf("將尋找建立時間超過 %d 天的快照...\n", minDaysOld)
		
		oldSnapshots, err := scanner.ScanOldSnapshots(context.TODO(), awsCfg, minDaysOld)
		if err != nil {
			fmt.Fprintf(os.Stderr, "掃描過舊快照失敗: %v\n", err)
			os.Exit(1)
		}
		if len(oldSnapshots) > 0 {
			if executeMode {
				fmt.Println("\n[執行模式] 開始清理過舊的快照：")
				for _, snap := range oldSnapshots {
					err := scanner.DeleteOldSnapshot(context.TODO(), awsCfg, snap.SnapshotID)
					if err != nil {
						fmt.Fprintf(os.Stderr, "==> [錯誤] 刪除 Snapshot ID %s 失敗: %v\n", snap.SnapshotID, err)
					} else {
						fmt.Printf("==> [成功] 已刪除 Snapshot ID: %s (來自 Volume: %s)\n", snap.SnapshotID, snap.VolumeID)
					}
				}
			} else {
				// Dry Run 模式下的報告
				fmt.Println("\n[預演模式] 掃描結果：發現以下過舊的 EBS 快照！")
				w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
				fmt.Fprintln(w, "Snapshot ID\tVolume ID\tSize (GiB)\tCreated")
				fmt.Fprintln(w, "------------\t----------\t----------\t----------")
				for _, snap := range oldSnapshots {
					fmt.Fprintf(w, "%s\t%s\t%d\t%s\n", snap.SnapshotID, snap.VolumeID, snap.VolumeSize, snap.StartTime)
				}
				w.Flush()
			}
		} else {
			fmt.Println("\n掃描結果：太棒了！沒有發現任何過舊的 EBS 快照。")
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().BoolVar(&executeMode, "execute", false, "執行實際的清理動作，而非僅顯示報告")
}
