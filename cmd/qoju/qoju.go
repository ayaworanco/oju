package main

import (
	"flag"
	"fmt"
	"oju/internal/commander"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprint(commander.USAGE))
	}

	if len(os.Args) == 1 {
		commander.UsageAndExit("Should have at least 1 argument")
	}

	var cmd *commander.Command

	switch os.Args[1] {
	case "version":
		cmd = commander.VersionCommand()
	case "watch":
		cmd = commander.WatchCommand()
	default:
		commander.UsageAndExit(fmt.Sprintf("qoju: '%s' is not a valid command.\n", os.Args[1]))
	}

	cmd.Init(os.Args[2:])
	cmd.Run()
}
