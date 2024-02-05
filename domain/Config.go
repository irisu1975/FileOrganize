package domain

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configs struct {
	Config []Config `json:"config"`
}

type Config struct {
	Extension []string	`json:"extension"`
	DirName		string		`json:"dir_name"`
}

// ReadConfig read json config file and return Configs.
func ReadConfig(path string) (Configs, error) {
	var jsonData Configs
	jsonFile, err := os.ReadFile(path)
	if err != nil {
		return Configs{}, fmt.Errorf("ReadConfig():Don't open JSON file: %s", path)
	}

	err = json.Unmarshal(jsonFile, &jsonData)
	if err != nil {
		return Configs{}, fmt.Errorf("ReadConfig():Please specify the correct JSON file.")
	}

	return jsonData, nil
}