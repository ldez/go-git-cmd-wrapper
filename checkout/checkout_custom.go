package checkout

import (
	"github.com/ldez/go-git-cmd-wrapper/v2/types"
)

// Branch to checkout; if it refers to a branch (i.e., a name that, when prepended with 'refs/heads/', is a valid ref), then that branch is checked out.
// Otherwise, if it refers to a valid commit, your HEAD becomes 'detached' and you are no longer on any branch (see below for details).
// <branch>
// or
// <new_branch> Name for the new branch.
func Branch(name string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions(name)
	}
}

// StartPoint The name of a commit at which to start the new branch; see git-branch(1) for details.
// Defaults to HEAD.
// <start_point>
func StartPoint(name string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions(name)
	}
}

// Path <paths>...
func Path(values ...string) types.Option {
	return func(g *types.Cmd) {
		for _, value := range values {
			g.AddOptions(value)
		}
	}
}

// TreeIsh Tree to checkout from (when paths are given).
// If not specified, the index will be used.
// <tree-ish>
func TreeIsh(value string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions(value)
	}
}
