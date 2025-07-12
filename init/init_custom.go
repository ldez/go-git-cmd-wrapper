package init

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/v2/types"
)

// SharedWithPerms Shared Specify that the Git repository is to be shared amongst several users.
// This allows users belonging to the same group to push into that repository.
// When specified, the config variable 'core.sharedRepository' is set so that files and directories under $GIT_DIR are created with the requested permissions.
// When not specified, Git will use permissions reported by umask(2).
// --shared[=(false|true|umask|group|all|world|everybody|0xxx)]
func SharedWithPerms(permissions string) types.Option {
	return func(g *types.Cmd) {
		if permissions == "" {
			g.AddOptions("--shared")
		} else {
			g.AddOptions(fmt.Sprintf("--shared=%s", permissions))
		}
	}
}

// Directory path.
func Directory(directory string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions(directory)
	}
}
