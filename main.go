package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"hypervctl/cmd/get"
)

var rootCmd = &cobra.Command{
	Use:   "hypervctl",
	Short: "hypervctl is a CLI tool for managing Hyper-V VMs",
	Long:  `hypervctl is a CLI tool designed to manage Hyper-V VMs, similar to how kubectl manages Kubernetes.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(get.Cmd)
}

func main() {
	Execute()
}
