package mq

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
	"mqtt_pro/config"
	"strconv"
	"time"
)

var mqttCfg, _ = config.InitConfig()

func GetMqttClient(cfg *config.MqttConfig) mqtt.Client {
	// Create MQTT client option
	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%s", cfg.Broker, strconv.Itoa(cfg.Port)))
	opts.SetClientID(cfg.ClientId)
	opts.SetUsername(cfg.Username)
	opts.SetPassword(cfg.Password)
	opts.SetCleanSession(true)
	opts.SetAutoReconnect(true)

	opts.SetKeepAlive(time.Duration(cfg.Alive) * time.Second)
	opts.SetDefaultPublishHandler(MessageDeal)

	// Create client
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return c
}

func MessageDeal(client mqtt.Client, msg mqtt.Message) {
	// Process received messages
	messageTopic := msg.Topic()
	payLoad := string(msg.Payload())
	reader := client.OptionsReader()
	clientId := reader.ClientID()
	for _, broker := range mqttCfg.Brokers {
		if clientId == broker.ClientId {
			for _, SubDeal := range broker.SubDealSlice {
				excludeFlag := false
				for _, excludeTopic := range SubDeal.ExcludeTopics {
					if messageTopic == excludeTopic {
						excludeFlag = true
						break
					}
				}
				if !excludeFlag {
					break
				} else {
					DealExcludeTopic()
				}
			}
			zap.S().Infof("Received message on topic: %s\nMessage: %s\n", messageTopic, payLoad)
		} else {
			continue
		}
	}

	//fmt.Printf("Received message on topic: %s\nMessage: %s\n", messageTopic, payLoad)
}

func DealExcludeTopic() {

}

func DealSubscribeTopic() {

}

func Subscribe(c mqtt.Client, topic string, qos byte) {
	if token := c.Subscribe(topic, qos, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func Publish(c mqtt.Client, topic string, payload []byte) {
	c.Publish(topic, 0, false, payload)
}

func DealSubMessage(cfg *config.MqttConfig) {
	// Subscription Topic
	dvgList := cfg.SubDealSlice
	for _, v := range dvgList {
		c := GetMqttClient(cfg)
		for _, topic := range v.Topics {
			Subscribe(c, topic, 0)
		}
	}
}
