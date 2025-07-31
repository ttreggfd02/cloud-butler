// cmd/scan.go

package cmd

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	// 匯入我們自己寫的套件
	"github.com/ttreggfd02/cloud-butler/internal/scanner"
	"github.com/ttreggfd02/cloud-butler/pkg/aws_session"

	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "掃描雲端資源以發現問題",
	Long:  `Scan 指令會根據設定對指定的雲端平台執行一系列的檢查，例如尋找閒置資源、不安全的設定等。`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("--- 初始化 AWS 連線 ---")
		// 建立 AWS 連線設定
		awsCfg, err := aws_session.NewConfig(context.TODO())
		if err != nil {
			fmt.Fprintf(os.Stderr, "建立 AWS session 失敗: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("AWS 連線成功！")

		fmt.Println("\n--- 開始執行 EBS 未掛載磁碟區掃描 ---")
		// 呼叫我們的掃描器
		unattachedVolumes, err := scanner.ScanUnattachedVolumes(context.TODO(), awsCfg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "掃描未掛載磁碟區失敗: %v\n", err)
			os.Exit(1)
		}

		// 使用 tabwriter 來格式化輸出，讓它像個表格
		if len(unattachedVolumes) > 0 {
			fmt.Println("\n掃描結果：發現以下未掛載的 EBS 磁碟區！")
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.FilterHTML)
			fmt.Fprintln(w, "Volume ID\tSize (GiB)\tCreated")
			fmt.Fprintln(w, "----------\t----------\t----------")
			for _, vol := range unattachedVolumes {
				fmt.Fprintf(w, "%s\t%d\t%s\n", vol.VolumeID, vol.Size, vol.Created)
			}
			w.Flush()
		} else {
			fmt.Println("\n掃描結果：太棒了！沒有發現任何未掛載的 EBS 磁碟區。")
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
