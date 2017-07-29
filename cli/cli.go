package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/chzyer/readline"
)

var completer = readline.NewPrefixCompleter(
	readline.PcItem("1 - new game"),
	readline.PcItem("2 - saved games"),
	readline.PcItem("3 - exit"),
)

func main() {
	l, err := readline.NewEx(&readline.Config{
		Prompt:          "»",
		InterruptPrompt: "^C",
		AutoComplete:    completer,
		EOFPrompt:       "exit",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	clt, err := getGRPCConnection()
	if err != nil {
		log.Fatal(err)
	}

	usage(l)

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
		menu := line[0:1] // to simplify the select statment just take the first character of the user input

	menu:
		switch {
		case menu == "1":
			g, err := newGallow(clt)
			if err != nil {
				fmt.Println(err)
				break menu
			}
			fmt.Printf(">>Game on!<<  the word hint is %v \n", g.WordMasked)
			l.SetPrompt("(CTRL+C to main menu) Enter letter: ")
			for {
				letter, err := l.Readline()
				if err != readline.ErrInterrupt {
					r, err := guessLetter(clt, g, letter)
					if err != nil {
						fmt.Println(err)

						if er := saveGallow(clt, g); er != nil {
							fmt.Println(err)
						}
						usage(l)
						break menu
					}
					fmt.Println(r)
					continue
				} else {
					if err := saveGallow(clt, g); err != nil {
						fmt.Println(err)
					}
					usage(l)
					break menu
				}
			}

		case menu == "2":
			r, err := listGallows(clt)
			if err != nil {
				fmt.Println(err)
				break menu
			}
			fmt.Println(r)

			l.SetPrompt("(CTRL+C to main menu) Enter game ID to resume: ")
			for {
				gameID, err := l.Readline()
				if err != readline.ErrInterrupt {
					r, err := resumeGallow(clt, gameID)
					if err != nil {
						log.Println(err)
						continue
					}

					fmt.Printf(">>Game on!<<  word hint is : %v \n", r.WordMasked)
					l.SetPrompt("(CTRL+C to main menu) Enter letter: ")
					for {
						letter, err := l.Readline()
						if err != readline.ErrInterrupt {
							rr, err := guessLetter(clt, r, letter)
							if err != nil {
								fmt.Println(err)
								if er := saveGallow(clt, r); er != nil {
									fmt.Println(err)
								}

								usage(l)
								break menu
							}
							fmt.Println(rr)
							continue
						} else {
							if err := saveGallow(clt, r); err != nil {
								fmt.Println(err)
							}
							usage(l)
							break menu
						}
					}
				} else {
					usage(l)
					break
				}
			}
		case menu == "3":
			os.Exit(1)
		default:
			usage(l)
		}
	}
}
