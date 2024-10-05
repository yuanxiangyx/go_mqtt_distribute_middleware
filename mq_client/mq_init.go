package mq_client

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"mqtt_pro/config"
)

type MqClientHandler struct {
	MqClient      mqtt.Client     `json:"mq_client"`
	SubDealConfig *config.SubDeal `json:"sub_deal_config"`
}

func (mq *MqClientHandler) SubProcess() error {
	cfg := mq.SubDealConfig
	token := mq.MqClient.Subscribe(cfg.SubTopic.Topic, cfg.SubTopic.Qos, MessageDeal)
	if token != nil {
		return token.Error()
	}
	return nil
}

func InitMqClient(cfg *config.Config) []MqClientHandler {
	var MqHandlers []MqClientHandler
	for _, broker := range cfg.Brokers {
		MqHandlers = append(MqHandlers, MqClientHandler{
			MqClient:      GetMqttClient(broker),
			SubDealConfig: broker.SubDealConfig,
		})
	}
	return MqHandlers
}
