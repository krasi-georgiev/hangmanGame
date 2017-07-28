package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/krasi-georgiev/hangmanGame/api"
)

func guessLetter(client api.HangmanClient, id int32, l string) error {
	if len(l) != 1 {
		fmt.Println("Please provide a single letter")
		return nil
	}
	ctx, cancel := appContext()
	defer cancel()

	if id > 0 {
		g, err := client.GuessLetter(ctx, &api.GuessRequest{GallowID: id, Letter: l})
		if err != nil {
			return err
		}

		if g.RetryLeft < 1 {
			return errors.New("\n>>>> Game Over Amigo , try another day <<<<\n\n\n")
		}

		if strings.Index(g.WordMasked, "_") == -1 {
			return fmt.Errorf("\n>>>> Well done, you guessed the word:%v  <<<<\n\n\n", g.WordMasked)
		}

		fmt.Println(gallowArt[(len(gallowArt) - int(g.RetryLeft) - 1)])
		fmt.Printf("Remaining attempts: %v \n", g.RetryLeft)
		fmt.Printf("Incorrect attempts: ")
		for _, v := range g.IncorrectGuesses {
			fmt.Print(v.Letter, " ")
		}
		fmt.Println("\nWord hint:", g.WordMasked)
	} else {
		return errors.New("Invalid Game ID")
	}
	return nil
}
