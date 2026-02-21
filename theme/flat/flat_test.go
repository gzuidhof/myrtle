package flat

import (
	"testing"

	"github.com/gzuidhof/myrtle/theme"
	"github.com/gzuidhof/myrtle/themetest"
)

func TestThemeSuite(t *testing.T) {
	themetest.RunSuite(t, func() theme.Theme {
		return New()
	})
}
