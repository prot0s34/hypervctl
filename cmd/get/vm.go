package get

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"hypervctl/config"
)

var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "Get a list of VMs",
	Long:  `Get a detailed list of Hyper-V VMs.`,
	Run: func(cmd *cobra.Command, args []string) {
		var cfg config.Config
		if err := viper.Unmarshal(&cfg); err != nil {
			fmt.Println("Unable to decode into struct:", err)
			return
		}

		fmt.Printf("Connecting to %s:%d...\n", cfg.Hypervisor.Host, cfg.Hypervisor.Port)

	},
}
