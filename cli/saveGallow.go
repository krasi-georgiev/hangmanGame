package main

import (
	"github.com/krasi-georgiev/hangmanGame/api"
)

func saveGallow(client api.HangmanClient, id int32) error {
	ctx, cancel := appContext()
	defer cancel()
	_, err := client.SaveGallow(ctx, &api.GallowRequest{Id: id})
	if err != nil {
		return err
	}
	return nil
}
