package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Go-gRPC-Server-Client/api/apipb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Api(ctx context.Context, request *apipb.ApiRequest) (*apipb.ApiResponse, error) {

	name := request.Name

	response := &apipb.ApiResponse{
		Greeting: "Welcome " + name,
	}

	return response, nil

}

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v  ", address)

	s := grpc.NewServer()

	apipb.RegisterApiServiceServer(s, &server{})

	s.Serve(lis)

}
