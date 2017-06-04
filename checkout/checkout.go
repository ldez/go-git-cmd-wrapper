package checkout

import "github.com/ldez/go-git-cmd-wrapper/git"

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

func Quiet(g *git.Cmd) {
	g.AddOptions("--quiet")
}

func Progress(g *git.Cmd) {
	g.AddOptions("--progress")
}

func NoProgress(g *git.Cmd) {
	g.AddOptions("--no-progress")
}

func Force(g *git.Cmd) {
	g.AddOptions("--force")
}

func Ours(g *git.Cmd) {
	g.AddOptions("--ours")
}

func Theirs(g *git.Cmd) {
	g.AddOptions("--theirs")
}

func Track(g *git.Cmd) {
	g.AddOptions("--track")
}

func NoTrack(g *git.Cmd) {
	g.AddOptions("--no-track")
}

func Detach(g *git.Cmd) {
	g.AddOptions("--detach")
}

func IgnoreSkipWorktreeBits(g *git.Cmd) {
	g.AddOptions("--ignore-skip-worktree-bits")
}

func Merge(g *git.Cmd) {
	g.AddOptions("--merge")
}

func Patch(g *git.Cmd) {
	g.AddOptions("--patch")
}

func IgnoreOtherWorktrees(g *git.Cmd) {
	g.AddOptions("--ignore-other-worktrees")
}

func RecurseSubmodules(g *git.Cmd) {
	g.AddOptions("--recurse-submodules")
}

func NoRecurseSubmodules(g *git.Cmd) {
	g.AddOptions("--no-recurse-submodules")
}

func NewBranch(name string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions("-b")
		g.AddOptions(name)
	}
}

func NewBranchForce(name string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions("-B")
		g.AddOptions(name)
	}
}

func Branch(name string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(name)
	}
}

func Orphan(branch string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions("--orphan")
		g.AddOptions(branch)
	}
}

//-l
//<new_branch>
//<start_point>
//<tree-ish>
//--conflict=<style>
