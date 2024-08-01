package main

import (
	config2 "mqtt_pro/config"
	"mqtt_pro/mq"
	"mqtt_pro/utils"
	"time"
)

func SubAllMqttMessage(cfg *config2.Config) {
	for _, val := range cfg.Brokers {
		mq.DealBrokerMessage(val)
	}
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
	SubAllMqttMessage(config)

	go func() {

	}()
	// Loop to maintain client connectivity
	for {
		time.Sleep(10 * time.Millisecond)
	}
}
