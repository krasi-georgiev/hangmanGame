package main

import (
	"fmt"

	"github.com/krasi-georgiev/hangmanGame/api"
)

func guessLetter() error {
	ctx, cancel := appContext()
	defer cancel()

	client, err := getGRPCConnection(&ctx)
	if err != nil {
		return err
	}
	guess, err := client.GuessLetter(ctx, &api.GuessRequest{})
	if err != nil {
		return err
	}

	fmt.Println("sdfsdf")
	fmt.Println(guess)

	return nil
}
