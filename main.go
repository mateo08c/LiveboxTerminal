package main

import (
	"fmt"
	"github.com/kataras/golog"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

const (
	IP = "192.168.1.1"
)

var ContextID string
var Cookie string

func main() {
	golog.SetTimeFormat("")
	golog.SetPrefix("LiveboxTerminal: ")

	golog.Info("Entrée votre mot de passe (CTRL+C pour quitter)")
	PrintCursor("Mot de passe: ")

	passwd, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		golog.Fatal(err)
		return
	}
	fmt.Print("\n")

	resp, err := GetContextID("admin", string(passwd))
	if err != nil {
		golog.Fatal(err)
		return
	}

	if resp.Status != 0 {
		golog.Error("Login failed")
		return
	}

	ContextID = resp.Data.ContextID
	golog.Info("Authentification avec succès: " + ContextID)

	go StartTerminal()

	select {} // block forever
}
