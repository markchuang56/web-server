package cmd

import (
	"context"
	"flag"
	"fmt"

	//"../../protocol/grpc"
	//"../../protocol/rest"
	//"../../service/v1"

	"go-heroku/fitbit-demo/web-server/pkg/protocol/grpc"
	"go-heroku/fitbit-demo/web-server/pkg/protocol/rest"
	"go-heroku/fitbit-demo/web-server/pkg/service/v1"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string

	// DB Database parameters setion
	// DatastoreDBHost is host of database
	DatastoreDBHost string

	// HTTP/REST gateway start parameters section
	// HTTPPort is TCP port to listen by HTTP/REST gateway
	HTTPPort string

	//TO DO ...
}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.HTTPPort, "http-port", "", "HTTP port to bind")
	flag.StringVar(&cfg.DatastoreDBHost, "db-host", "", "Database host")

	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	if len(cfg.HTTPPort) == 0 {
		return fmt.Errorf("invalid TCP port for HTTP gateway: '%s'", cfg.HTTPPort)
	}

	v1API := v1.NewToDoServiceServer()

	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
