package mq_client

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"mqtt_pro/config"
	"strconv"
	"time"
)

var mqttCfg, _ = config.InitConfig()

func GetMqttClient(cfg *config.MqttConfig) mqtt.Client {
	// Create MQTT client option
	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%s", cfg.BrokerIp, strconv.Itoa(cfg.BrokerPort)))
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
	fmt.Println(clientId, messageTopic, payLoad)
}

func Publish(c mqtt.Client, topic string, payload []byte) {
	c.Publish(topic, 0, false, payload)
}
