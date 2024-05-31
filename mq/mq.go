package mq

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
	"mqtt_pro/config"
	"os"
	"strconv"
	"time"
)

func GetMqttClient(cfg *config.MqttConfig) mqtt.Client {
	// Create MQTT client option
	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%s", cfg.Broker, strconv.Itoa(cfg.Port)))
	opts.SetClientID(cfg.ClientId)
	opts.SetUsername(cfg.Username)
	opts.SetPassword(cfg.Password)
	opts.SetCleanSession(true)

	opts.SetKeepAlive(time.Duration(cfg.Alive) * time.Second)
	opts.SetDefaultPublishHandler(MessageDeal)

	// Create client
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	return c
}

func MessageDeal(client mqtt.Client, msg mqtt.Message) {
	// Process received messages
	messageTopic := msg.Topic()
	payLoad := string(msg.Payload())
	zap.S().Infof("Received message on topic: %s\nMessage: %s\n", messageTopic, payLoad)
	//fmt.Println("Received message on topic: %s\nMessage: %s\n", messageTopic, payLoad)
}

func Subscribe(c mqtt.Client, topic string, qos byte) {
	if token := c.Subscribe(topic, qos, nil); token.Wait() && token.Error() != nil {
		zap.S().Info(token.Error())
		os.Exit(1)
	}
}

func Publish(c mqtt.Client, topic string, payload []byte) {
	c.Publish(topic, 0, false, payload)
}
