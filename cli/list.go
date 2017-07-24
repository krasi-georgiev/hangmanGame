package main

import (
	"fmt"

	"github.com/krasi-georgiev/hangmanGame/api"
	"github.com/urfave/cli"
)

var (
	hangmanListCommand = cli.Command{
		Name:    "list",
		Usage:   "list all games on the server",
		Aliases: []string{"ls"},
		Action:  hangmanListFn,
	}
)

func hangmanListFn(context *cli.Context) error {
	var (
		ctx, cancel = appContext(context)
	)
	defer cancel()

	client, err := getGRPCConnection(context)
	if err != nil {
		return err
	}
	games, err := client.List(ctx, &api.HangmanRequest{})
	if err != nil {
		return err
	}

	fmt.Println("sdfsdf")
	fmt.Println(games)

	return nil
}
