package cmd

import (
	"github.com/ghthor/gospec"
	. "github.com/ghthor/gospec"

	"github.com/ghthor/netmon/cmd/test_cmd"
)

func DescribeCommandCatalog(c gospec.Context) {
	c.Specify("a command catalog", func() {
		cat := NewCatalog()

		c.Specify("registers commands to verbs", func() {
			cmd1 := new(C)
			cmd2 := new(C)

			c.Assume(cat.Register("cmd1", cmd1), IsNil)
			c.Assume(cat.Register("cmd2", cmd2), IsNil)

			c.Expect(cat.MatchVerb("cmd1"), Equals, cmd1)
			c.Expect(cat.MatchVerb("cmd2"), Equals, cmd2)

			c.Specify("unless the verb is already in use", func() {
				c.Expect(cat.Register("cmd1", new(C)), Not(IsNil))
			})
		})

		c.Specify("registers a command package", func() {
			cmd := new(C)

			c.Assume(cat.RegisterAsPkg(cmd), IsNil)
			c.Assume(cat.RegisterAsPkg(test_cmd.C), IsNil)

			c.Expect(cat.MatchVerb("cmd"), Equals, cmd)
			c.Expect(cat.MatchVerb("test_cmd"), Equals, test_cmd.C)
		})
	})
}
