package main

import (
	"fmt"

	"github.com/krasi-georgiev/hangmanGame/api"
)

func guessLetter(client api.HangmanClient) error {
	ctx, cancel := appContext()
	defer cancel()
	guess, err := client.GuessLetter(ctx, &api.GuessRequest{})
	if err != nil {
		return err
	}

	fmt.Println("sdfsdf")
	fmt.Println(guess)

	return nil
}
