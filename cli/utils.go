package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/krasi-georgiev/hangmanGame/api"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

var grpcClient api.HangmanClient

func getGRPCConnection(context *cli.Context) (api.HangmanClient, error) {
	if grpcClient != nil {
		return grpcClient, nil
	}

	bindSocket := context.GlobalString("address")
	dialOpts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithTimeout(context.GlobalDuration("connect-timeout"))}
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
func appContext(clicontext *cli.Context) (context.Context, context.CancelFunc) {
	var (
		ctx     = context.Background()
		timeout = clicontext.GlobalDuration("timeout")
		cancel  = func() {}
	)

	if timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, timeout)
	} else {
		ctx, cancel = context.WithCancel(ctx)
	}

	return ctx, cancel
}
