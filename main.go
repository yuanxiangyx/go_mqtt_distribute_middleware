package main

import (
	config2 "mqtt_pro/config"
	"mqtt_pro/mq"
	"mqtt_pro/utils"
	"time"
)

func main() {
	config := config2.InitConfig()
	utils.InitLogger(config)
	c := mq.GetMqttClient(config)
	// Subscription Topic
	mq.Subscribe(c, "from/#", 0)
	// Publish Topic
	//go func() {
	//	for {
	//		topic := fmt.Sprintf("go/mqtt/%d", rand.Int63())
	//		Publish(c, topic, []byte(fmt.Sprintf("%s", time.Now())))
	//		time.Sleep(5 * time.Millisecond)
	//	}
	//}()

	// Loop to maintain client connectivity
	for {
		time.Sleep(10 * time.Millisecond)
	}
}
