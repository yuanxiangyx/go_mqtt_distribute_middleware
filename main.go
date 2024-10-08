package main

import (
	"mqtt_pro/apps"
	"mqtt_pro/config"
	"mqtt_pro/mq_client"
	"mqtt_pro/utils"
)

func SubAllMqttMessage(MqClientHandler []mq_client.MqClientHandler) error {
	for _, mqClient := range MqClientHandler {
		err := mqClient.SubProcess()
		if err != nil {
			return err
		}
	}
	return nil
}

func MqProcess(cfg *config.Config) {
	// Sub All Messages
	err := SubAllMqttMessage(mq_client.InitMqClient(cfg))
	if err != nil {
		panic(err)
	}
}

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	err = utils.InitLogger(cfg)
	if err != nil {
		panic(err)
	}
	go MqProcess(cfg)
	err = apps.WebApp()
	if err != nil {
		panic(err)
	}
}
