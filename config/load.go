package config

import (
	"errors"
	"github.com/spf13/viper"
)

func LoadConfig() (Config, error) {
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}

	if err := validateConfig(cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func validateConfig(cfg Config) error {
	if cfg.Hypervisor.Host == "" {
		return errors.New("configuration error: Host is missing")
	}
	if cfg.Hypervisor.Auth.Username == "" {
		return errors.New("configuration error: Username is missing")
	}
	if cfg.Hypervisor.Auth.Password == "" {
		return errors.New("configuration error: Password is missing")
	}
	return nil
}

