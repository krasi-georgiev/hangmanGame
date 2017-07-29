package main

import (
	"errors"
	"fmt"

	"github.com/krasi-georgiev/hangmanGame/api"
)

func listGallows(client api.HangmanClient) (string, error) {
	var reply string

	ctx, cancel := appContext()
	defer cancel()

	r, err := client.ListGallows(ctx, &api.GallowRequest{Id: -1})
	if err != nil {
		return "", err
	}

	if len(r.Gallow) > 0 {
		reply += "ID	Status		Attempts Left	Hint \n"
		for _, v := range r.Gallow {
			status := "          "
			if v.Status {
				status = "in progress"
			}
			reply += fmt.Sprint(v.Id, "	", status, "	", v.RetryLeft, "		", v.WordMasked, "\n")
		}
	} else {
		return "", errors.New("No saved games on the server!")
	}
	return reply, nil
}
