package main

import (
	"fmt"
	"net"
	"server/internal/util"
	"server/pkg/pb"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	DEFAULT_PORT = 10003
)

var (
	logEntry = logrus.NewEntry(logrus.StandardLogger())
)

type ServerService struct {
	pb.UnimplementedGreeterServer
}

func (s *ServerService) SayHello(stream pb.Greeter_SayHelloServer) error {
	logEntry.Info("SayHello called")

	for {
		resp, err := stream.Recv()
		if err != nil {
			logEntry.Errorf("Error reading from stream: %v", err)
			return err
		}

		logEntry.Infof("Received request: %v", resp)

		if err := stream.Send(&pb.HelloReply{Message: "Hello " + resp.Name}); err != nil {
			logEntry.Errorf("Error sending to stream: %v", err)
			return err
		}
	}
}

func main() {
	logEntry.Info("Starting server...")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", DEFAULT_PORT))
	if err != nil {
		logEntry.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer(util.NewGRPCServerOptions()...)

	pb.RegisterGreeterServer(server, &ServerService{})

	logEntry.Infof("Server is Listening on Port : %d", DEFAULT_PORT)

	if err := server.Serve(lis); err != nil {
		logEntry.Fatalf("failed to serve: %v", err)
	}
}
