package notes

// CODE GENERATED AUTOMATICALLY
// THIS FILE MUST NOT BE EDITED BY HAND

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/v2/types"
)

// Force When adding notes to an object that already has notes, overwrite the existing notes (instead of aborting).
// --force, -f
func Force(g *types.Cmd) {
	g.AddOptions("--force")
}

// Ref use notes from <notes-ref>
// --ref=<notes-ref>
func Ref(notesRef string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--ref=%s", notesRef))
	}
}
