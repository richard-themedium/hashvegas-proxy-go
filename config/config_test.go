package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	os.Setenv("CONNECTION_PROFILE_PATH", "testvalue")
	os.Setenv("MSP_CONFIG_PATH", "config value1")
	c := Config{}
	err := c.Load()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", c)
}
