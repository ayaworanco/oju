package config

import (
	"fmt"
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	AllowedApplications []Application `json:"allowed_applications"`
}

type Application struct {
	Name   string `json:"name"`
	AppKey string `json:"app_key"`
	Host   string `json:"host"`
}

func BuildConfig(config_file []byte) (Config, error) {
	var config Config

	err := json.Unmarshal(config_file, &config)
	if err != nil {
		return Config{}, err
	}

	for _, application := range config.AllowedApplications {
		fmt.Println("=> application loaded!")
		fmt.Println("[name]: ", application.Name)
		fmt.Println("[key]: ", application.AppKey)
		fmt.Println("[host]: ", application.Host)
		fmt.Println("--------------------------------------------")
	}

	return config, nil
}

func LoadConfigFile() ([]byte, error) {
	var config_file string
	config_file = os.Getenv("CONFIG_JSON_PATH")

	if config_file == "" {
		config_file = "./config.json"
	}

	file, read_error := os.ReadFile(config_file)
	if read_error != nil {
		return nil, errors.New(config_file + " file not found")
	}

	return file, nil
}
