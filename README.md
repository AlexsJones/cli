# cli

This is a simple interactive prompt for go that actually supports sub-commands, because I couldn't find one that did...
That said, it only supports one level of recursion at the moment, because that's all I need.

# Installation

```
go get github.com/AlexsJones/cli/cli
```

# Example

```
package main

import (
	"fmt"

	"github.com/AlexsJones/cli/cli"
	"github.com/AlexsJones/cli/command"
)

func main() {

	c := cli.NewCli()

	c.Unknowncommand = func(args []string) {
		fmt.Println("Unknown command")
	}

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
			},
			command.Command{
				Name: "logout",
				Help: "",
				Func: func(args []string) {
					fmt.Println("Logged out")
				},
			},
		},
	})

	c.Run()
}

```

# System commands

`help` & `exit`

Gives you information such as:

```

+--------+--------+------------------------+
| MODULE |  NAME  |          HELP          |
+--------+--------+------------------------+
| github | login  | access token to github |
| github | logout |                        |
+--------+--------+------------------------+

```
