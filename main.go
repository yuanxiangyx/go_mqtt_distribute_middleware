package main

import (
	"fmt"
	"mqtt_pro/mq_client"
	"mqtt_pro/requests"
	"mqtt_pro/schemas"
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

func main() {
	//cfg, err := config.InitConfig()
	//if err != nil {
	//	panic(err)
	//}
	//err = utils.InitLogger(cfg)
	//if err != nil {
	//	panic(err)
	//}
	//// Sub All Messages
	//err = SubAllMqttMessage(mq_client.InitMqClient(cfg))
	//if err != nil {
	//	panic(err)
	//}
	//// Loop to maintain client connectivity
	//for {
	//	time.Sleep(10 * time.Millisecond)
	//}
	res, err := requests.GrpcRequest("127.0.0.1:8972", schemas.MqSchema{Header: "sad", Body: "asd"})
	if err != nil {
		fmt.Println("err：", err)
		return
	}
	fmt.Println(res)
}
