package config

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Resources []Resource `json:"resources"`
}

type Resource struct {
	Name string `json:"name"`
	Key  string `json:"key"`
	Host string `json:"host"`
}

func LoadConfigFile() (Config, error) {
	var config_file string

	path, path_err := os.Getwd()

	if path_err != nil {
		log.Fatal(path_err)
	}

	config_file = get_config_file_by_env(path)

	file, read_error := os.ReadFile(config_file)
	if read_error != nil {
		return Config{}, errors.New(config_file + ": file not found")
	}

	return build_config(file)
}

func get_config_file_by_env(path string) string {
	file := os.Getenv("CONFIG_JSON_PATH")
	if file == "" {
		return path + "/config.json"
	}
	return file
}

func build_config(file []byte) (Config, error) {
	var config Config

	unmarshal_error := json.Unmarshal(file, &config)

	if unmarshal_error != nil {
		return Config{}, unmarshal_error
	}

	if len(config.Resources) == 0 {
		return Config{}, errors.New("config needs at least 1 resource")
	}

	print_resources_loaded(config.Resources)

	return config, nil
}

func print_resources_loaded(resources []Resource) {
	if flag.Lookup("test.v") == nil {
		for _, resource := range resources {
			fmt.Println("=> resource loaded!")
			fmt.Println("[name]: ", resource.Name)
			fmt.Println("[key]: ", resource.Key)
			fmt.Println("[host]: ", resource.Host)
			fmt.Println("--------------------------------------------")
		}
	}
}
