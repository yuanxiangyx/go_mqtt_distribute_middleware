package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Brokers    []*MqttConfig `json:"mqtt_brokers"`
	*LogOption `json:"log_config"`
}

type MqttConfig struct {
	ClientId      string   `json:"client_id"`
	Username      string   `json:"username"`
	Password      string   `json:"password"`
	Alive         int      `json:"alive"`
	BrokerIp      string   `json:"broker_ip"`
	BrokerPort    int      `json:"broker_port"`
	SubDealConfig *SubDeal `json:"sub_deal_config"`
}

type SubDeal struct {
	AppName         string      `json:"app_name"`
	AppId           string      `json:"app_id"`
	Enabled         bool        `json:"enabled"`
	CallbackMethod  string      `json:"callbackMethod"`
	CallbackAddress []string    `json:"callbackAddress"`
	SubTopic        TopicConfig `json:"subTopic"`
	ExcludeTopics   []string    `json:"excludeTopics"`
	Retry           int         `json:"retry"`
}

type TopicConfig struct {
	Topic string `json:"topic"`
	Qos   byte   `json:"qos"`
}

type LogOption struct {
	Level      string `json:"level"`
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
