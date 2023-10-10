package util

import (
	"time"

	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
)

func NewGRPCDialOptions() []grpc.DialOption {
	opts := []grpc.DialOption{}

	si := []grpc.StreamClientInterceptor{}

	ui := []grpc.UnaryClientInterceptor{}

	grpcLogger := logrus.NewEntry(logrus.StandardLogger())

	si = append(si, grpc_logrus.StreamClientInterceptor(grpcLogger))
	ui = append(ui, grpc_logrus.UnaryClientInterceptor(grpcLogger))

	return append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(si...)),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(ui...)),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                20 * time.Second,
			Timeout:             10 * time.Second,
			PermitWithoutStream: true,
		}),
	)
}
