package cmds

import (
	"github.com/ghthor/gospec"
	. "github.com/ghthor/gospec"

	"github.com/ghthor/netmon/cmds/daemon"
	"github.com/ghthor/netmon/cmds/status"
)

func DescribeRegisteredCommands(c gospec.Context) {
	c.Specify("commands that are registered are", func() {
		c.Specify("daemon", func() {
			c.Expect(MatchVerb("daemon"), Equals, daemon.Cmd)
		})

		c.Specify("status", func() {
			c.Expect(MatchVerb("status"), Equals, status.Cmd)
		})
	})
}
