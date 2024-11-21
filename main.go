package main

import (
	"fmt"
	"mqtt_pro/config"
	"mqtt_pro/logger"
	"mqtt_pro/mq_client"
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
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	logger.InitGroupLog(cfg.LogOption)
	if err != nil {
		panic(err)
	}
	// Sub All Messages
	err = SubAllMqttMessage(mq_client.InitMqClient(cfg))
	if err != nil {
		panic(err)
	}
	// Loop to maintain client connectivity
	fmt.Println("start...........success")
	fmt.Println(`
       ██     ██       
      █  █   █  █	
     █    █ █    █       
    █     ███     █      
`)
	for {
		time.Sleep(10 * time.Millisecond)
	}
}
