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

func (s *hangman) NewGallow(ctx context.Context, r *api.GallowRequest) (*api.Gallow, error) {
	if r.RetryLimit < 1 {
		return nil, errors.New("Please specify retry limit for this hangman")
	}
	// pick a random word
	rand.Seed(time.Now().UnixNano())
	wordID := rand.Intn(len(words))
	word := words[wordID]
	wordMAsked := strings.Repeat("_", utf8.RuneCountInString(word))
	gallowID := int32(len(s.slaughter)) // generate an id sequence

	if gallowID == 0 { // add one  gallow to fill the 0 slice element so our actual gallowID matches the slice positions
		s.slaughter = append(s.slaughter, &api.Gallow{Id: 0, Status: true})
		gallowID++
	}
	s.slaughter = append(s.slaughter, &api.Gallow{Id: gallowID, Word: word, WordMasked: wordMAsked, RetryLimit: r.RetryLimit, RetryLeft: r.RetryLimit, Status: true})
	g := *s.slaughter[gallowID] // need to dereference so we don't change the original struct
	g.Word = ""                 // don't sent the naked word to the client , to avoid cheating clients :)
	return &g, nil
}

func (s *hangman) ListGallows(context.Context, *api.GallowRequest) (*api.GallowArray, error) {
	d := &api.GallowArray{Gallow: s.slaughter}
	if d.Gallow == nil {
		return &api.GallowArray{}, nil
	}
	d.Gallow = d.Gallow[1:] // don't neet the 0 element as it is only a fake filling
	return d, nil
}

func (s *hangman) ResumeGallow(ctx context.Context, r *api.GallowRequest) (*api.Gallow, error) {
	// stay in range of the slice
	if r.Id > 0 && int32(len(s.slaughter)) > r.Id {
		if s.slaughter[r.Id].RetryLeft < 1 {
			return nil, errors.New("This game is over")
		}
		if s.slaughter[r.Id].Status {
			return nil, errors.New("Game is played by someone else")
		}

		s.slaughter[r.Id].Status = true // lock the game
		d := *s.slaughter[r.Id]         // need to dereference so we don't change the original struct
		d.Word = ""                     // don't sent the naked word to the client , to avoid cheating clients :)
		return &d, nil
	}
	return nil, errors.New("Invalid Game ID")
}

func (s *hangman) SaveGallow(ctx context.Context, r *api.GallowRequest) (*api.Gallow, error) {
	// stay in range of the slice
	if r.Id > 0 && int32(len(s.slaughter)) > r.Id {
		s.slaughter[r.Id].Status = false
		gg := *s.slaughter[r.Id] // need to dereference so we don't change the original struct
		gg.Word = ""             // don't sent the naked word to the client , to avoid cheating clients :)
		return &gg, nil
	}
	return nil, errors.New("Invalid Game ID")
}

func (s *hangman) GuessLetter(ctx context.Context, r *api.GuessRequest) (*api.Gallow, error) {
	// stay in range of the slice
	if r.GallowID > 0 && int32(len(s.slaughter)) > r.GallowID {
		r.Letter = strings.ToLower(r.Letter)
		g := s.slaughter[r.GallowID]
		if g.RetryLeft < 1 {
			return nil, errors.New("This game is over")
		}

		for k, v := range g.Word { // expose all letter occurencies
			if v == rune(r.Letter[0]) {
				g.WordMasked = g.WordMasked[:k] + r.Letter + g.WordMasked[k+1:]
			}
		}
		if strings.Index(g.Word, r.Letter) == -1 {
			contains := false
			for _, v := range g.IncorrectGuesses {
				if r.Letter == v.Letter {
					contains = true
				} else {

				}
			}
			if !contains {
				g.IncorrectGuesses = append(g.IncorrectGuesses, &api.GuessRequest{Letter: r.Letter})
				g.RetryLeft = g.RetryLeft - 1
			}
		}
		gg := *g     // need to dereference so we don't change the original struct
		gg.Word = "" // don't sent the naked word to the client , to avoid cheating clients :)
		return &gg, nil
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
