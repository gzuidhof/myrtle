package flat

import (
	"testing"

	"github.com/gzuidhof/myrtle/theme"
	"github.com/gzuidhof/myrtle/themetest"
)

func TestThemeSuite(t *testing.T) {
	t.Parallel()

	themetest.RunSuite(t, func() theme.Theme {
		return New()
	})
}
