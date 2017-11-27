package command

//ICommand interface
type ICommand interface {
	Count() int
}

//Count the number of subcommands
func Count(i ICommand) int {
	return i.Count()
}
