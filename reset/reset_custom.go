package reset

import "github.com/ldez/go-git-cmd-wrapper/types"

// HyphenHyphen add `--`
func HyphenHyphen(g *types.Cmd) {
	g.AddOptions("--")
}

// Path <paths>...
func Path(values ...string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		for _, value := range values {
			g.AddOptions(value)
		}
	}
}

// Commit [<commit>]
func Commit(hash string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(hash)
	}
}

// TreeIsh [<tree-ish>]
func TreeIsh(hash string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(hash)
	}
}
