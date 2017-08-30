package fetch

import (
	"github.com/ldez/go-git-cmd-wrapper/types"
)

// Remote name.
func Remote(name string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(name)
	}
}

// Group name.
func Group(name string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(name)
	}
}

// RefSpec name.
func RefSpec(ref string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(ref)
	}
}
