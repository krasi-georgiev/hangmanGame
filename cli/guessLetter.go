package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/krasi-georgiev/hangmanGame/api"
)

func guessLetter(client api.HangmanClient, g *api.Gallow, l string) (string, error) {

	var reply string

	if len(l) != 1 {
		return "Please provide a single letter", nil
	}
	ctx, cancel := appContext()
	defer cancel()

	if g.Id > 0 {
		gg, err := client.GuessLetter(ctx, &api.GuessRequest{GallowID: g.Id, Letter: l})
		if err != nil {
			return "", err
		}

		if gg.RetryLeft < 1 {
			return "", errors.New("\n>>>> Game Over Amigo , try another day <<<<\n\n\n")
		}

		if strings.Index(gg.WordMasked, "_") == -1 {
			return "", fmt.Errorf("\n>>>> Well done, you guessed the word:%v  <<<<\n\n\n", gg.WordMasked)
		}

		reply += gallowArt[(len(gallowArt) - int(gg.RetryLeft) - 1)]
		reply += fmt.Sprintf("\nRemaining attempts: %v", gg.RetryLeft)
		reply += ("\nIncorrect attempts: ")
		for _, v := range gg.IncorrectGuesses {
			reply += fmt.Sprint(v.Letter, " ")
		}
		reply += fmt.Sprint("\nWord hint:", gg.WordMasked)
	} else {
		return "", errors.New("Invalid Game ID")
	}
	return reply, nil
}
