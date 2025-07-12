package lsfiles

import "github.com/ldez/go-git-cmd-wrapper/v2/types"

// HyphenHyphen add `--`.
// Do not interpret any more arguments as options.
func HyphenHyphen(g *types.Cmd) {
	g.AddOptions("--")
}

// Files [<file>...].
// Files to show.
// If no files are given all files which match the other specified criteria are shown.
func Files(files ...string) types.Option {
	return func(g *types.Cmd) {
		for _, file := range files {
			g.AddOptions(file)
		}
	}
}
