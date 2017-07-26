package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"

	"github.com/krasi-georgiev/hangmanGame/api"
	"google.golang.org/grpc"
)

type hangman struct {
	slaughterhouse []api.Gallow
}

func (s *hangman) NewGallow(context.Context, *api.HangmanRequest) (*api.HangmanReply, error) {
	s.slaughterhouse = append(s.slaughterhouse, api.Gallow{Id: int32(len(s.slaughterhouse) + 1), Status: true})
	for k, v := range s.slaughterhouse {
		log.Println(k)
		log.Println(v)
	}
	return nil, nil
}
func (s *hangman) ListGallows(context.Context, *api.HangmanRequest) (*api.HangmanReply, error) {
	fmt.Println("LIST")
	return nil, nil
}
func (s *hangman) ResumeGallow(context.Context, *api.HangmanRequest) (*api.HangmanReply, error) {
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
