package notes

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/v2/types"
)

// Message Use the given <msg> as the commit message.
// If multiple -m options are given, their values are concatenated as separate paragraphs.
// -m <msg>, --message=<msg>
func Message(msg string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--message=%s", msg))
	}
}
