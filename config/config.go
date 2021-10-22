package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Server server `json:"server"`
	App    app    `json:"app"`
}

type server struct {
	Port int `json:"port"`
}

type app struct {
	Path      string `json:"path"`
	SyncInMin int    `json:"syncInMin"`
}

func (receiver *Config) Init() error {
	filename, err := filepath.Abs("./config/config.json")
	if err != nil {
		return fmt.Errorf("error while finding config path \n\t%v", err)
	}
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error while reading config file \n\t%v", err)
	}
	err = json.Unmarshal(bytes, &receiver)
	if err != nil {
		return fmt.Errorf("error while parsing config file \n\t%v", err)
	}
	return nil
}
