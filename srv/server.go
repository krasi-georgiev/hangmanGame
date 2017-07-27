package main

import (
	"errors"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/net/context"

	"github.com/krasi-georgiev/hangmanGame/api"
	"google.golang.org/grpc"
)

type hangman struct {
	slaughter []*api.Gallow
}

func (s *hangman) NewGallow(ctx context.Context, r *api.GallowRequest) (*api.GallowReply, error) {
	if r.RetryLimit < 1 {
		return nil, errors.New("Please specify retry limit for this hangman")
	}
	// pick a random word
	rand.Seed(time.Now().UnixNano())
	wordID := rand.Intn(len(words))
	word := words[wordID]
	wordMAsked := strings.Repeat("_", utf8.RuneCountInString(word))
	gallowID := int32(len(s.slaughter) + 1) // generate an id sequence starting from 1
	s.slaughter = append(s.slaughter, &api.Gallow{Id: gallowID, Word: word, WordMasked: wordMAsked, RetryLimit: r.RetryLimit, RetryLeft: r.RetryLimit})

	g := s.slaughter[gallowID-1 : gallowID]
	d := *g[0]  // need to dereference so we don't change the original struct
	d.Word = "" // don't sent the naked word to the client , to avoid cheating clients :)
	return &api.GallowReply{Gallow: []*api.Gallow{&d}}, nil
}

func (s *hangman) ListGallows(context.Context, *api.GallowRequest) (*api.GallowReply, error) {
	d := &api.GallowReply{Gallow: s.slaughter}
	return d, nil
}
func (s *hangman) ResumeGallow(context.Context, *api.GallowRequest) (*api.GallowReply, error) {
	return nil, nil
}
func (s *hangman) GuessLetter(ctx context.Context, r *api.GuessRequest) (*api.GuessReply, error) {
	// stay in range of the slice
	if int32(len(s.slaughter)) >= r.GallowID {
		r.Letter = strings.ToLower(r.Letter)
		g := s.slaughter[r.GallowID-1]

		if g.RetryLeft < 1 {
			return nil, errors.New("This game is over")
		}

		for k, v := range g.Word { // expose all letter occurencies
			if v == rune(r.Letter[0]) {
				g.WordMasked = g.WordMasked[:k] + r.Letter + g.WordMasked[k+1:]
			}
		}
		if strings.Index(g.Word, r.Letter) == -1 {
			g.RetryLeft = g.RetryLeft - 1

			contains := false
			for _, v := range g.IncorrectGuesses {
				if r.Letter == v.Letter {
					contains = true
				}
			}
			if !contains {
				g.IncorrectGuesses = append(g.IncorrectGuesses, &api.GuessRequest{Letter: r.Letter})
			}
		}
		return &api.GuessReply{Gallow: g}, nil
	}
	return nil, errors.New("Invalid Game ID")
}

func main() {
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterHangmanServer(s, &hangman{})
	log.Println("listening!")
	s.Serve(lis)
}
