package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/qmdx00/lifecycle"
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

	slog.Info("starting up - starting http server")

	if err := server.internal.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error(fmt.Sprintf("starting up - starting http server error: %s", err.Error()))
		return err
	}

	return nil
}

func (server *HttpServer) Stop(ctx context.Context) error {

	slog.Info("shutting down - stopping http server")

	if err := server.internal.Shutdown(ctx); err != nil {
		slog.Error(fmt.Sprintf("shutting down - forced shutdown: %s", err.Error()))
		return err
	}

	slog.Info("shutting down - http server stopped")
	return nil
}
