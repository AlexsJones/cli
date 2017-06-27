# cli

[![Build Status](https://travis-ci.org/AlexsJones/cli.svg?branch=master)](https://travis-ci.org/AlexsJones/cli)

This is a simple interactive prompt for go that actually supports sub-commands, because I couldn't find one that did...
Supports unlimited subcommand nesting.

It looks a bit like this:
```
>>>github login auth alex
Hit auth
Authenticated with alex

>>>github logout
Logged out
Failed logout

>>>help
github sub commands:
	[github] login: access token to github
		[login] auth: login sub command
			[auth] sub: login sub-sub command
	[github] logout: allows you to logout from github
```


# Installation

```
go get github.com/AlexsJones/cli/cli
```

# Example

```go
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
				Help: "access token to github",
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
						Help: "login sub command",
						Func: func(args []string) {
							fmt.Println("Hit auth")
							if len(args) == 0 {
								fmt.Println("Failed login")
								return
							}
							fmt.Printf("Authenticated with %s\n", args[0])
						},
						SubCommands: []command.Command{
							command.Command{
								Name: "sub",
								Help: "login sub-sub command",
								Func: func(args []string) {
									if len(args) == 0 {
										fmt.Println("Failed login")
										return
									}
									fmt.Printf("Logged in with username %s\n", args[0])
								},
							},
						},
					},
				},
			},
			command.Command{
				Name: "logout",
				Help: "allows you to logout from github",
				Func: func(args []string) {
					fmt.Println("Logged out")
					if len(args) == 0 {
						fmt.Println("Failed logout")
						return
					}
					fmt.Printf("Logged out with username %s\n", args[0])
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
```

# System commands

`help` & `exit`

Gives you information such as:

*Needs to be improved and recurse in a better way*

```
auth: github login auth; login sub command
login: github login; access token to github
logout: github logout; allows you to logout from github
github: github primary command interface

```
