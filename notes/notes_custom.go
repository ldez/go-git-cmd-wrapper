package notes

import "github.com/ldez/go-git-cmd-wrapper/v2/types"

// Object related to `copy` sub-command (`<from-object> <to-object>`).
func Object(from, to string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(from)
		if to != "" {
			g.AddOptions(to)
		}
	}
}

// NotesRef related to `merge` sub-command (`<notes-ref>`).
func NotesRef(ref string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(ref)
	}
}
