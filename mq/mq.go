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

func GetMqttClient(config *config.Config) mqtt.Client {
	// Create MQTT client option
	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%s", config.Broker, strconv.Itoa(config.Port)))
	opts.SetClientID(config.ClientId)
	opts.SetUsername(config.Username)
	opts.SetPassword(config.Password)
	opts.SetCleanSession(true)
	opts.SetKeepAlive(time.Duration(config.Alive) * time.Second)
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
	// 处理接收到的消息
	zap.S().Infof("Received message on topic: %s\nMessage: %s\n", msg.Topic(), string(msg.Payload()))
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", msg.Topic(), string(msg.Payload()))
}

func Subscribe(c mqtt.Client, topic string, qos byte) {
	if token := c.Subscribe(topic, qos, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}

func Publish(c mqtt.Client, topic string, payload []byte) {
	c.Publish(topic, 0, false, payload)
}
