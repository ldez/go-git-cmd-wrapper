package revparse

import (
	"github.com/ldez/go-git-cmd-wrapper/v2/types"
)

// Args Flags and parameters to be parsed.
func Args(args ...string) types.Option {
	return func(g *types.Cmd) {
		for _, arg := range args {
			g.AddOptions(arg)
		}
	}
}
