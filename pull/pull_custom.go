package pull

import (
	"github.com/ldez/go-git-cmd-wrapper/v2/types"
)

// Repository The "remote" repository that is the source of a fetch or pull operation.
// This parameter can be either a URL (see the section GIT URLS below) or the name of a remote (see the section REMOTES below).
func Repository(remote string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions(remote)
	}
}

// Refspec Specifies which refs to fetch and which local refs to update.
// When no <refspec>s appear on the command line, the refs to fetch are read from remote.<repository>.fetch variables instead.
func Refspec(refs ...string) types.Option {
	return func(g *types.Cmd) {
		for _, ref := range refs {
			g.AddOptions(ref)
		}
	}
}
