package mq_client

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
	"mqtt_pro/config"
	"mqtt_pro/logger"
	"mqtt_pro/requests"
	"mqtt_pro/schemas"
	"mqtt_pro/utils"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type MqClientHandler struct {
	MqClient      mqtt.Client     `json:"mq_client"`
	SubDealConfig *config.SubDeal `json:"sub_deal_config"`
}

// SubProcess  Subscription topic
func (mq MqClientHandler) SubProcess() error {
	cfg := mq.SubDealConfig
	token := mq.MqClient.Subscribe(cfg.SubTopic.Topic, cfg.SubTopic.Qos, mq.MessageDeal)
	if token != nil {
		return token.Error()
	}
	return nil
}

// MessageDeal Process received messages
func (mq MqClientHandler) MessageDeal(client mqtt.Client, msg mqtt.Message) {
	messageTopic := msg.Topic()
	payLoad := string(msg.Payload())
	//reader := client.OptionsReader()
	//clientId := reader.ClientID()

	if mq.SubDealConfig.Enabled {
		exclude := mq.DealExcludeTopic(messageTopic)
		if exclude {
			return
		}
		mq.DistributeTopicContent(payLoad)
	}
}

// DealExcludeTopic Topic for handling filtering
func (mq MqClientHandler) DealExcludeTopic(receiveTopic string) bool {
	for _, i := range mq.SubDealConfig.ExcludeTopics {
		exclude, err := regexp.MatchString(i, receiveTopic)
		if err != nil {
			zap.Error(err)
			return true
		}
		if exclude {
			return true
		}
	}
	return false
}

func (mq MqClientHandler) DistributeTopicContent(payLoad string) {
	callbackMethod := strings.ToUpper(mq.SubDealConfig.CallbackMethod)
	go mq.RetryCallBack(callbackMethod, payLoad)
}

func (mq MqClientHandler) RetryCallBack(callbackMethod string, payLoad string) {
	retryCount := 0
	for retryCount <= mq.SubDealConfig.Retry {
		var err error
		switch callbackMethod {
		case "HTTP":
			err = mq.HttpCallBackDeal(payLoad)
		case "HTTPS":
			err = mq.HttpCallBackDeal(payLoad)
		case "GRPC":
			err = mq.GrpcCallBackDeal(payLoad)
		}
		if err == nil {
			return
		} else {
			retryCount += 1
			time.Sleep(1 * time.Second)
		}
	}
}

func (mq MqClientHandler) HttpCallBackDeal(payLoad string) (err error) {
	for _, addr := range mq.SubDealConfig.CallbackAddress {
		mapData, _ := utils.ParserPayLoadDataToMap(payLoad)
		data, err := requests.Post(requests.Args{
			Url:  addr,
			Json: mapData,
		})
		if err != nil {
			logger.WError(fmt.Sprintf("Post data err response: %s", string(data)))
			return err
		}
		logger.WInfo(fmt.Sprintf("%s-->%s Post data response: %s", mq.SubDealConfig.AppId, mq.SubDealConfig.AppName, string(data)))
	}
	return nil
}

func (mq MqClientHandler) GrpcCallBackDeal(payLoad string) error {
	for _, addr := range mq.SubDealConfig.CallbackAddress {
		headString, bodyString := utils.ParserPayLoadDataToString(payLoad)
		stringSchema := schemas.MqStringSchema{
			Header: headString,
			Body:   bodyString,
		}
		data, err := requests.GrpcRequest(addr, stringSchema)
		if err != nil {
			logger.WError(fmt.Sprintf("Grpc response: %s", data))
			return err
		}
		logger.WInfo(fmt.Sprintf("%s-->%s Grpc response: %s", mq.SubDealConfig.AppId, mq.SubDealConfig.AppName, data))
	}
	return nil
}

func (mq MqClientHandler) Publish(topic string, payload []byte) {
	mq.MqClient.Publish(topic, 0, false, payload)
}

// GetMqttClient Create MQTT client option
func GetMqttClient(cfg *config.MqttConfig) mqtt.Client {
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
