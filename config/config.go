package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var FILE = "./config.json"

type Config struct {
	Command     string
	Template    string
	StoragePath string
}

func CreateVMCommand() string {
	config := new(Config)
	config.loadConfig()
	return config.Command
}

func Template() string {
	config := new(Config)
	config.loadConfig()
	return config.Template
}

func StoragePath() string {
	config := new(Config)
	config.loadConfig()
	return config.StoragePath
}

func (c *Config) loadConfig() error {
	bytes, err := ioutil.ReadFile(FILE)

	if err != nil {
		log.Printf("Failed to load %s: %v\n", FILE, err)
		return err
	}

	err = json.Unmarshal(bytes, &c)

	if err != nil {
		log.Printf("Failed parse %s: %v\n", FILE, err)
		return err
	}

	return nil
}
