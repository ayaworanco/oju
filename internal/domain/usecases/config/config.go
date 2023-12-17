package usecases

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"oju/internal/domain/entities"
	"os"
)

func BuildConfig(config_file []byte) (entities.Config, error) {
	var config entities.Config

	unmarshal_err := json.Unmarshal(config_file, &config)

	if unmarshal_err != nil {
		return entities.Config{}, unmarshal_err
	}

	if len(config.Resources) == 0 {
		return entities.Config{}, errors.New("malformed config file")
	}

	for _, application := range config.Resources {
		fmt.Println("=> application loaded!")
		fmt.Println("[name]: ", application.Name)
		fmt.Println("[key]: ", application.Key)
		fmt.Println("[host]: ", application.Host)
		fmt.Println("--------------------------------------------")
	}

	return config, nil
}

func LoadConfigFile() ([]byte, error) {
	var config_file string

	path, path_err := os.Getwd()

	if path_err != nil {
		log.Fatal(path_err)
	}

	config_file = os.Getenv("CONFIG_JSON_PATH")
	if config_file == "" {
		config_file = path + "/config.json"
	}

	file, read_error := os.ReadFile(config_file)
	if read_error != nil {
		return nil, errors.New(config_file + ": file not found")
	}

	return file, nil
}
