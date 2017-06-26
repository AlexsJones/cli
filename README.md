# cli

This is a simple interactive prompt for go that actually supports sub-commands, because I couldn't find one that did...
That said, it only supports one level of recursion at the moment, because that's all I need.

# Installation

```
go get github.com/AlexsJones/cli
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
