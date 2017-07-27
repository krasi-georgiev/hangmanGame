package main

import (
	"errors"

	"github.com/krasi-georgiev/hangmanGame/api"
)

func guessLetter(client api.HangmanClient, id int32, l string) (*api.GuessReply, error) {
	ctx, cancel := appContext()
	defer cancel()

	if id > 0 {
		guess, err := client.GuessLetter(ctx, &api.GuessRequest{GallowID: id, Letter: l})
		if err != nil {
			return nil, err
		}
		return guess, nil
	}
	return nil, errors.New("Invalid Game ID")
}
