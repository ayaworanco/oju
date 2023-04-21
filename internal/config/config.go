package config

import (
	"errors"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	AllowedApplications []Application `yaml:"allowed_applications"`
}

type Application struct {
	Name, AppKey string
}

func GetAllowedApplications() []Application {
	config_file, load_config_file_error := load_config_file()
	if load_config_file_error != nil {
		log.Fatalln(load_config_file_error.Error())
	}

	config, load_error := build_config(config_file)
	if load_error != nil {
		log.Fatalln(load_error.Error())
	}

	return config.AllowedApplications
}

func build_config(config_file []byte) (Config, error) {
	var config Config

	err := yaml.Unmarshal(config_file, &config)
	if err != nil {
		println(err)
		return Config{}, err
	}

	return config, nil
}

func load_config_file() ([]byte, error) {
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
