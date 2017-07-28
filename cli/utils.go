package main

import (
	"context"
	"fmt"
	"time"
	"io"


	"github.com/krasi-georgiev/hangmanGame/api"
	"google.golang.org/grpc"
	"github.com/chzyer/readline"

)

var grpcClient api.HangmanClient

var gallowArt = []string{
	`
    _________
    |/      |
    |
    |
    |
    |
    |
____|____`,
	`
    _________
    |/      |
    |      (_)
    |
    |
    |
    |
____|____`,
	`
    __________
    |/      |
    |      (_)
    |       |
    |       |
    |
    |
____|____`,
	`
    __________
    |/      |
    |      (_)
    |      \|
    |       |
    |
    |
____|____`,
	`
    __________
    |/      |
    |      (_)
    |      \|/
    |       |
    |
    |
____|____`,
	`
    __________
    |/      |
    |      (_)
    |      \|/
    |       |
    |      /
    |
____|____`,
	`
    __________
    |/      |
    |      (_)
    |      \|/
    |       |
    |      / \
    |
____|____`,
}

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

func usage(w *readline.Instance) {
	w.SetPrompt("»")
	io.WriteString(w.Stdout(), completer.Tree("    "))
}
