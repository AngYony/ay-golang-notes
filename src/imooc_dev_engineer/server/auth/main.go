package main

import (
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("不能创建Logger:%v", err)
	}

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Fatal("不能监听", zap.Error(err))
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		Logger: logger,
	})

	err = s.Serve(lis)
	logger.Fatal("不能启动server", zap.Error(err))
}
