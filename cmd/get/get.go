package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "get",
	Short: "Get resources",
	Long:  `Get one of the resources like vm, vswitch, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Specify what resource to get (e.g., 'vm', 'switch')")
	},
}

func init() {
	Cmd.AddCommand(vmCmd)
	Cmd.AddCommand(switchCmd)
}
