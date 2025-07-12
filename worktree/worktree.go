package worktree

import (
	"github.com/ldez/go-git-cmd-wrapper/v2/types"
)

// Add git worktree add [-f] [--detach] [--checkout] [--lock] [-b <new-branch>] <path> [<branch>]
func Add(path, branch string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("add")
		g.AddOptions(path)

		if branch != "" {
			g.AddOptions(branch)
		}
	}
}

// List git worktree list [--porcelain]
func List(g *types.Cmd) {
	g.AddOptions("list")
}

// Lock git worktree lock [--reason <string>] <worktree>
func Lock(worktree string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("lock")
		g.AddOptions(worktree)
	}
}

// Prune git worktree prune [-n] [-v] [--expire <expire>]
func Prune(g *types.Cmd) {
	g.AddOptions("prune")
}

// UnLock git worktree unlock <worktree>
func UnLock(worktree string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("unlock")
		g.AddOptions(worktree)
	}
}

// Force [-f | --force]
func Force(g *types.Cmd) {
	g.AddOptions("--force")
}

// Branch [-b]
func Branch(branch string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("-b")

		if branch != "" {
			g.AddOptions(branch)
		}
	}
}

// BranchOverride [-B]
func BranchOverride(branch string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("-B")

		if branch != "" {
			g.AddOptions(branch)
		}
	}
}

// Detach [--detach]
func Detach(g *types.Cmd) {
	g.AddOptions("--detach")
}

// Checkout [--checkout]
func Checkout(g *types.Cmd) {
	g.AddOptions("--checkout")
}

// NoCheckout [--no-checkout]
func NoCheckout(g *types.Cmd) {
	g.AddOptions("--no-checkout")
}

// LockOption [--lock]
func LockOption(g *types.Cmd) {
	g.AddOptions("--lock")
}

// DryRun [-n | --dry-run]
func DryRun(g *types.Cmd) {
	g.AddOptions("--dry-run")
}

// Porcelain [--porcelain]
func Porcelain(g *types.Cmd) {
	g.AddOptions("--porcelain")
}

// Verbose [-v | --verbose]
func Verbose(g *types.Cmd) {
	g.AddOptions("--verbose")
}

// Expire [--expire <time>]
func Expire(time string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("--expire")

		if time != "" {
			g.AddOptions(time)
		}
	}
}

// Reason [--reason <string>]
func Reason(value string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("--reason")

		if value != "" {
			g.AddOptions(value)
		}
	}
}
