package main

import (
	"github.com/krasi-georgiev/hangmanGame/api"
)

func listGallows(client api.HangmanClient) ([]*api.Gallow, error) {
	ctx, cancel := appContext()
	defer cancel()
	games, err := client.ListGallows(ctx, &api.GallowRequest{Id: -1})
	if err != nil {
		return nil, err
	}

	return games.Gallow, nil
}
