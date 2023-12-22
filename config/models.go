package config

// add https true/false to Hypervisor.Auth sub-struct
type Config struct {
	Hypervisor struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
		Auth struct {
			Username string `mapstructure:"username"`
			Password string `mapstructure:"password"`
		} `mapstructure:"auth"`
	} `mapstructure:"hypervisor"`
}
