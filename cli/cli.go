package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "hangman cli"
	app.Version = "0.1"
	app.Usage = `hangman usage`

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "address, a",
			Usage: "address for the hangman server",
			Value: "localhost:9999",
		},
		cli.DurationFlag{
			Name:  "timeout",
			Usage: "total timeout for any commands",
			Value: time.Duration(2 * time.Second),
		},
		cli.DurationFlag{
			Name:  "connect-timeout",
			Usage: "timeout for connecting to the hangman server",
			Value: time.Duration(2 * time.Second),
		},
	}

	app.Commands = append([]cli.Command{
		hangmanListCommand,
	})

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "ctr: %s\n", err)
		os.Exit(1)
	}
}
