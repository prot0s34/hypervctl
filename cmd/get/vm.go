package get

import (
	"bytes"
	"context"
	"fmt"
	"hypervctl/config"

	"github.com/masterzen/winrm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "Get a list of VMs",
	Long:  `Get a detailed list of Hyper-V VMs.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := loadConfig()
		if err != nil {
			fmt.Println("Error loading config:", err)
			return
		}

		// move to separate place and use it more generic-way
		if cfg.Hypervisor.Host == "" || cfg.Hypervisor.Auth.Username == "" || cfg.Hypervisor.Auth.Password == "" {
			fmt.Println("Configuration error: Host, Username, or Password is missing.")
			return
		}

		// exluce that part to standalone function/pkg with all connect-init logic
		endpoint := winrm.NewEndpoint(cfg.Hypervisor.Host, cfg.Hypervisor.Port, false, true, nil, nil, nil, 0)
		client, err := winrm.NewClient(endpoint, cfg.Hypervisor.Auth.Username, cfg.Hypervisor.Auth.Password)
		if err != nil {
			fmt.Println("Error creating WinRM client:", err)
			return
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// string-by-string copying output of powershell looks little ugly. Need to redesign that part, but call winrm broken by design in that approach %)
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

// move loadconfig-related functional to separate file to avoid duplicating
func loadConfig() (config.Config, error) {
	var cfg config.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
