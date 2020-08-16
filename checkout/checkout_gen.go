package checkout

// CODE GENERATED AUTOMATICALLY
// THIS FILE MUST NOT BE EDITED BY HAND

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/v2/types"
)

// Conflict The same as --merge option above, but changes the way the conflicting hunks are presented, overriding the merge.conflictStyle configuration variable.
// Possible values are 'merge' (default) and 'diff3' (in addition to what is shown by 'merge' style, shows the original contents).
// --conflict=<style>
func Conflict(style string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--conflict=%s", style))
	}
}

// Detach Rather than checking out a branch to work on it, check out a commit for inspection and discardable experiments.
// This is the default behavior of 'git checkout <commit>' when <commit> is not a branch name.
// See the 'DETACHED HEAD' section below for details.
// --detach
func Detach(g *types.Cmd) {
	g.AddOptions("--detach")
}

// Force When switching branches, proceed even if the index or the working tree differs from HEAD.
// This is used to throw away local changes.
// When checking out paths from the index, do not fail upon unmerged entries; instead, unmerged entries are ignored.
// -f, --force
func Force(g *types.Cmd) {
	g.AddOptions("--force")
}

// IgnoreOtherWorktrees git checkout refuses when the wanted ref is already checked out by another worktree.
// This option makes it check the ref out anyway.
// In other words, the ref can be held by more than one worktree.
// --ignore-other-worktrees
func IgnoreOtherWorktrees(g *types.Cmd) {
	g.AddOptions("--ignore-other-worktrees")
}

// IgnoreSkipWorktreeBits In sparse checkout mode, git checkout -- <paths> would update only entries matched by <paths> and sparse patterns in $GIT_DIR/info/sparse-checkout.
// This option ignores the sparse patterns and adds back any files in <paths>.
// --ignore-skip-worktree-bits
func IgnoreSkipWorktreeBits(g *types.Cmd) {
	g.AddOptions("--ignore-skip-worktree-bits")
}

// Merge When switching branches, if you have local modifications to one or more files that are different between the current branch and the branch to which you are switching, the command refuses to switch branches in order to preserve your modifications in context.
// However, with this option, a three-way merge between the current branch, your working tree contents, and the new branch is done, and you will be on the new branch.
// -m, --merge
func Merge(g *types.Cmd) {
	g.AddOptions("--merge")
}

// NewBranch Create a new branch named <new_branch> and start it at <start_point>; see git-branch(1) for details.
// -b [new_branch]
func NewBranch(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("-b")
		if len(value) != 0 {
			g.AddOptions(value)
		}
	}
}

// NewBranchForce Creates the branch <new_branch> and start it at <start_point>; if it already exists, then reset it to <start_point>.
// This is equivalent to running 'git branch' with '-f'; see git-branch(1) for details.
// -B [new_branch]
func NewBranchForce(newBranch string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("-B")
		if len(newBranch) != 0 {
			g.AddOptions(newBranch)
		}
	}
}

// NewBranchReflog Create the new branch’s reflog; see git-branch(1) for details.
// -l
func NewBranchReflog(g *types.Cmd) {
	g.AddOptions("-l")
}

// NoProgress Progress status is reported on the standard error stream by default when it is attached to a terminal, unless --quiet is specified.
// This flag enables progress reporting even if not attached to a terminal, regardless of --quiet.
// --no-progress
func NoProgress(g *types.Cmd) {
	g.AddOptions("--no-progress")
}

// NoRecurseSubmodules Using --recurse-submodules will update the content of all initialized submodules according to the commit recorded in the superproject.
// If local modifications in a submodule would be overwritten the checkout will fail unless -f is used.
// If nothing (or --no-recurse-submodules) is used, the work trees of submodules will not be updated.
// --no-recurse-submodules
func NoRecurseSubmodules(g *types.Cmd) {
	g.AddOptions("--no-recurse-submodules")
}

// NoTrack Do not set up 'upstream' configuration, even if the branch.autoSetupMerge configuration variable is true.
// --no-track
func NoTrack(g *types.Cmd) {
	g.AddOptions("--no-track")
}

// Orphan Create a new orphan branch, named <new_branch>, started from <start_point> and switch to it.
// The first commit made on this new branch will have no parents and it will be the root of a new history totally disconnected from all the other branches and commits.
// --orphan <new_branch>
func Orphan(newBranch string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--orphan")
		g.AddOptions(newBranch)
	}
}

// Ours When checking out paths from the index, check out stage #2 (ours) or #3 (theirs) for unmerged paths.
// --ours, --theirs
func Ours(g *types.Cmd) {
	g.AddOptions("--ours")
}

// Patch Interactively select hunks in the difference between the <tree-ish> (or the index, if unspecified) and the working tree.
// The chosen hunks are then applied in reverse to the working tree (and if a <tree-ish> was specified, the index).
// -p, --patch
func Patch(g *types.Cmd) {
	g.AddOptions("--patch")
}

// Progress Progress status is reported on the standard error stream by default when it is attached to a terminal, unless --quiet is specified.
// This flag enables progress reporting even if not attached to a terminal, regardless of --quiet.
// --progress
func Progress(g *types.Cmd) {
	g.AddOptions("--progress")
}

// Quiet Quiet, suppress feedback messages.
// -q, --quiet
func Quiet(g *types.Cmd) {
	g.AddOptions("--quiet")
}

// RecurseSubmodules Using --recurse-submodules will update the content of all initialized submodules according to the commit recorded in the superproject.
// If local modifications in a submodule would be overwritten the checkout will fail unless -f is used.
// If nothing (or --no-recurse-submodules) is used, the work trees of submodules will not be updated.
// --recurse-submodules
func RecurseSubmodules(g *types.Cmd) {
	g.AddOptions("--recurse-submodules")
}

// Theirs When checking out paths from the index, check out stage #2 (ours) or #3 (theirs) for unmerged paths.
// --ours, --theirs
func Theirs(g *types.Cmd) {
	g.AddOptions("--theirs")
}

// Track When creating a new branch, set up 'upstream' configuration.
// See '--track' in git-branch(1) for details.
// -t, --track
func Track(g *types.Cmd) {
	g.AddOptions("--track")
}
