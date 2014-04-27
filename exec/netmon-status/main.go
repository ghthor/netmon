package main

import (
	"flag"
	"os"

	verb "github.com/ghthor/netmon/cmd_verbs/status"
)

const (
	EC_OK int = iota
	EC_CMD_ERROR
)

func showUsageAndExit(exitCode int) {
	flag.Usage()
	os.Exit(exitCode)
}

func main() {
	flag.Parse()

	args := flag.Args()

	// Execute the command
	err := verb.Cmd.Exec(args)
	if err != nil {
		showUsageAndExit(EC_CMD_ERROR)
	}
}
