package mq_client

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"mqtt_pro/config"
	"strconv"
	"time"
)

type MqClientHandler struct {
	MqClient      mqtt.Client     `json:"mq_client"`
	SubDealConfig *config.SubDeal `json:"sub_deal_config"`
}

func (mq *MqClientHandler) SubProcess() error {
	cfg := mq.SubDealConfig
	token := mq.MqClient.Subscribe(cfg.SubTopic.Topic, cfg.SubTopic.Qos, mq.MessageDeal)
	if token != nil {
		return token.Error()
	}
	return nil
}

func (mq *MqClientHandler) MessageDeal(client mqtt.Client, msg mqtt.Message) {
	// Process received messages
	messageTopic := msg.Topic()
	payLoad := string(msg.Payload())
	reader := client.OptionsReader()
	clientId := reader.ClientID()
	fmt.Println(clientId, messageTopic, payLoad)
}

func (mq *MqClientHandler) Publish(topic string, payload []byte) {
	mq.MqClient.Publish(topic, 0, false, payload)
}

func GetMqttClient(cfg *config.MqttConfig) mqtt.Client {
	// Create MQTT client option
	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%s", cfg.BrokerIp, strconv.Itoa(cfg.BrokerPort)))
	opts.SetClientID(cfg.ClientId)
	opts.SetUsername(cfg.Username)
	opts.SetPassword(cfg.Password)
	opts.SetCleanSession(true)
	opts.SetAutoReconnect(true)

	opts.SetKeepAlive(time.Duration(cfg.Alive) * time.Second)

	// Create client
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return c
}
