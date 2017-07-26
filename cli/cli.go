package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/chzyer/readline"
)

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

	// log.SetOutput(l.Stderr())

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

		switch {

		case line == "1" || line == "1 - new game":
			r, err := newGallow()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(r)
			fmt.Println("Game on!")
			l.SetPrompt("Enter letter: ")
			for {
				letter, err := l.Readline()
				if err != readline.ErrInterrupt {
					if len(letter) != 0 {
						guessLetter()
					}
					continue
				} else {
					l.SetPrompt("»")
					usage(l.Stdout())
					break
				}
			}

		case line == "2" || line == "2 - saved games":
			listGallows()
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
