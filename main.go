package main

import (
	"LiveboxTerminal/cmd"
	"github.com/kataras/golog"
	"log"
)

func main() {
	golog.SetLevel("debug")

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
