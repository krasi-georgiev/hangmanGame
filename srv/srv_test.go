package main

import (
	"context"
	"log"
	"testing"

	"github.com/krasi-georgiev/hangmanGame/api"
)

var g *api.Gallow
var h *hangman

func init() {

	var err error
	h = &hangman{}
	if _, err = h.NewGallow(context.Background(), &api.GallowRequest{Id: -1}); err == nil {
		log.Panic("Gallow initialization without Retry limit should fail")
	}
	g, err = h.NewGallow(context.Background(), &api.GallowRequest{RetryLimit: 5})
	if err != nil {
		log.Panic("Gallow initialization returned an error:", err)
	}
}

func TestNewGallow(t *testing.T) {
	if g.Id != 1 {
		t.Logf("Gallow initialization expected ID:%v, actual ID:%v", 1, g.Id)
		t.Fail()
	}
}

func TestListGallows(t *testing.T) {
	l, err := h.ListGallows(context.Background(), &api.GallowRequest{Id: -1})
	if err != nil {
		t.Log("Gallow listing error:", err)
		t.Fail()
	}
	if len(l.Gallow) == 0 {
		t.Log("Gallow listing returned unexpected 0 length")
		t.Fail()
	}
	if l.Gallow[0].Id != 1 {
		t.Logf("Gallow listing returned unexpected Id:%v for the first element", l.Gallow[0].Id)
		t.Fail()
	}
}

func TestResumeGallow(t *testing.T) {
	if _, err := h.ResumeGallow(context.Background(), &api.GallowRequest{Id: 1}); err == nil {
		t.Log("Gallow resume should fail for locked gallows")
		t.Fail()
	}
	if _, err := h.SaveGallow(context.Background(), &api.GallowRequest{Id: 1}); err != nil {
		t.Logf("Gallow save error:%v", err)
		t.Fail()
	}

	g1, err := h.ResumeGallow(context.Background(), &api.GallowRequest{Id: 1})
	if err != nil {
		t.Logf("Gallow resume error:%v", err)
		t.Fail()
	}

	if g1.Id != 1 {
		t.Logf("Gallow ID expected:%v, actual:%v", 1, g1.Id)
		t.Fail()
	}
	if _, err := h.ResumeGallow(context.Background(), &api.GallowRequest{Id: -1}); err == nil {
		t.Log("Gallow didn't fail with an invalid Gallow ID")
		t.Fail()
	}
}

func TestGuesslLetter(t *testing.T) {
	g, err := h.SaveGallow(context.Background(), &api.GallowRequest{Id: 1})
	if err != nil {
		t.Logf("Gallow save error:%v", err)
		t.Fail()

	}
	gg, err := h.GuessLetter(context.Background(), &api.GuessRequest{GallowID: 1, Letter: "~"})
	if err != nil {
		t.Logf("Gallow letter guess error:%v", err)
		t.Fail()
	}
	if g.RetryLeft-gg.RetryLeft != 1 {
		t.Logf("Retry Limit decrease expected:1 actual:%v", (g.RetryLeft - gg.RetryLeft))
		t.Fail()
	}
	if _, err := h.GuessLetter(context.Background(), &api.GuessRequest{GallowID: -1, Letter: "~"}); err == nil {
		t.Log("Letter guess didn't fail with an invalid Gallow ID")
		t.Fail()
	}
}
