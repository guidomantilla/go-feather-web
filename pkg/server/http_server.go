package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/qmdx00/lifecycle"
	"go.uber.org/zap"
)

var _ lifecycle.Server = (*HttpServer)(nil)

type HttpServer struct {
	internal *http.Server
}

func BuildHttpServer(server *http.Server) lifecycle.Server {
	return &HttpServer{
		internal: server,
	}
}

func (server *HttpServer) Run(ctx context.Context) error {

	zap.L().Info("starting up - starting http server")

	if err := server.internal.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		zap.L().Error(fmt.Sprintf("starting up - starting http server error: %s", err.Error()))
		return err
	}

	return nil
}

func (server *HttpServer) Stop(ctx context.Context) error {

	zap.L().Info("shutting down - stopping http server")

	if err := server.internal.Shutdown(ctx); err != nil {
		zap.L().Error(fmt.Sprintf("shutting down - forced shutdown: %s", err.Error()))
		return err
	}

	zap.L().Info("shutting down - http server stopped")
	return nil
}
