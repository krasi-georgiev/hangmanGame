package main

import (
	"fmt"

	"github.com/krasi-georgiev/hangmanGame/api"
)

func newGallow() error {
	var (
		ctx, cancel = appContext()
	)
	defer cancel()

	client, err := getGRPCConnection(&ctx)
	if err != nil {
		return err
	}
	games, err := client.ListGallows(ctx, &api.HangmanRequest{})
	if err != nil {
		return err
	}

	fmt.Println(games)

	return nil
}
