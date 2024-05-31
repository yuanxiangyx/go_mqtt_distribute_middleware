package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	*MqttConfig `json:"mqtt_config"`
	*LogConfig  `json:"log_config"`
}

type MqttConfig struct {
	ClientId string `json:"client_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Alive    int    `json:"alive"`
	Broker   string `json:"broker"`
	Port     int    `json:"port"`
}

type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxSize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

var config = new(Config)

func InitConfig() *Config {
	content, err := os.ReadFile("config.json")
	// Check if the folder exists
	dir, err := filepath.Abs(filepath.Dir("."))
	info, _ := os.Stat(filepath.Join(dir, "logs"))
	if info == nil {
		_ = os.Mkdir(filepath.Join(dir, "logs"), 0777)
	}
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, &config)
	if err != nil {
		panic(err)
	}
	return config
}
