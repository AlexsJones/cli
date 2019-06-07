# cli

[![Build Status](https://travis-ci.org/AlexsJones/cli.svg?branch=master)](https://travis-ci.org/AlexsJones/cli)

[![GoDoc](https://godoc.org/github.com/AlexsJones/cli/cli?status.svg)](https://godoc.org/github.com/AlexsJones/cli/cli)

[![Maintainability](https://api.codeclimate.com/v1/badges/3a06871c361d5e8e70ae/maintainability)](https://codeclimate.com/github/AlexsJones/cli/maintainability)

This is a simple interactive prompt for go that actually supports sub-commands, because I couldn't find one that did...
Supports unlimited subcommand nesting.

It looks a bit like this (Once you wire up your commands: see example):
```
>>>github login auth alex
Hit auth
Authenticated with alex

>>>github logout
Logged out

>>>help
npm sub commands:
	[npm] file: relink an npm package locally<prefix> <string>
	[npm] remove: remove a dep from package.json <string>
	[npm] usage: find usage of a package within submodules <string>
github sub commands:
	[github] pr: pr command palette
		[pr] attach: attach the current issue to a pr <reponame> <owner> <prnumber>
	[github] issue: Issue command palette
		[issue] set: set the current working issue <issue url>
		[issue] unset: unset the current working issue
		[issue] show: show the current working issue
	[github] login: use an access token to login to github
submodule sub commands:
	[submodule] exec: execute in all submodules <command string>

```


# Installation

```
go get github.com/AlexsJones/cli/cli
```
# Simple example

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
	})

	c.Run()

}
```

This gives you something like:

```
>>>github
I do nothing...

```


# Recursive subcommand example

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

```
>>>help
npm sub commands:
	[npm] file: relink an npm package locally<prefix> <string>
	[npm] remove: remove a dep from package.json <string>
	[npm] usage: find usage of a package within submodules <string>
github sub commands:
	[github] pr: pr command palette
		[pr] attach: attach the current issue to a pr <reponame> <owner> <prnumber>
	[github] issue: Issue command palette
		[issue] set: set the current working issue <issue url>
		[issue] unset: unset the current working issue
		[issue] show: show the current working issue
	[github] login: use an access token to login to github
submodule sub commands:
	[submodule] exec: execute in all submodules <command string>


```
