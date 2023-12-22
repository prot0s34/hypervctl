package winrmclient

import (
	"context"
	"github.com/masterzen/winrm"
	"hypervctl/config"
)

func InitializeClient(cfg config.Config) (*winrm.Client, context.Context, context.CancelFunc, error) {
	endpoint := winrm.NewEndpoint(cfg.Hypervisor.Host, cfg.Hypervisor.Port, false, true, nil, nil, nil, 0)
	client, err := winrm.NewClient(endpoint, cfg.Hypervisor.Auth.Username, cfg.Hypervisor.Auth.Password)
	if err != nil {
		return nil, nil, nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	return client, ctx, cancel, nil
}

