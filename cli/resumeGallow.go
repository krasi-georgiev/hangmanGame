package main

import (
	"errors"
	"strconv"

	"github.com/krasi-georgiev/hangmanGame/api"
)

func resumeGallow(client api.HangmanClient, id string) (*api.Gallow, error) {
	i, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("Invalid Game ID")
	}

	ctx, cancel := appContext()
	defer cancel()
	g, err := client.ResumeGallow(ctx, &api.GallowRequest{Id: int32(i)})
	if err != nil {
		return nil, err
	}
	return g, nil
}
