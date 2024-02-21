package config

import (
	"os"
	"testing"
)

func TestReturnAllowedResources(t *testing.T) {
	os.Setenv("CONFIG_JSON_PATH", "config_testdata.json")
	config, load_error := LoadConfigFile()

	if load_error != nil {
		t.Error(load_error)
	}

	if len(config.Resources) == 0 {
		t.Error("should be at least 1 resource")
	}
}
