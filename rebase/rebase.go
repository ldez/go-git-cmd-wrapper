package rebase

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/types"
)

// Abort Abort the rebase operation and reset HEAD to the original branch. If <branch> was provided when the rebase operation was started, then HEAD will be reset to <branch>. Otherwise HEAD will be reset to where it was when the rebase operation was started.
// --abort
func Abort(g *types.Cmd) {
	g.AddOptions("--abort")
}

// Autosquash When the commit log message begins with 'squash! ...' (or 'fixup! ...'), and there is a commit whose title begins with the same ..., automatically modify the todo list of rebase -i so that the commit marked for squashing comes right after the commit to be modified, and change the action of the moved commit from pick to squash (or fixup). Ignores subsequent 'fixup! ' or 'squash! ' after the first, in case you referred to an earlier fixup/squash with git commit --fixup/--squash. This option is only valid when the --interactive option is used.
// --autosquash, --no-autosquash
func Autosquash(g *types.Cmd) {
	g.AddOptions("--autosquash")
}

// Autostash Automatically create a temporary stash before the operation begins, and apply it after the operation ends. This means that you can run rebase on a dirty worktree. However, use with care: the final stash application after a successful rebase might result in non-trivial conflicts.
// --autostash, --no-autostash
func Autostash(g *types.Cmd) {
	g.AddOptions("--autostash")
}

// CommitterDateIsAuthorDate These flags are passed to git am to easily change the dates of the rebased commits (see git-am(1)). Incompatible with the --interactive option.
// --ignore-date
func CommitterDateIsAuthorDate(g *types.Cmd) {
	g.AddOptions("--committer-date-is-author-date")
}

// Continue Restart the rebasing process after having resolved a merge conflict.
// --continue
func Continue(g *types.Cmd) {
	g.AddOptions("--continue")
}

// EditTodo Edit the todo list during an interactive rebase.
// --edit-todo
func EditTodo(g *types.Cmd) {
	g.AddOptions("--edit-todo")
}

// Exec Append 'exec <cmd>' after each line creating a commit in the final history. <cmd> will be interpreted as one or more shell commands.
// -x <cmd>, --exec <cmd>
func Exec(cmd string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--exec")
		g.AddOptions(cmd)
	}
}

// ForceRebase Force a rebase even if the current branch is up-to-date and the command without --force would return without doing anything.
// -f, --force-rebase
func ForceRebase(g *types.Cmd) {
	g.AddOptions("--force-rebase")
}

// ForkPoint Use reflog to find a better common ancestor between <upstream> and <branch> when calculating which commits have been introduced by <branch>.
// --fork-point, --no-fork-point
func ForkPoint(g *types.Cmd) {
	g.AddOptions("--fork-point")
}

// GpgSign GPG-sign commits. The keyid argument is optional and defaults to the committer identity; if specified, it must be stuck to the option without a space.
// -S[<keyid>], --gpg-sign[=<keyid>]
func GpgSign(keyid string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(keyid) == 0 {
			g.AddOptions("--gpg-sign")
		} else {
			g.AddOptions(fmt.Sprintf("--gpg-sign=%s", keyid))
		}
	}
}

// IgnoreDate These flags are passed to git am to easily change the dates of the rebased commits (see git-am(1)). Incompatible with the --interactive option.
// --ignore-date
func IgnoreDate(g *types.Cmd) {
	g.AddOptions("--ignore-date")
}

// IgnoreWhitespace These flag are passed to the git apply program (see git-apply(1)) that applies the patch. Incompatible with the --interactive option.
// --ignore-whitespace
func IgnoreWhitespace(g *types.Cmd) {
	g.AddOptions("--ignore-whitespace")
}

// Interactive Make a list of the commits which are about to be rebased. Let the user edit that list before rebasing. This mode can also be used to split commits (see SPLITTING COMMITS below).
// -i, --interactive
func Interactive(g *types.Cmd) {
	g.AddOptions("--interactive")
}

// KeepEmpty Keep the commits that do not change anything from its parents in the result.
// --keep-empty
func KeepEmpty(g *types.Cmd) {
	g.AddOptions("--keep-empty")
}

// Merge Use merging strategies to rebase. When the recursive (default) merge strategy is used, this allows rebase to be aware of renames on the upstream side.
// -m, --merge
func Merge(g *types.Cmd) {
	g.AddOptions("--merge")
}

// NoAutosquash When the commit log message begins with 'squash! ...' (or 'fixup! ...'), and there is a commit whose title begins with the same ..., automatically modify the todo list of rebase -i so that the commit marked for squashing comes right after the commit to be modified, and change the action of the moved commit from pick to squash (or fixup). Ignores subsequent 'fixup! ' or 'squash! ' after the first, in case you referred to an earlier fixup/squash with git commit --fixup/--squash. This option is only valid when the --interactive option is used.
// --autosquash, --no-autosquash
func NoAutosquash(g *types.Cmd) {
	g.AddOptions("--no-autosquash")
}

// NoAutostash Automatically create a temporary stash before the operation begins, and apply it after the operation ends. This means that you can run rebase on a dirty worktree. However, use with care: the final stash application after a successful rebase might result in non-trivial conflicts.
// --autostash, --no-autostash
func NoAutostash(g *types.Cmd) {
	g.AddOptions("--no-autostash")
}

// NoFf With --interactive, cherry-pick all rebased commits instead of fast-forwarding over the unchanged ones. This ensures that the entire history of the rebased branch is composed of new commits. Without --interactive, this is a synonym for --force-rebase.
// --no-ff
func NoFf(g *types.Cmd) {
	g.AddOptions("--no-ff")
}

// NoForkPoint Use reflog to find a better common ancestor between <upstream> and <branch> when calculating which commits have been introduced by <branch>.
// --fork-point, --no-fork-point
func NoForkPoint(g *types.Cmd) {
	g.AddOptions("--no-fork-point")
}

// NoStat Do not show a diffstat as part of the rebase process.
// -n, --no-stat
func NoStat(g *types.Cmd) {
	g.AddOptions("--no-stat")
}

// NoVerify This option bypasses the pre-rebase hook. See also githooks(5).
// --no-verify
func NoVerify(g *types.Cmd) {
	g.AddOptions("--no-verify")
}

// Onto Starting point at which to create the new commits. If the --onto option is not specified, the starting point is <upstream>. May be any valid commit, and not just an existing branch name.
// --onto <newbase>
func Onto(newbase string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--onto")
		g.AddOptions(newbase)
	}
}

// PreserveMerges Recreate merge commits instead of flattening the history by replaying commits a merge commit introduces. Merge conflict resolutions or manual amendments to merge commits are not preserved.
// -p, --preserve-merges
func PreserveMerges(g *types.Cmd) {
	g.AddOptions("--preserve-merges")
}

// Quiet Be quiet. Implies --no-stat.
// -q, --quiet
func Quiet(g *types.Cmd) {
	g.AddOptions("--quiet")
}

// Quit Abort the rebase operation but HEAD is not reset back to the original branch. The index and working tree are also left unchanged as a result.
// --quit
func Quit(g *types.Cmd) {
	g.AddOptions("--quit")
}

// Root Rebase all commits reachable from <branch>, instead of limiting them with an <upstream>. This allows you to rebase the root commit(s) on a branch. When used with --onto, it will skip changes already contained in <newbase> (instead of <upstream>) whereas without --onto it will operate on every change. When used together with both --onto and --preserve-merges, all root commits will be rewritten to have <newbase> as parent instead.
// --root
func Root(g *types.Cmd) {
	g.AddOptions("--root")
}

// Signoff This flag is passed to git am to sign off all the rebased commits (see git-am(1)). Incompatible with the --interactive option.
// --signoff
func Signoff(g *types.Cmd) {
	g.AddOptions("--signoff")
}

// Skip Restart the rebasing process by skipping the current patch.
// --skip
func Skip(g *types.Cmd) {
	g.AddOptions("--skip")
}

// Stat Show a diffstat of what changed upstream since the last rebase. The diffstat is also controlled by the configuration option rebase.stat.
// --stat
func Stat(g *types.Cmd) {
	g.AddOptions("--stat")
}

// Strategy Use the given merge strategy. If there is no -s option git merge-recursive is used instead. This implies --merge.
// -s <strategy>, --strategy=<strategy>
func Strategy(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--strategy=%s", value))
	}
}

// StrategyOption Pass the <strategy-option> through to the merge strategy. This implies --merge and, if no strategy has been specified, -s recursive. Note the reversal of ours and theirs as noted above for the -m option.
// -X <strategy-option>, --strategy-option=<strategy-option>
func StrategyOption(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--strategy-option=%s", value))
	}
}

// Verbose Be verbose. Implies --stat.
// -v, --verbose
func Verbose(g *types.Cmd) {
	g.AddOptions("--verbose")
}

// Verify Allows the pre-rebase hook to run, which is the default. This option can be used to override --no-verify. See also githooks(5).
// --verify
func Verify(g *types.Cmd) {
	g.AddOptions("--verify")
}

// Whitespace These flag are passed to the git apply program (see git-apply(1)) that applies the patch. Incompatible with the --interactive option.
// --whitespace=<option>
func Whitespace(option string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--whitespace=%s", option))
	}
}

// Upstream Upstream branch to compare against. May be any valid commit, not just an existing branch name. Defaults to the configured upstream for the current branch.
func Upstream(name string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(name)
	}
}

// Branch Working branch; defaults to HEAD.
func Branch(name string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(name)
	}
}

//-C<n>
