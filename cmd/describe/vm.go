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
	Args:  cobra.ExactArgs(1), // Expect exactly one argument: vmname
	Run: func(cmd *cobra.Command, args []string) {
		vmName := args[0] // The name of the VM to describe

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

		// Commands to run
		commands := []string{
			fmt.Sprintf("powershell -Command \"Get-VM -Name '%s' | Format-List *\"", vmName),
			fmt.Sprintf("powershell -Command \"Get-VMProcessor -VMName '%s' | Format-List *\"", vmName),
			fmt.Sprintf("powershell -Command \"Get-VMMemory -VMName '%s' | Format-List *\"", vmName),
			fmt.Sprintf("powershell -Command \"Get-VMCheckpoint -VMName '%s' | Format-List *\"", vmName),
			fmt.Sprintf("powershell -Command \"Get-VMNetworkAdapter -VMName '%s' | Format-List *\"", vmName),
			fmt.Sprintf("powershell -Command \"Get-VMHardDiskDrive -VMName '%s' | Format-List *\"", vmName),
		}

		var combinedOutput bytes.Buffer
		for _, command := range commands {
			var stdout, stderr bytes.Buffer
			_, err = client.RunWithContext(ctx, command, &stdout, &stderr)
			if err != nil {
				fmt.Println("Error running command:", err)
				fmt.Println("STDERR:", stderr.String())
				return
			}
			combinedOutput.WriteString(stdout.String())
			combinedOutput.WriteString("\n------------------------------------------------\n")
		}

		fmt.Println(combinedOutput.String())
	},
}

