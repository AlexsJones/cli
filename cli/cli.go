package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"

	"github.com/AlexsJones/cli/command"
	"github.com/fatih/color"
)

//Cli control object
type Cli struct {
	Commands []command.Command
}

//NewCli initialize
func NewCli() *Cli {
	c := &Cli{}
	return c
}

//AddCommand to Cli
func (cli *Cli) AddCommand(c command.Command) {
	cli.Commands = append(cli.Commands, c)
}

func (cli *Cli) recurseHelp(c []command.Command, parent command.Command) {
	for _, cmd := range c {
		if parent.Name != "" {
			fmt.Printf("[parent command: %s][current command: %s]: %s\n", parent.Name, cmd.Name, cmd.Help)
		}
		if len(cmd.SubCommands) > 0 {
			cli.recurseHelp(cmd.SubCommands, cmd)
		}
	}
}

func (cli *Cli) parseSystemCommands(input []string) error {
	if input[0] == "exit" {
		fmt.Println("Bye")
		os.Exit(0)
	}
	if input[0] == "clear" {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
	}
	if input[0] == "help" {
		cli.recurseHelp(cli.Commands, command.Command{})
	}

	return nil
}

func (cli *Cli) recurse(c []command.Command, args []string, i int) error {

	for _, cmd := range c {
		if i > len(args) {
			return nil
		}
		if cmd.Name == args[i] {
			if len(cmd.SubCommands) > 0 && len(args) > i+1 {
				cli.recurse(cmd.SubCommands, args, i+1)
			} else {
				cmd.Func(args[i+1:])
				fmt.Printf("\n")
			}
		}
	}
	return nil
}
func (cli *Cli) findCommand(input string) error {
	parsed := strings.Fields(input)
	if len(parsed) == 0 {
		fmt.Println("No input detected")
		return nil
	}
	if systemCmd := cli.parseSystemCommands(parsed); systemCmd != nil {
		return nil
	}
	currentCommands := cli.Commands
	error := cli.recurse(currentCommands, parsed, 0)
	if error != nil {
		return error
	}
	return nil
}

//Run the primary entry point
func (cli *Cli) Run() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			fmt.Printf("\nBye\n")
			os.Exit(0)
		}
	}()
reset:
	//Get user input
	fmt.Print(">>>")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	err := cli.findCommand(text)
	if err != nil {
		color.Red(err.Error())
	}
	goto reset
}
