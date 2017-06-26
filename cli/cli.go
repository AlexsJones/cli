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
	Commands []command.Command
}

//NewCli initialize
func NewCli() *Cli {

	return &Cli{}
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

func (cli *Cli) parseSystemCommands(input []string) bool {
	if input[0] == "exit" {
		fmt.Println("Bye")
		os.Exit(0)
	}
	if input[0] == "help" {
		cli.printHelp()
		return true
	}
	return false
}

func (cli *Cli) findCommand(input string) *command.Command {

	parsed := strings.Fields(input)
	if len(parsed) == 0 {
		fmt.Println("No input detected")
		return nil
	}
	if cli.parseSystemCommands(parsed) {
		return nil
	}

	currentCommands := cli.Commands
	//Maybe recurse this one day...
	for _, primary := range currentCommands {

		if parsed[0] == primary.Name {

			for _, secondary := range primary.SubCommands {

				if len(parsed) > 1 {
					if parsed[1] == secondary.Name {

						return &secondary
					}
				}
			}

			return &primary
		}
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
	fmt.Println(text)

	command := cli.findCommand(text)
	if command != nil {
		command.Func()
	}
	goto reset
}
