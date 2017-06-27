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
		Func: func(args []string) {
			fmt.Println("I do nothing...")
		},
		SubCommands: []command.Command{
			command.Command{
				Name: "login",
				Help: "github login; access token to github",
				Func: func(args []string) {
					if len(args) == 0 {
						fmt.Println("Failed login")
						return
					}
					fmt.Printf("Logged in %s", args[0])
				},
				SubCommands: []command.Command{
					command.Command{
						Name: "auth",
						Help: "github login auth; login sub command",
						Func: func(args []string) {
							if len(args) == 0 {
								fmt.Println("Failed login")
								return
							}
							fmt.Printf("Logged in with auth! %s\n", args[0])
						},
						SubCommands: []command.Command{
							command.Command{
								Name: "sub",
								Help: "github login auth sub; login sub-sub command",
								Func: func(args []string) {
									if len(args) == 0 {
										fmt.Println("Failed login")
										return
									}
									fmt.Printf("Logged in with auth! %s\n", args[0])
								},
							},
						},
					},
				},
			},
			command.Command{
				Name: "logout",
				Help: "github logout; allows you to logout from github",
				Func: func(args []string) {
					fmt.Println("Logged out")
				},
			},
		},
	})
	c.AddCommand(command.Command{
		Name: "sql",
		Help: "sql primary command interface",
		Func: func(args []string) {
			fmt.Println("I do nothing...")
		}})
	c.Run()
}
