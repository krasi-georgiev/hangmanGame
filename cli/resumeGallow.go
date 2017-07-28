package main

import (
	"github.com/krasi-georgiev/hangmanGame/api"
)

func resumeGallow(client api.HangmanClient, id int) (*api.Gallow, error) {
	ctx, cancel := appContext()
	defer cancel()
	g, err := client.ResumeGallow(ctx, &api.GallowRequest{Id: int32(id)})
	if err != nil {
		return nil, err
	}
	return g, nil
}
