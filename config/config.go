package config

import (
	"fmt"
	// "github.com/Netflix/go-env"
	env "github.com/Netflix/go-env"
)

type Config struct {
	ConnectionProfilePath string `env:"CONNECTION_PROFILE_PATH"`
	MspConfigPath         string `env:"MSP_CONFIG_PATH"`
}

func (c *Config) Load() error {
	// os.Setenv("CONNECTION_PROFILE_PATH", "testvalue")
	es, err := env.UnmarshalFromEnviron(c)
	if err != nil {
		return err
	}
	fmt.Printf("Got the following: %+v\n", es)

	return nil
}
