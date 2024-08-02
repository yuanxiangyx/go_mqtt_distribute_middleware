package mq

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
	"mqtt_pro/config"
	"regexp"
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
		//fmt.Println(clientId, broker.ClientId)
		if clientId != broker.ClientId {
			continue
		} else {
			for _, SubDeal := range broker.SubDealSlice {
				for _, excludeTopic := range SubDeal.ExcludeTopics {
					// match topic name
					ok, err := regexp.MatchString(excludeTopic, messageTopic)
					if err != nil {
						zap.S().Errorf("%s", err)
					}
					// exclude topic
					if ok {
						continue
					} else {
						for _, url := range SubDeal.ApiCallbackUrl {
							go DealSubscribeTopic(SubDeal.CallbackMethod, url, payLoad, SubDeal.Retry.MaxAttempts)
						}
						zap.S().Infof("ClientId：%s %s Message: %s", clientId, messageTopic, payLoad)
					}
				}
			}
		}
	}
}

func DealSubscribeTopic(callbackMethod string, url string, payLoad string, retry int) {
	fmt.Println(callbackMethod, url, payLoad)
}

func Subscribe(c mqtt.Client, topic string, qos byte) {
	if token := c.Subscribe(topic, qos, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func Publish(c mqtt.Client, topic string, payload []byte) {
	c.Publish(topic, 0, false, payload)
}

func DealBrokerMessage(cfg *config.MqttConfig) {
	// Subscription Topic
	dvgList := cfg.SubDealSlice
	for _, v := range dvgList {
		c := GetMqttClient(cfg)
		for _, topic := range v.Topics {
			Subscribe(c, topic.Topic, topic.Qos)
		}
	}
}
