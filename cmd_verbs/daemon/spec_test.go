package daemon

import (
	"github.com/ghthor/gospec"
	"testing"
)

func TestUnitSpecs(t *testing.T) {
	r := gospec.NewRunner()

	/* r.AddSpec() */

	gospec.MainGoTest(r, t)
}
