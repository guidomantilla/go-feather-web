package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	feather_commons_log "github.com/guidomantilla/go-feather-commons/pkg/log"
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

	feather_commons_log.Info(fmt.Sprintf("starting up - starting grpc server: %s", server.address))

	var err error
	var listener net.Listener
	if listener, err = net.Listen("tcp", server.address); err != nil {
		feather_commons_log.Error(fmt.Sprintf("starting up - starting grpc server error: %s", err.Error()))
		return err
	}

	if err = server.internal.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
		feather_commons_log.Error(fmt.Sprintf("starting up - starting grpc server error: %s", err.Error()))
		return err
	}

	return nil
}

func (server *GrpcServer) Stop(ctx context.Context) error {

	feather_commons_log.Info("shutting down - stopping grpc server")
	server.internal.GracefulStop()
	feather_commons_log.Info("shutting down - grpc server stopped")

	return nil
}
