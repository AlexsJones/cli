package cli

import "github.com/AlexsJones/cli/command"

//ICli is the common mockable interface for Cli
type ICli interface {
	AddCommand(c command.ICommand)
}

//AddCommand for Cli commands
func AddCommand(i ICli, c command.ICommand) {
	i.AddCommand(c)
}
