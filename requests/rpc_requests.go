package requests

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	mq_rpc_message "mqtt_pro/mq_grpc/pb"
	"time"
)

func GrpcRequest() {
	conn, err := grpc.NewClient("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := mq_rpc_message.NewMqGreeterRpcServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendMqMessage(ctx, &mq_rpc_message.MqRpcRequest{Header: "aaa", Body: "vvv"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
