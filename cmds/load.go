package cmds

import (
	"github.com/ghthor/netmon/cmd"
	"github.com/ghthor/netmon/cmds/daemon"
	"github.com/ghthor/netmon/cmds/status"
)

var c cmd.Catalog = cmd.NewCatalog()

func init() {
	c.RegisterAsPkg(daemon.Cmd)
	c.RegisterAsPkg(status.Cmd)
}

func MatchVerb(verb string) cmd.Cmd {
	return c.MatchVerb(verb)
}
