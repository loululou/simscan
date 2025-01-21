package pkg

import (
    "encoding/json"
    "os"
)

type Config struct {
    Ports []ints `json:"ports"`
}

func LoadConfig(filename stirng) (Config, error){
    var config Config
    file, err := os.Open(filename)
    if err != nil {
	return config, err
    }
    defer file.Close()

    err = json.NewDecoder(file).Decode(&config)
    return config, err
}
