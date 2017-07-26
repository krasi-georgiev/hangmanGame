package main

import (
	"fmt"

	"github.com/krasi-georgiev/hangmanGame/api"
)

func listGallows(client api.HangmanClient) error {
	ctx, cancel := appContext()
	defer cancel()
	games, err := client.ListGallows(ctx, &api.GallowRequest{})
	if err != nil {
		return err
	}

	fmt.Println("sdfsdf")
	fmt.Println(games)

	return nil
}
