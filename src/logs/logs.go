package logs

import (
	"context"
	"fmt"

	sidelogspb "github.com/ThibaultLonguepee/sidelogs-api/gen/sidelogs-proto"
)

type LogsServer struct {
	sidelogspb.UnimplementedLogServiceServer
}

func CreateLogsServer() *LogsServer {
	server := &LogsServer{}
	fmt.Println("LogsServer has been instantiated.")
	return server
}

func (srv *LogsServer) WriteLog(ctx context.Context, log *sidelogspb.Log) (*sidelogspb.Empty, error) {
	fmt.Println("A log was written!")
	return &sidelogspb.Empty{}, nil
}
