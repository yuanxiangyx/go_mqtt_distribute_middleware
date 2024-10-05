package main

import (
	config2 "mqtt_pro/config"
	"mqtt_pro/mq_client"
	"mqtt_pro/utils"
	"time"
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

func main() {
	config, err := config2.InitConfig()
	if err != nil {
		panic(err)
	}
	err = utils.InitLogger(config)
	if err != nil {
		panic(err)
	}
	// Sub All Messages
	err = SubAllMqttMessage(mq_client.InitMqClient(config))
	if err != nil {
		panic(err)
	}
	// Loop to maintain client connectivity
	for {
		time.Sleep(10 * time.Millisecond)
	}
}
