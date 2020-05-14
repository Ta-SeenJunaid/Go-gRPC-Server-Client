package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"github.com/Go-gRPC-Server-Client/api/apipb"
)

func main() {
	fmt.Println("Hello client ....")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := apipb.NewApiServiceClient(cc)
	request := &apipb.ApiRequest{Name: "Satoshi"}

	resp, _ := client.Api(context.Background(), request)
	fmt.Printf("Receive response => [%v]", resp.Greeting)

}
