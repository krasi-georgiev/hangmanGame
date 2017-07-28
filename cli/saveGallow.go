package main

import (
	"github.com/krasi-georgiev/hangmanGame/api"
)

func saveGallow(client api.HangmanClient, g *api.Gallow) error {
	ctx, cancel := appContext()
	defer cancel()
	_, err := client.SaveGallow(ctx, &api.GallowRequest{Id: g.Id})
	if err != nil {
		return err
	}
	return nil
}
