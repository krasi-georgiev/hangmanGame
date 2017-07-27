package main

import (
	"errors"
	"log"
	"math/rand"
	"net"
	"strings"
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
	id := int32(len(s.slaughter) + 1)
	wordID := rand.Intn(int(len(words)))
	word := words[wordID]
	wordMAsked := strings.Repeat("_", utf8.RuneCountInString(word))
	s.slaughter = append(s.slaughter, &api.Gallow{Id: id, Word: word, WordMasked: wordMAsked, RetryLimit: r.RetryLimit, RetryLeft: r.RetryLimit})

	g := s.slaughter[id-1 : id]
	g[0].Word = "" // don't sent the naked word to the client , only the server knows the word
	return &api.GallowReply{Gallow: g}, nil
}

func (s *hangman) ListGallows(context.Context, *api.GallowRequest) (*api.GallowReply, error) {
	d := &api.GallowReply{Gallow: s.slaughter}
	return d, nil
}
func (s *hangman) ResumeGallow(context.Context, *api.GallowRequest) (*api.GallowReply, error) {
	return nil, nil
}
func (s *hangman) GuessLetter(context.Context, *api.GuessRequest) (*api.GuessReply, error) {
	return nil, nil
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
