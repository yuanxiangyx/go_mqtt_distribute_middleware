package main

import (
	"github.com/gin-gonic/gin"
	config "mqtt_pro/config"
	"mqtt_pro/mq_client"
	"mqtt_pro/utils"
	"net/http"
)

func SubAllMqttMessage(MqClientHandler []mq_client.MqClientHandler) error {
	for _, mqClient := range MqClientHandler {
		err := mqClient.SubProcess()
		if err != nil {
			return err
		}
	}
	return nil
}

func MqProcess(cfg *config.Config) {
	// Sub All Messages
	err := SubAllMqttMessage(mq_client.InitMqClient(cfg))
	if err != nil {
		panic(err)
	}
	// Loop to maintain client connectivity
	//for {
	//	time.Sleep(10 * time.Millisecond)
	//}
}

func WebApp() error {
	r := gin.Default()
	// 返回一个json数据
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    http.StatusOK,
			"message": "ok",
		})
	})

	return r.Run()
}

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	err = utils.InitLogger(cfg)
	if err != nil {
		panic(err)
	}
	go MqProcess(cfg)
	err = WebApp()
	if err != nil {
		panic(err)
	}
}
