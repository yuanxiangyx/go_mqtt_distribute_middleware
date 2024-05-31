package main

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"math/rand"
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
	// Loop to maintain client connectivity
	//for {
	//	time.Sleep(10 * time.Millisecond)
	//}
}

func DealMqttMessage(cfg *config2.Config) {
	for _, val := range cfg.Brokers {
		//client := mq.GetMqttClient(val)
		DealSubMessage(val)
	}
}

func SubMessage(c mqtt.Client) {
	// Publish Topic
	go func() {
		for {
			topic := fmt.Sprintf("go/mqtt/%d", rand.Int63())
			mq.Publish(c, topic, []byte(fmt.Sprintf("%s", time.Now())))
			time.Sleep(5 * time.Millisecond)
		}
	}()
}

func DealSubMessage(cfg *config2.MqttConfig) {
	// Subscription Topic
	dvgList := cfg.SubDealSlice
	for _, v := range dvgList {
		res, _ := json.Marshal(v)
		fmt.Println(string(res))
		fmt.Println(v.Topics)
		c := mq.GetMqttClient(cfg)
		for topic := range v.Topics {
			fmt.Println(topic, v.Topics[topic])
			mq.Subscribe(c, "from/#", 0)
		}
	}
}
