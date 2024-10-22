package requests

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	mq_rpc_message "mqtt_pro/mq_grpc/pb"
	"mqtt_pro/schemas"
	"time"
)

func GrpcRequest(rpcAddr string, mqMessage schemas.MqSchema) (data string, err error) {
	conn, err := grpc.NewClient(rpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}
	defer conn.Close()
	c := mq_rpc_message.NewMqGreeterRpcServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	r, err := c.SendMqMessage(ctx, &mq_rpc_message.MqRpcRequest{Header: mqMessage.Header, Body: mqMessage.Body})
	if err != nil {
		return "", err
	}
	return r.String(), nil
}
