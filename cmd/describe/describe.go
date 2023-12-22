package describe

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe choosen resources",
	Long:  `Get definition of the resources like vm, switch, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Specify what resource to describe (e.g., 'vm', 'switch')")
	},
}

func init() {
	Cmd.AddCommand(vmCmd)
}
