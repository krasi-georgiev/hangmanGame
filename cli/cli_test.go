package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"testing"

	"github.com/krasi-georgiev/hangmanGame/api"
)

var clt api.HangmanClient
var gallow *api.Gallow

func TestMain(m *testing.M) {
	cmd := exec.Command("go", "build", "../srv/srv.go", "../srv/dict.go")
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

	pwd, err := os.Getwd()
	var path string
	if err == nil {
		path = pwd + "/srv"
	} else {
		path = "./srv"
	}
	cmd = exec.Command(path)
	if err := cmd.Start(); err != nil {
		cmd.Wait()
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
	fmt.Printf("Started hangman server with PID: %v \n", cmd.Process.Pid)

	clt, err = getGRPCConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	gallow, err = newGallow(clt)
	result := m.Run()

	if cmd.Process != nil {
		if err == nil {
			if err := cmd.Process.Kill(); err != nil {
				fmt.Printf("Can't kill the hangman server, error:%v \n", err)
			} else {
				fmt.Printf("Killed hangman server with PID: %v \n", cmd.Process.Pid)
			}
		}
		// cmd.Wait()

	}
	if err := os.Remove(path); err != nil {
		fmt.Println(err)
	}

	os.Exit(result)
}

func TestNewGallow(t *testing.T) {
	if gallow.Id != 1 {
		t.Logf("Gallow initialization expected ID:%v, actual ID:%v", 1, gallow.Id)
		t.Fail()
	}
}

func TestListGallows(t *testing.T) {
	if _, err := listGallows(clt); err != nil {
		t.Logf("Gallow listing error:%v", err)
		t.Fail()
	}
}

func TestResumeGallow(t *testing.T) {
	if _, err := resumeGallow(clt, strconv.Itoa(int(gallow.Id))); err == nil {
		t.Log("Gallow resume didn't fail for a locked game")
		t.Fail()
	}
	if err := saveGallow(clt, gallow); err != nil {
		t.Logf("Gallow save error:%v", err)
		t.Fail()
	}
	if _, err := resumeGallow(clt, "-1"); err == nil {
		t.Log("Galows resume didn't  fail with an invalid ID")
		t.Fail()
	}
	g, err := resumeGallow(clt, strconv.Itoa(int(gallow.Id)))
	if err != nil {
		t.Logf("Gallow resume error:%v", err)
		t.FailNow()
	}
	if g.Id != gallow.Id {
		t.Logf("Galows id expected:%v, actual:%v", gallow.Id, g.Id)
		t.Fail()
	}
}

func TestGuesslLetter(t *testing.T) {

}

func TestGamePlay(t *testing.T) {

}
