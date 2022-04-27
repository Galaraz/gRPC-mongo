package main

import (
	user "Galaraz/gRPC-mongo/proto/gen"
	"Galaraz/gRPC-mongo/server"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":8200")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("testo aleatorio")
	usr := server.Server{}

	grpcServer := grpc.NewServer()

	user.RegisterUserServiceServer(grpcServer, &usr)

	log.Println("Listening on Port: 8200!")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
