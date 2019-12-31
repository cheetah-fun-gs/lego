package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// 参数
var (
	EnvName    string
	IsAutoInit bool // 是否自动init  非web服务不自动加载
	isHelp     bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "lego",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && args[0] == "help" {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&EnvName, "env", "e", "local", "configs环境")
	rootCmd.PersistentFlags().BoolVarP(&IsAutoInit, "auto-init", "a", true, "是否自动Init")
	rootCmd.PersistentFlags().BoolVarP(&isHelp, "help", "h", false, "帮助信息")

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}

	if isHelp {
		os.Exit(0)
	}
}
