package main

import (
	"flag"
	"os"

	"github.com/ghthor/netmon/cmd_verbs"
)

const (
	EC_OK int = iota
	EC_NO_CMD
	EC_CMD_ERROR
	EC_UNKNOWN_COMMAND
)

func showUsageAndExit(exitCode int) {
	flag.Usage()
	os.Exit(exitCode)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	// Check that a verb exists in the arguments
	args := flag.Args()
	if len(args) == 0 {
		showUsageAndExit(EC_NO_CMD)
	}

	// Retrieve the command bound to the verb
	cmd := cmd_verbs.MatchVerb(args[0])
	if cmd == nil {
		showUsageAndExit(EC_UNKNOWN_COMMAND)
	}

	// Execute the command
	err := cmd.Exec(args[1:])
	if err != nil {
		showUsageAndExit(EC_CMD_ERROR)
	}
}
