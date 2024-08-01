package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Brokers    []*MqttConfig `json:"mqtt_brokers"`
	*LogConfig `json:"log_config"`
}

type MqttConfig struct {
	ClientId     string     `json:"client_id"`
	Username     string     `json:"username"`
	Password     string     `json:"password"`
	Alive        int        `json:"alive"`
	Broker       string     `json:"broker"`
	Port         int        `json:"port"`
	SubDealSlice []*SubDeal `json:"sub_deal_config"`
}

type SubDeal struct {
	SubId          string      `json:"sub_id"`
	AppName        string      `json:"app_name"`
	Enabled        string      `json:"enabled"`
	CallbackMethod string      `json:"callbackMethod"`
	ApiCallbackUrl []string    `json:"apiCallbackUrl"`
	Topics         []string    `json:"topics"`
	ExcludeTopics  []string    `json:"excludeTopics"`
	Retry          RetryConfig `json:"retry"`
}

type RetryConfig struct {
	MaxAttempts int `json:"max_attempts"`
}

type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxSize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

var config = new(Config)

func InitConfig() (cfg *Config, err error) {
	content, err := os.ReadFile("config.json")
	// Check if the folder exists
	dir, err := filepath.Abs(filepath.Dir("."))
	info, _ := os.Stat(filepath.Join(dir, "logs"))
	if info == nil {
		_ = os.Mkdir(filepath.Join(dir, "logs"), 0777)
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, &config)
	if err != nil {
		panic(err)
	}
	return config, err
}
