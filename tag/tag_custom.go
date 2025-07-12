package tag

import "github.com/ldez/go-git-cmd-wrapper/v2/types"

// Name [<tagname>]
// The name of the tag to create, delete, or describe.
// The new tag name must pass all checks defined by git-check-ref-format(1).
// Some of these checks may restrict the characters allowed in a tag name.
func Name(tagName string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions(tagName)
	}
}

// Commit <commit>, <object>
// The object that the new tag will refer to, usually a commit. Defaults to HEAD.
func Commit(commit string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions(commit)
	}
}

// Format <format>
// A string that interpolates %(fieldname) from a tag ref being shown and the object it points at.
// The format is the same as that of git-for-each-ref(1).
// When unspecified, defaults to %(refname:strip=2).
func Format(format string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions(format)
	}
}
