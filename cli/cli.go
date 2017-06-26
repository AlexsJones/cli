package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/AlexsJones/cli/command"
	"github.com/olekukonko/tablewriter"
)

//Cli control object
type Cli struct {
	Commands       []command.Command
	Unknowncommand func(args []string)
}

//NewCli initialize
func NewCli() *Cli {
	c := &Cli{}
	c.Unknowncommand = nil
	return c
}

//AddCommand to Cli
func (cli *Cli) AddCommand(c command.Command) {
	cli.Commands = append(cli.Commands, c)
}

func (cli *Cli) printHelp() {
	data := [][]string{}

	for _, commands := range cli.Commands {
		for _, subCommands := range commands.SubCommands {
			data = append(data, []string{commands.Name, subCommands.Name, subCommands.Help})
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Module", "Name", "Help"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}

func (cli *Cli) parseSystemCommands(input []string) *command.Command {
	if input[0] == "exit" {
		fmt.Println("Bye")
		os.Exit(0)
	}
	if input[0] == "help" {
		return &command.Command{
			Func: func(arg []string) {
				cli.printHelp()
			},
		}
	}
	return nil
}

func (cli *Cli) findCommand(input string) (*command.Command, []string) {
	parsed := strings.Fields(input)
	if len(parsed) == 0 {
		fmt.Println("No input detected")
		return nil, parsed[1:]
	}
	if systemCmd := cli.parseSystemCommands(parsed); systemCmd != nil {
		return systemCmd, parsed[1:]
	}

	currentCommands := cli.Commands
	//Maybe recurse this one day...
	for _, primary := range currentCommands {
		if parsed[0] == primary.Name {
			for _, secondary := range primary.SubCommands {
				if len(parsed) > 1 {
					if parsed[1] == secondary.Name {
						return &secondary, parsed[2:]
					}
				}
			}
			return &primary, parsed[1:]
		}
	}
	return nil, nil
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

	command, args := cli.findCommand(text)
	if command != nil {
		if command.Func != nil {
			command.Func(args)
			fmt.Printf("\n")
		}
	} else {
		if cli.Unknowncommand != nil {
			cli.Unknowncommand(args)
		}
	}
	goto reset
}
