package branch

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/v2/types"
)

// Abbrev use <n> digits to display SHA-1s
// --abbrev[=<n>]
func Abbrev(n string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--abbrev=%s", n))
	}
}

// All list both remote-tracking and local branches
// -a, --all
func All(g *types.Cmd) {
	g.AddOptions("--all")
}

// BranchName branch name.
func BranchName(name string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(name)
	}
}

// Color use colored output
// --color[=<when>]
func Color(when string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--color=%s", when))
	}
}

// Column list branches in columns
// --column[=<style>]
func Column(style string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--column=%s", style))
	}
}

// Contains print only branches that contain the commit
// --contains <commit>
func Contains(commit string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--contains")
		g.AddOptions(commit)
	}
}

// NoContains print only branches that don't contain the commit
// --no-contains <commit>
func NoContains(commit string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--no-contains")
		g.AddOptions(commit)
	}
}

// CreateReflog create the branch's reflog
// -l, --create-reflog
func CreateReflog(g *types.Cmd) {
	g.AddOptions("--create-reflog")
}

// Delete delete fully merged branch
// -d, --delete
func Delete(g *types.Cmd) {
	g.AddOptions("--delete")
}

// DeleteForce delete branch (even if not merged)
// -D
func DeleteForce(g *types.Cmd) {
	g.AddOptions("-D")
}

// EditDescription edit the description for the branch
// --edit-description
func EditDescription(g *types.Cmd) {
	g.AddOptions("--edit-description")
}

// Force force creation, move/rename, deletion
// -f, --force
func Force(g *types.Cmd) {
	g.AddOptions("--force")
}

// Format format to use for the output
// --format <format>
func Format(format string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--format")
		g.AddOptions(format)
	}
}

// IgnoreCase sorting and filtering are case insensitive
// -i, --ignore-case
func IgnoreCase(g *types.Cmd) {
	g.AddOptions("--ignore-case")
}

// List list branch names
// --list
func List(g *types.Cmd) {
	g.AddOptions("--list")
}

// PointsAt print only branches of the object
// --points-at <object>
func PointsAt(object string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--points-at")
		g.AddOptions(object)
	}
}

// Merged print only branches that are merged
// --merged <commit>
func Merged(commit string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--merged")
		g.AddOptions(commit)
	}
}

// NoMerged print only branches that are not merged
// --no-merged <commit>
func NoMerged(commit string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--no-merged")
		g.AddOptions(commit)
	}
}

// Move move/rename a branch and its reflog
// -m, --move
func Move(g *types.Cmd) {
	g.AddOptions("--move")
}

// MoveForce move/rename a branch, even if target exists
// -M
func MoveForce(g *types.Cmd) {
	g.AddOptions("-M")
}

// Quiet suppress informational messages
// -q, --quiet
func Quiet(g *types.Cmd) {
	g.AddOptions("--quiet")
}

// Remotes act on remote-tracking branches
// -r, --remotes
func Remotes(g *types.Cmd) {
	g.AddOptions("--remotes")
}

// ShowCurrent print the name of the current branch. In detached HEAD state, nothing is printed.
// --show-current
func ShowCurrent(g *types.Cmd) {
	g.AddOptions("--show-current")
}

// SetUpstream change upstream info
// --set-upstream
func SetUpstream(g *types.Cmd) {
	g.AddOptions("--set-upstream")
}

// UnsetUpstream change the upstream info
// -u, --set-upstream-to <upstream>
func UnsetUpstream(g *types.Cmd) {
	g.AddOptions("--unset-upstream")
}

// SetUpstreamTo change the upstream info
// -u, --set-upstream-to <upstream>
func SetUpstreamTo(upstream string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--set-upstream-to")
		g.AddOptions(upstream)
	}
}

// Sort field name to sort on
// --sort <key>
func Sort(key string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--sort")
		g.AddOptions(key)
	}
}

// Track set up tracking mode (see git-pull(1))
// -t, --track
func Track(g *types.Cmd) {
	g.AddOptions("--track")
}

// Verbose show hash and subject, give twice for upstream branch
// -v, --verbose
func Verbose(g *types.Cmd) {
	g.AddOptions("--verbose")
}
