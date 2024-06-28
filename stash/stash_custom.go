package stash

import "github.com/ldez/go-git-cmd-wrapper/v2/types"

// Push git stash push [-p|--patch] [-S|--staged] [-k|--[no-]keep-index] [-u|--include-untracked] [-a|--all] [-q|--quiet] [(-m|--message) <message>] [--pathspec-from-file=<file> [--pathspec-file-nul]] [--] [<pathspec>...]
func Push(message string, options ...types.Option) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("push")
		g.ApplyOptions(options...)

		if message != "" {
			g.AddOptions(message)
		}
	}
}

// Save git stash save [-p|--patch] [-S|--staged] [-k|--[no-]keep-index] [-u|--include-untracked] [-a|--all] [-q|--quiet] [<message>]
func Save(message string, options ...types.Option) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("save")
		g.ApplyOptions(options...)

		if message != "" {
			g.AddOptions(message)
		}
	}
}

// List git stash list [<log-options>]
func List(options ...types.Option) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("list")
		g.ApplyOptions(options...)
	}
}

// Show git stash show [-u|--include-untracked|--only-untracked] [<diff-options>] [<stash>]
func Show(stash string, options ...types.Option) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("show")
		g.ApplyOptions(options...)

		if stash != "" {
			g.AddOptions(stash)
		}
	}
}

// Pop git stash pop [--index] [-q|--quiet] [<stash>]
func Pop(stash string, options ...types.Option) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("pop")
		g.ApplyOptions(options...)

		if stash != "" {
			g.AddOptions(stash)
		}
	}
}

// Apply git stash apply [--index] [-q|--quiet] [<stash>]
func Apply(stash string, options ...types.Option) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("apply")
		g.ApplyOptions(options...)

		if stash != "" {
			g.AddOptions(stash)
		}
	}
}

// Branch git stash branch <branchname> [<stash>]
func Branch(branchName, stash string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("branch")
		g.AddOptions(branchName)

		if stash != "" {
			g.AddOptions(stash)
		}
	}
}

// Clear git stash clear
func Clear() func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("clear")
	}
}

// Drop git stash drop [-q|--quiet] [<stash>]
func Drop(stash string, options ...types.Option) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("drop")
		g.ApplyOptions(options...)

		if stash != "" {
			g.AddOptions(stash)
		}
	}
}

// Create git stash create
func Create() func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("create")
	}
}

// Store git stash store
func Store() func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("store")
	}
}

// HyphenHyphen add `--`
func HyphenHyphen(g *types.Cmd) {
	g.AddOptions("--")
}
