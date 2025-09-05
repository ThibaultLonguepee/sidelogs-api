package main

import (
	"fmt"
	"log"
	"net"

	sidelogspb "github.com/ThibaultLonguepee/sidelogs-api/gen/sidelogs-proto"
	"github.com/ThibaultLonguepee/sidelogs-api/src/logs"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen. %v", err)
	}

	grpcServer := grpc.NewServer()

	sidelogspb.RegisterLogServiceServer(grpcServer, logs.CreateLogsServer())
	fmt.Println("LogsServer is now listening on", listener.Addr())

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve. %v", err)
	}
}
