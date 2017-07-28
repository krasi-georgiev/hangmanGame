package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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

	menu:
		switch {
		case line == "1" || line == "1 - new game":
			r, err := newGallow(clt)
			if err != nil {
				log.Println(err)
				break
			}
			fmt.Printf(">>Game on!<<  the word is %v characters \n", len(r.WordMasked))
			l.SetPrompt("(CTRL+C to main menu) Enter letter: ")
			for {
				letter, err := l.Readline()
				if err != readline.ErrInterrupt {
					if err := guessLetter(clt, r.Id, letter); err != nil {
						fmt.Println(err)

						if er := saveGallow(clt, r.Id); er != nil {
							fmt.Println(err)
						}

						l.SetPrompt("»")
						usage(l.Stdout())
						break menu
					}
					continue
				} else {
					if err := saveGallow(clt, r.Id); err != nil {
						fmt.Println(err)
					}
					l.SetPrompt("»")
					usage(l.Stdout())
					break menu
				}
			}

		case line == "2" || line == "2 - saved games":
			r, err := listGallows(clt)
			if err != nil {
				log.Println(err)
				break
			}
			if r != nil {
				fmt.Println("ID	Status		Attempts Left	Hint")
				for _, v := range r {
					status := "      "
					if v.Status {
						status = "in progress"
					}
					fmt.Println(v.Id, "	", status, "	", v.RetryLeft, "		", v.WordMasked)
				}
			} else {
				fmt.Println("No saved games on the server!")
				break menu
			}
			l.SetPrompt("Enter game ID to resume: ")
			for {
				gameID, err := l.Readline()
				if err != readline.ErrInterrupt {
					if len(gameID) != 0 {
						i, _ := strconv.Atoi(gameID)
						if err != nil {
							log.Println("Invalid Game ID")
							continue
						} else {
							r, err := resumeGallow(clt, i)
							if err != nil {
								log.Println(err)
								continue
							}

							fmt.Printf(">>Game on!<<  word hint is : %v \n", r.WordMasked)
							l.SetPrompt("(CTRL+C to main menu) Enter letter: ")
							for {
								letter, err := l.Readline()
								if err != readline.ErrInterrupt {
									if err := guessLetter(clt, r.Id, letter); err != nil {
										fmt.Println(err)

										if er := saveGallow(clt, r.Id); er != nil {
											fmt.Println(err)
										}

										l.SetPrompt("»")
										usage(l.Stdout())
										break menu
									}
									continue
								} else {
									if err := saveGallow(clt, r.Id); err != nil {
										fmt.Println(err)
									}
									l.SetPrompt("»")
									usage(l.Stdout())
									break menu
								}
							}

						}
					}
					continue
				} else {
					l.SetPrompt("»")
					usage(l.Stdout())
					break
				}
			}
		case line == "3" || line == "3 - exit":
			os.Exit(1)
		default:
			usage(l.Stdout())
		}
	}
}
