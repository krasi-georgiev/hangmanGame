package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/chzyer/readline"
)

var currentGame int32

func main() {
	l, err := readline.NewEx(&readline.Config{
		Prompt:          "»",
		InterruptPrompt: "^C",
		AutoComplete:    completer,
		EOFPrompt:       "exit",
	})
	if err != nil {
		panic(err)
	}
	defer l.Close()

	clt, err := getGRPCConnection()
	if err != nil {
		log.Fatal(err)
	}

	usage(l.Stdout())

	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		// fmt.Print("\033[H\033[2J") // clear the screen
	menu:
		switch {
		case line == "1" || line == "1 - new game":
			r, err := newGallow(clt)
			currentGame = r.Id
			if err != nil {
				log.Println(err)
				break
			}
			fmt.Printf(">>Game on!<<  the word is %v characters \n", len(r.WordMasked))
			l.SetPrompt("(CTRL+C to main menu) Enter letter: ")
			for {
				letter, err := l.Readline()
				if err != readline.ErrInterrupt {
					if len(letter) == 1 {
						g, err := guessLetter(clt, currentGame, letter)
						if err != nil {
							fmt.Println(err)
							continue
						}

						if g.Gallow.RetryLeft < 1 {
							fmt.Print("\n>>>> Game Over Amigo , try another day! <<<<\n\n\n")
							l.SetPrompt("»")
							usage(l.Stdout())
							break menu
						}

						fmt.Println(gallowArt[(len(gallowArt) - int(g.Gallow.RetryLeft))])
						fmt.Printf("Remaining attempts: %v \n", g.Gallow.RetryLeft)
						fmt.Printf("Incorrect attempts: ")
						for _, v := range g.Gallow.IncorrectGuesses {
							fmt.Print(v.Letter, " ")
						}

						fmt.Println("\nWord hint:", g.Gallow.WordMasked)
					} else {
						fmt.Println("Please provide a single letter")
					}
					continue
				} else {
					l.SetPrompt("»")
					usage(l.Stdout())
					break
				}
			}

		case line == "2" || line == "2 - saved games":
			r, err := listGallows(clt)
			if err != nil {
				log.Println(err)
				break
			}

			fmt.Println(r)
			l.SetPrompt("Enter game ID to continue playing: ")
			for {
				gameID, err := l.Readline()
				if err != readline.ErrInterrupt {
					if len(gameID) != 0 {

					}
					continue
				}
			}

		case line == "3" || line == "3 - exit":
			os.Exit(1)
		default:
			usage(l.Stdout())
		}
	}
}

func usage(w io.Writer) {
	io.WriteString(w, completer.Tree("    "))
}

var completer = readline.NewPrefixCompleter(
	readline.PcItem("1 - new game"),
	readline.PcItem("2 - saved games"),
	readline.PcItem("3 - exit"),
)
