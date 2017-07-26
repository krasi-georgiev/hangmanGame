package main

import (
	"github.com/krasi-georgiev/hangmanGame/api"
)

func newGallow() (*api.Gallow, error) {
	ctx, cancel := appContext()
	defer cancel()

	client, err := getGRPCConnection(&ctx)
	if err != nil {
		return nil, err
	}
	r, err := client.NewGallow(ctx, &api.GallowRequest{RetryLimit: 5})
	if err != nil {
		return nil, err
	}
	return r.Gallow[0], nil
}
