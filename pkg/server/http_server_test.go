package server

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/qmdx00/lifecycle"
)

func TestBuildHttpServer(t *testing.T) {
	type args struct {
		server *http.Server
	}
	tests := []struct {
		name string
		args args
		want lifecycle.Server
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildHttpServer(tt.args.server); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildHttpServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHttpServer_Run(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		server  *HttpServer
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.server.Run(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("HttpServer.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHttpServer_Stop(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		server  *HttpServer
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.server.Stop(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("HttpServer.Stop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
