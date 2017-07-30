package main

import (
	"context"
	"fmt"
	"time"

	"github.com/krasi-georgiev/hangmanGame/api"
	"google.golang.org/grpc"
)

var grpcClient api.HangmanClient

func getGRPCConnection() (api.HangmanClient, error) {
	if grpcClient != nil {
		return grpcClient, nil
	}

	bindSocket := ":9999"
	dialOpts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}

	ctx, cancel := appContext()
	defer cancel()
	conn, err := grpc.DialContext(ctx, bindSocket, dialOpts...)

	if err != nil {
		return nil, fmt.Errorf("failed to dial %q: %s", bindSocket, err.Error())
	}
	grpcClient = api.NewHangmanClient(conn)
	return grpcClient, nil
}

func appContext() (context.Context, context.CancelFunc) {
	var (
		ctx     = context.Background()
		timeout = 2 * time.Second
	)
	return context.WithTimeout(ctx, timeout)
}
