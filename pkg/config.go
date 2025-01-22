package pkg

import (
	"encoding/json"
	"os"
)

type Config struct {
	Ports []int `json:"ports"`
}

func LoadConfig(filename string)(Config, error) {
	var config Config
	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)
	return config, err 
}