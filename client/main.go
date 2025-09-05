package main

import (
	"context"
	"log"

	sidelogspb "github.com/ThibaultLonguepee/sidelogs-api/gen/sidelogs-proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	grpcClient, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect client. %v", err)
	}

	defer grpcClient.Close()

	logsClient := sidelogspb.NewLogServiceClient(grpcClient)
	_, err = logsClient.WriteLog(
		context.Background(),
		&sidelogspb.Log{
			Id:      0,
			Title:   "Sample",
			Message: "This is a sample message",
		},
	)
	if err != nil {
		log.Fatalf("Error trying to write log. %v", err)
	}
}
