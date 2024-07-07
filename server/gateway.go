package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"yadro/gen/go/proto"
)

func GatewayRun(httpAddr, rpcAddr string) error {
	conn, err := grpc.NewClient(
		rpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("couldn't dial the server: %w", err)
	}

	mux := runtime.NewServeMux()
	err = proto.RegisterDnsServiceHandler(context.Background(), mux, conn)
	if err != nil {
		return fmt.Errorf("couldn't register the handler: %w", err)
	}

	fileServer := http.FileServer(http.Dir("server/static"))
	server := &http.Server{
		Addr: httpAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/api") {
				mux.ServeHTTP(w, r)
			} else {
				fileServer.ServeHTTP(w, r)
			}
		}),
	}

	return server.ListenAndServe()
}
