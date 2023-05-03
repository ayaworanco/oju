package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	AllowedApplications []Application `yaml:"allowed_applications"`
}

type Application struct {
	Name   string `yaml:"name"`
	AppKey string `yaml:"app_key"`
}

func BuildConfig(config_file []byte) (Config, error) {
	var config Config

	err := yaml.Unmarshal(config_file, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func LoadConfigFile() ([]byte, error) {
	var config_file string
	config_file = os.Getenv("CONFIG_YAML_PATH")

	if config_file == "" {
		config_file = "./config.yaml"
	}

	file, read_error := os.ReadFile(config_file)
	if read_error != nil {
		return nil, errors.New(config_file + " file not found")
	}

	return file, nil
}
