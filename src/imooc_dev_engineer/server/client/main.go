package main

import (
	"context"
	trippb "coolcar/proto/gen/go"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.Lshortfile)
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("不能连接到服务：%v", err)
	}
	tsClient := trippb.NewTripServiceClient(conn)

	r, err := tsClient.GetTrip(context.Background(), &trippb.GetTripRequest{
		Id: "trip456",
	})

	if err != nil {
		log.Fatalf("无法调用GetTrip：%v", err)
	}

	fmt.Println(r)

}
