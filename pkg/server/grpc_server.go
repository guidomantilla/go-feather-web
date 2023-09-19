package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"github.com/qmdx00/lifecycle"
	"google.golang.org/grpc"
)

var _ lifecycle.Server = (*GrpcServer)(nil)

type GrpcServer struct {
	address  string
	internal *grpc.Server
}

func BuildGrpcServer(address string, server *grpc.Server) lifecycle.Server {
	return &GrpcServer{
		address:  address,
		internal: server,
	}
}

func (server *GrpcServer) Run(ctx context.Context) error {

	slog.Info("starting up - starting grpc server")

	var err error
	var listener net.Listener
	if listener, err = net.Listen("tcp", server.address); err != nil {
		slog.Error(fmt.Sprintf("starting up - starting grpc server error: %s", err.Error()))
	}

	if err = server.internal.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error(fmt.Sprintf("starting up - starting grpc server error: %s", err.Error()))
		return err
	}

	return nil
}

func (server *GrpcServer) Stop(ctx context.Context) error {

	slog.Info("shutting down - stopping grpc server")
	server.internal.GracefulStop()
	slog.Info("shutting down - grpc server stopped")

	return nil
}
