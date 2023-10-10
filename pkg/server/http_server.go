package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	feather_commons_log "github.com/guidomantilla/go-feather-commons/pkg/log"
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

	feather_commons_log.Info(fmt.Sprintf("starting up - starting http server: %s", server.internal.Addr))

	if err := server.internal.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		feather_commons_log.Error(fmt.Sprintf("starting up - starting http server error: %s", err.Error()))
		return err
	}

	return nil
}

func (server *HttpServer) Stop(ctx context.Context) error {

	feather_commons_log.Info("shutting down - stopping http server")

	if err := server.internal.Shutdown(ctx); err != nil {
		feather_commons_log.Error(fmt.Sprintf("shutting down - forced shutdown: %s", err.Error()))
		return err
	}

	feather_commons_log.Info("shutting down - http server stopped")
	return nil
}
