package describe

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"hypervctl/config"
	"hypervctl/winrmclient"
)

var vmCmd = &cobra.Command{
	Use:   "vm [vmname]",
	Short: "Describe a VM",
	Long:  `Describe detailed information about a specific VM.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vmName := args[0]

		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Println(err)
			return
		}

		client, ctx, cancel, err := winrmclient.InitializeClient(cfg)
		if err != nil {
			fmt.Println("Error creating WinRM client:", err)
			return
		}
		defer cancel()

		var stdout, stderr bytes.Buffer
		command := fmt.Sprintf("powershell -Command \"Get-VM -Name %s | Format-List \"", vmName)
		_, err = client.RunWithContext(ctx, command, &stdout, &stderr)
		if err != nil {
			fmt.Println("Error running command:", err)
			fmt.Println("STDERR:", stderr.String())
			return
		}

		fmt.Println(stdout.String())

	},
}
