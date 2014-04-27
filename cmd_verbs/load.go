package cmd_verbs

import (
	"github.com/ghthor/netmon/cmd"
	"github.com/ghthor/netmon/cmd_verbs/daemon"
	"github.com/ghthor/netmon/cmd_verbs/status"
)

var c cmd.Catalog = cmd.NewCatalog()

func init() {
	c.RegisterAsPkg(daemon.Cmd)
	c.RegisterAsPkg(status.Cmd)
}

func MatchVerb(verb string) cmd.Cmd {
	return c.MatchVerb(verb)
}
