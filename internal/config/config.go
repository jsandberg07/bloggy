package config

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() Config {

	fileName := getConfigFilePath()
	// fmt.Printf("// filename: %v\n", fileName)

	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading data from file: %v\n", err)
		os.Exit(1)
	}

	cfg := Config{}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Printf("Error unmarshaling: %v\n", err)
		os.Exit(1)
	}

	return cfg

}

func (cfg Config) SetUser(user string) {
	cfg.CurrentUserName = user
	write(cfg)
}

func write(cfg Config) {
	fileName := getConfigFilePath()
	data, err := json.Marshal(cfg)
	if err != nil {
		fmt.Printf("Error marshaling: %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile(fileName, data, fs.ModeDevice)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}
}

func getConfigFilePath() string {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		os.Exit(1)
	}
	return homeDir + "/.gatorconfig.json"
}

func Coco() string {
	return "coco"
}
