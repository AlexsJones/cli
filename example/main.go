package main

import (
	"fmt"

	"github.com/AlexsJones/cli/cli"
	"github.com/AlexsJones/cli/command"
)

func main() {

	c := cli.NewCli()

	c.AddCommand(command.Command{
		Name: "github",
		Help: "github primary command interface",
		Func: func() {
			fmt.Println("I do nothing...")
		},
		SubCommands: []command.Command{
			command.Command{
				Name: "login",
				Help: "access token to github",
				Func: func() {
					fmt.Println("Logged in")
				},
			},
			command.Command{
				Name: "logout",
				Help: "",
				Func: func() {
					fmt.Println("Logged out")
				},
			},
		},
	})

	c.Run()
}
