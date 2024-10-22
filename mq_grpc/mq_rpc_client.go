package mq_grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mqtt_pro/mq_grpc/pb"
	"net"
)

type server struct {
	__.MqGreeterRpcServiceServer
}

func (s *server) SendMqMessage(ctx context.Context, in *__.MqRpcRequest) (*__.MqRpcResponse, error) {
	fmt.Println(in.String())
	return &__.MqRpcResponse{Code: 200, Message: "Success"}, nil
}

func RpcRun() {
	lis, err := net.Listen("tcp", ":8982")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	__.RegisterMqGreeterRpcServiceServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
