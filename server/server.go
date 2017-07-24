package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"

	"github.com/krasi-georgiev/hangmanGame/api"
	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) New(context.Context, *api.HangmanRequest) (*api.HangmanReply, error) {
	return nil, nil
}
func (s *server) List(context.Context, *api.HangmanRequest) (*api.HangmanReply, error) {
	fmt.Println("LIST")
	return nil, nil
}
func (s *server) Resume(context.Context, *api.HangmanRequest) (*api.HangmanReply, error) {
	return nil, nil
}
func (s *server) Guess(context.Context, *api.GuessRequest) (*api.GuessReply, error) {
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterHangmanServer(s, &server{})
	fmt.Println("listening")
	s.Serve(lis)
}
