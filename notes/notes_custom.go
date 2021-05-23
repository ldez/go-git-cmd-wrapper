package notes

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/v2/types"
)

// Add git notes add [-f] [--allow-empty] [-F <file> | -m <msg> | (-c | -C) <object>] [<object>]
func Add(g *types.Cmd) {
	g.AddOptions("add")
}

// List git notes list [<object]
func List(g *types.Cmd) {
	g.AddOptions("list")
}

// Message Use the given <msg> as the commit message.
// If multiple -m options are given, their values are concatenated as separate paragraphs.
// -m <msg>, --message=<msg>
func Message(msg string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--message=%s", msg))
	}
}

// Show git notes show [<object]
func Show(notesRef string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("show")
		if notesRef != "" {
			g.AddOptions(notesRef)
		}
	}
}
