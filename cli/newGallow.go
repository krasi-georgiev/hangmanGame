package main

import (
	"github.com/krasi-georgiev/hangmanGame/api"
)

func newGallow(client api.HangmanClient) (*api.Gallow, error) {
	ctx, cancel := appContext()
	defer cancel()
	r, err := client.NewGallow(ctx, &api.GallowRequest{RetryLimit: 5})
	if err != nil {
		return nil, err
	}
	return r.Gallow[0], nil
}
