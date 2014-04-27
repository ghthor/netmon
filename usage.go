package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/ghthor/netmon/cmd_verbs"
)

func usage() {
	fmt.Print(usagePrefix)
	flag.PrintDefaults()
	verbsTmpl.Execute(os.Stdout, cmd_verbs.Usages())
}

var usagePrefix = `
netmon is a distributed network status/life monitor

Usage:
    netmon [options] <subcommand verb> [verb options]

Options:
`
var verbsTmpl = template.Must(template.New("verbs").Parse(
	`
Verbs:{{range .}}
    {{.Verb | printf "%-10s"}} {{.Summary}}{{end}}
`))
