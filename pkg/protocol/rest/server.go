package rest

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	//"../../api/v1"
	"go-heroku/fitbit-demo/web-server/pkg/api/v1"
)

// RusServer runs HTTP/REST gatway
func RunServer(ctx context.Context, grpcPort, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	port := os.Getenv("PORT")
	if port == "" {
		port = grpcPort
	}
	//fmt.Println(port)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := v1.RegisterToDoServiceHandlerFromEndpoint(ctx, mux, "localhost:"+port, opts); err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	log.Println("starting HTTP/REST gateway ...")
	return srv.ListenAndServe()
}
