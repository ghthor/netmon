package cmds

import (
	"github.com/ghthor/gospec"
	"testing"
)

func TestUnitSpecs(t *testing.T) {
	r := gospec.NewRunner()

	r.AddSpec(DescribeRegisteredCommands)

	gospec.MainGoTest(r, t)
}
