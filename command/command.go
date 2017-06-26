package command

//Command structure
type Command struct {
	Name        string
	Help        string
	Func        func()
	SubCommands []Command
}

//NewCommand ...
func NewCommand() *Command {
	c := new(Command)
	c.SubCommands = []Command{}
	return c
}
