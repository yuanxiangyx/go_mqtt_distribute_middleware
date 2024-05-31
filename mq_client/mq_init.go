package mq_client

import (
	"mqtt_pro/config"
)

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
