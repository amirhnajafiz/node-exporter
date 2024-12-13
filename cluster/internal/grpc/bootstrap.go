package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/F24-CSE535/2pcbyz/cluster/internal/storage"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Bootstrap is a wrapper that holds every required thing for the gRPC server starting.
type Bootstrap struct {
	ServicePort int
	Logger      *zap.Logger
	Storage     *storage.Storage
	Queue       chan context.Context // this queue is the handler's input channel
}

// ListenAnsServer creates a new gRPC instance with all required services.
func (b *Bootstrap) ListenAndServer() error {
	// on the local network, listen to a port
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", b.ServicePort))
	if err != nil {
		return fmt.Errorf("failed to start the listener server: %v", err)
	}

	// create a new grpc instance
	server := grpc.NewServer(
		grpc.UnaryInterceptor(b.allUnaryInterceptor),   // set an unary interceptor
		grpc.StreamInterceptor(b.allStreamInterceptor), // set a stream interceptor
	)

	// register all gRPC services

	// starting the server
	b.Logger.Info("grpc started", zap.Int("port", b.ServicePort))
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to start the server: %v", err)
	}

	return nil
}
