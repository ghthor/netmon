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
netmond is the network status monitoring daemon of netmon

Usage:
    netmond [options]

Options:
`
