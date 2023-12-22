package get

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"hypervctl/config"
	"hypervctl/winrmclient"
)

var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "Get a list of VMs",
	Long:  `Get a detailed list of Hyper-V VMs.`,
	Run: func(cmd *cobra.Command, args []string) {
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

		// string-by-string copying output of powershell looks little ugly.
		// Strongly needed to redesign that part, but call to winrm for hyperv communitation broken by design in that case. %)
		var stdout, stderr bytes.Buffer
		command := "powershell -Command \"Get-VM | Format-Table -Property Name, State, Status -AutoSize | Out-String -Width 4096\""
		_, err = client.RunWithContext(ctx, command, &stdout, &stderr)
		if err != nil {
			fmt.Println("Error running command:", err)
			fmt.Println("STDERR:", stderr.String())
			return
		}

		fmt.Println(stdout.String())
	},
}
