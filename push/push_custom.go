package push

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/types"
)

// ForceWithLease alone, without specifying the details, will protect all remote refs that are going to be updated by requiring their current value to be the same as the remote-tracking branch we have for them.
// Usually, "git push" refuses to update a remote ref that is not an ancestor of the local ref used to overwrite it.
// This option overrides this restriction if the current value of the remote ref is the expected value. "git push" fails otherwise.
// --force-with-lease
func ForceWithLease(g *types.Cmd) {
	g.AddOptions("--force-with-lease")
}

// NoForceWithLease Usually, "git push" refuses to update a remote ref that is not an ancestor of the local ref used to overwrite it.
// This option overrides this restriction if the current value of the remote ref is the expected value. "git push" fails otherwise.
// --[no-]force-with-lease
func NoForceWithLease(g *types.Cmd) {
	g.AddOptions("--no-force-with-lease")
}

// ForceWithLeaseRef without specifying the expected value, will protect the named ref (alone), if it is going to be updated, by requiring its current value to be the same as the remote-tracking branch we have for it.
// Usually, "git push" refuses to update a remote ref that is not an ancestor of the local ref used to overwrite it.
// This option overrides this restriction if the current value of the remote ref is the expected value. "git push" fails otherwise.
// --force-with-lease=<refname>
func ForceWithLeaseRef(refname string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--force-with-lease=%s", refname))
	}
}

// ForceWithLeaseExpect will protect the named ref (alone), if it is going to be updated, by requiring its current value to be the same as the specified value <expect> (which is allowed to be different from the remote-tracking branch we have for the refname, or we do not even have to have such a remote-tracking branch when this form is used). If <expect> is the empty string, then the named ref must not already exist.
// Usually, "git push" refuses to update a remote ref that is not an ancestor of the local ref used to overwrite it.
// This option overrides this restriction if the current value of the remote ref is the expected value. "git push" fails otherwise.
// --force-with-lease=<refname>:<expect>
func ForceWithLeaseExpect(refname, expect string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--recurse-submodules=%s:%s", refname, expect))
	}
}

// Remote The "remote" repository that is destination of a push operation.
// This parameter can be either a URL (see the section GIT URLS below) or the name of a remote (see the section REMOTES below).
// <repository>
func Remote(repository string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(repository)
	}
}

// RefSpecSpecify what destination ref to update with what source object.
// The format of a <refspec> parameter is an optional plus +, followed by the source object <src>, followed by a colon :, followed by the destination ref <dst>.
// <refspec>...
func RefSpec(refspecs ...string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		for _, refspec := range refspecs {
			g.AddOptions(refspec)
		}
	}
}
