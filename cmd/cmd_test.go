package cmd

import (
	"github.com/ghthor/gospec"
	. "github.com/ghthor/gospec"
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
	})
}
