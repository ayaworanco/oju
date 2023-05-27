package commander

import (
	"flag"
)

type Command struct {
	flags   *flag.FlagSet
	Execute func(cmd *Command, args []string)
}

func (command *Command) Init(args []string) error {
	return command.flags.Parse(args)
}

func (command *Command) Called() bool {
	return command.flags.Parsed()
}

func (command *Command) Run() {
	command.Execute(command, command.flags.Args())
}
