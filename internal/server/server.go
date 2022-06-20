package server

import (
	"context"
	"github.com/firekitz/fk-daemon-iam/config"
	iampb "github.com/firekitz/fk-daemon-iam/internal/proto/iam"
	log "github.com/firekitz/fk-lib-log-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net"
	"net/http"
)

// Run starts the example gRPC service.
// "network" and "address" are passed to net.Listen.
func Run(ctx context.Context, network, address string) error {
	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			log.E("Failed to close %s %s: %v", network, address, err)
		}
	}()

	s := config.Middleware()
	iampb.RegisterIamServer(s, new(iamServer))
	log.I("Starting listening at %s", address)
	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()
	return s.Serve(l)
}

//RunInProcessGateway starts the invoke in process http gateway.
func RunInProcessGateway(ctx context.Context, addr string, opts ...runtime.ServeMuxOption) error {
	mux := runtime.NewServeMux(opts...)
	err := iampb.RegisterIamHandlerServer(ctx, mux, new(iamServer))
	if err != nil {
		log.E("Don't Handled HTTP: %v", err)
	}

	s := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		log.I("Shutting down the http gateway server")
		if err := s.Shutdown(context.Background()); err != nil {
			log.E("Failed to shutdown http gateway server: %v", err)
		}
	}()

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.E("Failed to listen and serve: %v", err)
		return err
	}
	return nil

}
