package main

import (
	"errors"
	"fmt"
	"log"
	"net"

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
	s.slaughter = append(s.slaughter, &api.Gallow{Id: id, RetryLimit: r.RetryLimit, RetryLeft: r.RetryLimit})
	return &api.GallowReply{Gallow: s.slaughter[id-1 : id]}, nil
}
func (s *hangman) ListGallows(context.Context, *api.GallowRequest) (*api.GallowReply, error) {
	fmt.Println("LIST")
	return nil, nil
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
