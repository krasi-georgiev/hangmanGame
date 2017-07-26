package main

import (
	"fmt"

	"github.com/krasi-georgiev/hangmanGame/api"
)

func listGallows() error {
	var (
		ctx, cancel = appContext()
	)
	defer cancel()

	client, err := getGRPCConnection(&ctx)
	if err != nil {
		return err
	}
	games, err := client.ListGallows(ctx, &api.GallowRequest{})
	if err != nil {
		return err
	}

	fmt.Println("sdfsdf")
	fmt.Println(games)

	return nil
}
