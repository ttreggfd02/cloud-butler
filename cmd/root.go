/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
        "fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cloud-butler",
	Short: "一個自動化雲端資源巡檢與清理的 Go 工具",
	Long: `Cloud-Butler 是一個 CLI 工具，旨在幫助 DevOps 工程師自動化雲端基礎設施的巡檢。
它可以發現閒置或不合規的資源，以節省成本並提升安全性。`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cloud-butler.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cobra.OnInitialize(initConfig)
}
func initConfig() {
	// 設定檔名稱為 config (不用副檔名)
	viper.SetConfigName("config")
	// 設定檔類型為 YAML
	viper.SetConfigType("yaml")
	// 在當前目錄下尋找設定檔
	viper.AddConfigPath(".")

	// 嘗試讀取設定檔
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 如果設定檔不存在，可以忽略錯誤，後續可用環境變數等
			fmt.Fprintln(os.Stderr, "Config file not found; using defaults or environment variables.")
		} else {
			// 如果設定檔被找到但解析出錯
			fmt.Fprintln(os.Stderr, "Error reading config file:", err)
			os.Exit(1)
		}
	}
}


