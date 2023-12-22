package main

import (
	"fmt"
	"os"

	"hypervctl/cmd/describe"
	"hypervctl/cmd/get"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "hypervctl",
	Short: "hypervctl is a CLI tool for managing Hyper-V VMs",
	Long:  `hypervctl is a CLI tool designed to manage Hyper-V VMs, similar to how kubectl manages Kubernetes.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initConfig()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hyperv/config)")
	rootCmd.AddCommand(get.Cmd)
	rootCmd.AddCommand(describe.Cmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		viper.SetConfigType("yaml")
	} else {
		viper.AddConfigPath("$HOME/.hyperv")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
