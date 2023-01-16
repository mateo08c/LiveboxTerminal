package main

import (
	"bufio"
	"fmt"
	"github.com/kataras/golog"
	"os"
)

type Command struct {
	Name          string
	Description   string
	Handler       func(c *Command, cmd *Command)
	SubCommands   []*Command
	BeforeCommand *Command
}

var Commands = &Command{
	Name:        "LiveboxTerminal",
	Description: "main",
	Handler:     nil,
}

func StartTerminal() {
	Commands.AddSubCommand("exit", "Quitter le terminal", ExitTerminal)
	Commands.AddSubCommand("ipv6", "Entrer dans la configuration IPv6", Ipv6)

	Commands.Read(nil)
}

func (cmd *Command) Read(parent *Command) {
	cmd.BeforeCommand = parent
	name := cmd.Name
	PrintCursor(name)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		if text == "help" {
			for _, d := range cmd.SubCommands {
				golog.Info(d.Name, " - ", d.Description)
			}
		}

		if text == "" {
			PrintCursor(name)
			continue
		}

		for _, d := range cmd.SubCommands {
			if d.Name == text {
				d.Execute(cmd)
				break
			}
		}

		PrintCursor(name)
	}
}

func PrintCursor(name string) {
	fmt.Print("\u001B[1;32m" + name + " > \u001B[0m")
}

func (cmd *Command) Execute(main *Command) {
	cmd.Handler(main, cmd)
}

func (cmd *Command) AddSubCommand(name string, description string, handler func(*Command, *Command)) {
	cmd.SubCommands = append(cmd.SubCommands, &Command{name, description, handler, nil, nil})
}
