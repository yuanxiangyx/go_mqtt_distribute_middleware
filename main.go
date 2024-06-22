package main

import (
	config2 "mqtt_pro/config"
	"mqtt_pro/mq"
	"mqtt_pro/utils"
	"time"
)

func main() {
	config, err := config2.InitConfig()
	if err != nil {
		panic(err)
	}
	err = utils.InitLogger(config)
	if err != nil {
		panic(err)
	}

	DealMqttMessage(config)

	go func() {

	}()
	// Loop to maintain client connectivity
	for {
		time.Sleep(10 * time.Millisecond)
	}
}

func DealMqttMessage(cfg *config2.Config) {
	for _, val := range cfg.Brokers {
		mq.DealSubMessage(val)
	}
}
