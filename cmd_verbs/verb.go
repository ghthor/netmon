package cmd_verbs

import (
	"github.com/ghthor/netmon/cmd"
)

var c cmd.Catalog = cmd.NewCatalog()

func MatchVerb(verb string) cmd.Cmd {
	return c.MatchVerb(verb)
}

type VerbUsage struct {
	Verb, Summary string
}

func Usages() (usages []VerbUsage) {
	for verb, cmd := range c {
		usages = append(usages, VerbUsage{verb, cmd.Summary()})
	}
	return
}
