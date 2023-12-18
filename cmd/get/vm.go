package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "Get a list of VMs",
	Long:  `Get a detailed list of Hyper-V VMs.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Displaying list of VMs...")
	},
}
