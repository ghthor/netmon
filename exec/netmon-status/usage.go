package main

import (
	"flag"
	"fmt"
)

func usage() {
	fmt.Print(usagePrefix)
	flag.PrintDefaults()
}

var usagePrefix = `
netmod-status is the cli interface to query a netmond network

Usage:
    netmon-status [options]

Options:
`
