package scfgw

import (
	"os"

	"github.com/cheetah-fun-gs/lego/internal/common"
	"github.com/spf13/cobra"
)

// 参数
var (
	envName string
	isHelp  bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "lego",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && args[0] == "help" {
			cmd.Help()
			os.Exit(0)
		}
		common.Init(&common.Opt{
			EnvName:   envName,
			StartType: common.StartTypeScfgw,
		})
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&envName, "env", "e", "local", "configs环境")
	rootCmd.PersistentFlags().BoolVarP(&isHelp, "help", "h", false, "帮助信息")

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}

	if isHelp {
		os.Exit(0)
	}
}
