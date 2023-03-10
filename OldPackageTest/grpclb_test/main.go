package main

import (
	"context"
	"fmt"
	"grpclb_test/proto"
	"log"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(
		"consul://192.168.28.100:8500/user-srv?wait=14s&tag=lyc",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), //采用了轮询法
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for i := 0; i < 10; i++ {
		userSrvClient := proto.NewUserClient(conn)
		rsp, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
			Pn:    1,
			PSize: 2,
		})
		if err != nil {
			panic(err)
		}
		for index, data := range rsp.Data {
			fmt.Println(index, data)
		}
	}

}
