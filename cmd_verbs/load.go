package cmd_verbs

import (
	"github.com/ghthor/netmon/cmd_verbs/daemon"
	"github.com/ghthor/netmon/cmd_verbs/status"
)

func init() {
	c.RegisterAsPkg(daemon.Cmd)
	c.RegisterAsPkg(status.Cmd)
}
