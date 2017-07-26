package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/krasi-georgiev/hangmanGame/api"
	"google.golang.org/grpc"
)

var grpcClient api.HangmanClient

func getGRPCConnection(context *context.Context) (api.HangmanClient, error) {
	if grpcClient != nil {
		return grpcClient, nil
	}

	bindSocket := ":9999"
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	dialOpts = append(dialOpts,
		grpc.WithDialer(func(addr string, timeout time.Duration) (net.Conn, error) {
			return net.DialTimeout("tcp", bindSocket, timeout)
		},
		))

	conn, err := grpc.Dial(bindSocket, dialOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to dial %q: %s", bindSocket, err.Error())
	}
	grpcClient = api.NewHangmanClient(conn)
	return grpcClient, nil
}

// appContext returns the context for a command.
func appContext() (context.Context, context.CancelFunc) {
	var (
		ctx     = context.Background()
		timeout = 2 * time.Second
		cancel  = func() {}
	)

	if timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, timeout)
	} else {
		ctx, cancel = context.WithCancel(ctx)
	}

	return ctx, cancel
}
