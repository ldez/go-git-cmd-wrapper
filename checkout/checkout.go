// Checkout https://git-scm.com/docs/git-checkout
// git checkout [-q] [-f] [-m] [<branch>]
// git checkout [-q] [-f] [-m] --detach [<branch>]
// git checkout [-q] [-f] [-m] [--detach] <commit>
// git checkout [-q] [-f] [-m] [[-b|-B|--orphan] <new_branch>] [<start_point>]
// git checkout [-f|--ours|--theirs|-m|--conflict=<style>] [<tree-ish>] [--] <paths>…​
// git checkout [-p|--patch] [<tree-ish>] [--] [<paths>…​]
// --quiet
// --[no-]progress
// --force
// --ours
// --theirs
// -b <new_branch>
// -B <new_branch>
// --track
// --no-track
// -l
// --detach
// --orphan <new_branch>
// --ignore-skip-worktree-bits
// --merge
// --conflict=<style>
// --patch
// --ignore-other-worktrees
// --[no-]recurse-submodules
// <branch>
// <new_branch>
// <start_point>
// <tree-ish>
package checkout

import (
	"fmt"
	"github.com/ldez/go-git-cmd-wrapper/types"
)

func Quiet(g *types.Cmd) {
	g.AddOptions("--quiet")
}

func Progress(g *types.Cmd) {
	g.AddOptions("--progress")
}

func NoProgress(g *types.Cmd) {
	g.AddOptions("--no-progress")
}

func Force(g *types.Cmd) {
	g.AddOptions("--force")
}

func Ours(g *types.Cmd) {
	g.AddOptions("--ours")
}

func Theirs(g *types.Cmd) {
	g.AddOptions("--theirs")
}

func Track(g *types.Cmd) {
	g.AddOptions("--track")
}

func NoTrack(g *types.Cmd) {
	g.AddOptions("--no-track")
}

func Detach(g *types.Cmd) {
	g.AddOptions("--detach")
}

func IgnoreSkipWorktreeBits(g *types.Cmd) {
	g.AddOptions("--ignore-skip-worktree-bits")
}

func Merge(g *types.Cmd) {
	g.AddOptions("--merge")
}

func Patch(g *types.Cmd) {
	g.AddOptions("--patch")
}

func IgnoreOtherWorktrees(g *types.Cmd) {
	g.AddOptions("--ignore-other-worktrees")
}

func RecurseSubmodules(g *types.Cmd) {
	g.AddOptions("--recurse-submodules")
}

func NoRecurseSubmodules(g *types.Cmd) {
	g.AddOptions("--no-recurse-submodules")
}

func NewBranch(g *types.Cmd) {
	g.AddOptions("-b")
}

func NewBranchForce(g *types.Cmd) {
	g.AddOptions("-B")
}

func Branch(name string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(name)
	}
}

func Orphan(branch string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--orphan")
		g.AddOptions(branch)
	}
}

func StartPoint(name string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(name)
	}
}

// Conflict
// --conflict=<style>
func Conflict(style string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--conflict=%s", style))
	}
}

//-l
//<new_branch>
//<tree-ish>
