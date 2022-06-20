package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/firekitz/fk-daemon-iam/config"
	postgres "github.com/firekitz/fk-daemon-iam/internal/Database"
	"github.com/firekitz/fk-daemon-iam/internal/gateway"
	"github.com/firekitz/fk-daemon-iam/internal/server"
	log "github.com/firekitz/fk-lib-log-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func matcher(key string) (string, bool) {
	if strings.EqualFold(key, "x-request-fk-src-service-name") {
		return key, true
	}
	return "", false
}

func main() {
	flag.Parse()

	_, err := config.LoadConfig(".")
	if err != nil {
		panic(fmt.Errorf("Fatal error Config file: %w\n", err))
	}

	postgres.DatabaseInit()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		cancel()
		fmt.Println("Exiting server on ", sig)
		os.Exit(0)
	}()

	var (
		serverAddr = flag.String("serverAddr", ":"+config.LoadedConfig.GRPC_PORT, "endpoint of the gRPC server")
		network    = flag.String("network", "tcp", `one of "tcp" or "unix". Must be consistent to -endpoint`)
	)

	var muxOptions []runtime.ServeMuxOption
	muxOptions = append(muxOptions, runtime.WithIncomingHeaderMatcher(matcher))
	opts := gateway.Options{
		Addr: ":" + config.LoadedConfig.HTTP_PORT,
		GRPCServer: gateway.Endpoint{
			Network: *network,
			Addr:    *serverAddr,
		},
		Mux: muxOptions,
	}

	go func() {
		//fmt.Println("Starting HTTP gateway on", *gatewayAddr)
		if err := gateway.Run(ctx, opts); err != nil {
			log.F(err)
		}
	}()

	fmt.Println("Starting gRPC server on", *serverAddr)
	if err := server.Run(ctx, *network, *serverAddr); err != nil {
		log.F(err)
	}
}
