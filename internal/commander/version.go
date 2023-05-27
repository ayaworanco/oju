package commander

import (
	"flag"
	"fmt"
	"os"
)

var (
	build   = "alpha"
	version = "0.1.2"
)

var version_handler = func(cmd *Command, args []string) {
	fmt.Printf("version %s-%s\n", version, build)
	os.Exit(0)
}

func VersionCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("version", flag.ExitOnError),
		Execute: version_handler,
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, version_usage)
	}
	return cmd
}
