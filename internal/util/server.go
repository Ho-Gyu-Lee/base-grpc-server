package util

import (
	"time"

	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
)

func NewGRPCServerOptions() []grpc.ServerOption {
	opts := []grpc.ServerOption{}

	si := []grpc.StreamServerInterceptor{
		grpc_recovery.StreamServerInterceptor(),
	}

	ui := []grpc.UnaryServerInterceptor{
		grpc_recovery.UnaryServerInterceptor(),
	}

	grpcLogger := logrus.NewEntry(logrus.StandardLogger())

	si = append(si, grpc_logrus.StreamServerInterceptor(grpcLogger))
	ui = append(ui, grpc_logrus.UnaryServerInterceptor(grpcLogger))

	return append(opts,
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(si...)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(ui...)),
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             10 * time.Second,
				PermitWithoutStream: true,
			},
		),
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				Time:    20 * time.Second,
				Timeout: 10 * time.Second,
			},
		),
	)
}
