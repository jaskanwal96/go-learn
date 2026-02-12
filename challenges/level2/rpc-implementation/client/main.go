package main

import (
	"context"
	"log"
	"time"

	hellopb "challenge1_2/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := hellopb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &hellopb.HelloRequest{
		Name: "Sunny",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.Message)
}
