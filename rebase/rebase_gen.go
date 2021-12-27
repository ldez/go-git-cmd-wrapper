package rebase

// CODE GENERATED AUTOMATICALLY
// THIS FILE MUST NOT BE EDITED BY HAND

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/v2/types"
)

// Abort Abort the rebase operation and reset HEAD to the original branch.
// If <branch> was provided when the rebase operation was started, then HEAD will be reset to <branch>.
// Otherwise HEAD will be reset to where it was when the rebase operation was started.
// --abort
func Abort(g *types.Cmd) {
	g.AddOptions("--abort")
}

// AllowEmptyMessage No-op.
// Rebasing commits with an empty message used to fail and this option would override that behavior, allowing commits with empty messages to be rebased.
// Now commits with an empty message do not cause rebasing to halt.
// --allow-empty-message
func AllowEmptyMessage(g *types.Cmd) {
	g.AddOptions("--allow-empty-message")
}

// Apply Use applying strategies to rebase (calling git-am internally).
// This option may become a no-op in the future once the merge backend handles everything the apply one does.
// --apply
func Apply(g *types.Cmd) {
	g.AddOptions("--apply")
}

// Autosquash When the commit log message begins with "squash! ..." or "fixup! ..." or "amend! ...", and there is already a commit in the todo list that matches the same ..., automatically modify the todo list of rebase -i, so that the commit marked for squashing comes right after the commit to be modified, and change the action of the moved commit from pick to squash or fixup or fixup -C respectively.
// A commit matches the ...  if the commit subject matches, or if the ...  refers to the commit’s hash.
// As a fall-back, partial matches of the commit subject work, too.
// The recommended way to create fixup/amend/squash commits is by using the --fixup, --fixup=amend: or --fixup=reword: and --squash options respectively of git-commit(1).
//
// If the --autosquash option is enabled by default using the configuration variable rebase.autoSquash, this option can be used to override and disable this setting.
// --autosquash, --no-autosquash
func Autosquash(g *types.Cmd) {
	g.AddOptions("--autosquash")
}

// Autostash Automatically create a temporary stash entry before the operation begins, and apply it after the operation ends.
// This means that you can run rebase on a dirty worktree.
// However, use with care: the final stash application after a successful rebase might result in non-trivial conflicts.
// --autostash, --no-autostash
func Autostash(g *types.Cmd) {
	g.AddOptions("--autostash")
}

// CommitterDateIsAuthorDate Instead of using the current time as the committer date, use the author date of the commit being rebased as the committer date.
// This option implies --force-rebase.
// --committer-date-is-author-date
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

// Empty How to handle commits that are not empty to start and are not clean cherry-picks of any upstream commit, but which become empty after rebasing (because they contain a subset of already upstream changes).
// With drop (the default), commits that become empty are dropped.
// With keep, such commits are kept.
// With ask (implied by --interactive), the rebase will halt when an empty commit is applied allowing you to choose whether to drop it, edit files more, or just commit the empty changes.
// Other options, like --exec, will use the default of drop unless -i/--interactive is explicitly specified.
//
// Note that commits which start empty are kept (unless --no-keep-empty is specified), and commits which are clean cherry-picks (as determined by git log --cherry-mark ...) are detected and dropped as a preliminary step (unless --reapply-cherry-picks is passed).
// --empty={drop,keep,ask}
func Empty(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--empty=%s", value))
	}
}

// EnsureContext Ensure at least <n> lines of surrounding context match before and after each change.
// When fewer lines of surrounding context exist they all must match.
// By default no context is ever ignored.
// Implies --apply.
// -C<n>
func EnsureContext(n string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("-C")
		g.AddOptions(n)
	}
}

// Exec Append "exec <cmd>" after each line creating a commit in the final history.
// <cmd> will be interpreted as one or more shell commands.
// Any command that fails will interrupt the rebase, with exit code 1.
// -x <cmd>, --exec <cmd>
func Exec(cmd string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--exec")
		g.AddOptions(cmd)
	}
}

// ForceRebase Individually replay all rebased commits instead of fast-forwarding over the unchanged ones.
// This ensures that the entire history of the rebased branch is composed of new commits.
//
// You may find this helpful after reverting a topic branch merge, as this option recreates the topic branch with fresh commits so it can be remerged successfully without needing to "revert the reversion" (see the revert-a-faulty-merge How-To[1] for details).
// --no-ff, --force-rebase, -f
func ForceRebase(g *types.Cmd) {
	g.AddOptions("--force-rebase")
}

// ForkPoint Use reflog to find a better common ancestor between <upstream> and <branch> when calculating which commits have been introduced by <branch>.
//
// When --fork-point is active, fork_point will be used instead of <upstream> to calculate the set of commits to rebase, where fork_point is the result of git merge-base --fork-point <upstream> <branch> command (see git-merge-base(1)).
// If fork_point ends up being empty, the <upstream> will be used as a fallback.
//
// If <upstream> is given on the command line, then the default is --no-fork-point, otherwise the default is --fork-point.
// See also rebase.forkpoint in git-config(1).
//
// If your branch was based on <upstream> but <upstream> was rewound and your branch contains commits which were dropped, this option can be used with --keep-base in order to drop those commits from your branch.
// --fork-point, --no-fork-point
func ForkPoint(g *types.Cmd) {
	g.AddOptions("--fork-point")
}

// GpgSign GPG-sign commits.
// The keyid argument is optional and defaults to the committer identity; if specified, it must be stuck to the option without a space.
// --no-gpg-sign is useful to countermand both commit.gpgSign configuration variable, and earlier --gpg-sign.
// -S[<keyid>], --gpg-sign[=<keyid>], --no-gpg-sign
func GpgSign(keyID string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(keyID) == 0 {
			g.AddOptions("--gpg-sign")
		} else {
			g.AddOptions(fmt.Sprintf("--gpg-sign=%s", keyID))
		}
	}
}

// IgnoreDate Instead of using the author date of the original commit, use the current time as the author date of the rebased commit.
// This option implies --force-rebase.
// --ignore-date, --reset-author-date
func IgnoreDate(g *types.Cmd) {
	g.AddOptions("--ignore-date")
}

// IgnoreWhitespace Ignore whitespace differences when trying to reconcile differences.
// Currently, each backend implements an approximation of this behavior:
//
// apply backend: When applying a patch, ignore changes in whitespace in context lines.
// Unfortunately, this means that if the "old" lines being replaced by the patch differ only in whitespace from the existing file, you will get a merge conflict instead of a successful patch application.
//
// merge backend: Treat lines with only whitespace changes as unchanged when merging.
// Unfortunately, this means that any patch hunks that were intended to modify whitespace and nothing else will be dropped, even if the other side had no changes that conflicted.
// --ignore-whitespace
func IgnoreWhitespace(g *types.Cmd) {
	g.AddOptions("--ignore-whitespace")
}

// Interactive Make a list of the commits which are about to be rebased.
// Let the user edit that list before rebasing.
// This mode can also be used to split commits (see SPLITTING COMMITS below).
//
// The commit list format can be changed by setting the configuration option rebase.instructionFormat.
// A customized instruction format will automatically have the long commit hash prepended to the format.
// -i, --interactive
func Interactive(g *types.Cmd) {
	g.AddOptions("--interactive")
}

// KeepBase Set the starting point at which to create the new commits to the merge base of <upstream> <branch>.
// Running git rebase --keep-base <upstream> <branch> is equivalent to running git rebase --onto <upstream>... <upstream>.
//
// This option is useful in the case where one is developing a feature on top of an upstream branch.
// While the feature is being worked on, the upstream branch may advance and it may not be the best idea to keep rebasing on top of the upstream but to keep the base commit as-is.
//
// Although both this option and --fork-point find the merge base between <upstream> and <branch>, this option uses the merge base as the starting point on which new commits will be created, whereas --fork-point uses the merge base to determine the set of commits which will be rebased.
// --keep-base
func KeepBase(g *types.Cmd) {
	g.AddOptions("--keep-base")
}

// KeepEmpty Do not keep commits that start empty before the rebase (i.e. that do not change anything from its parent) in the result.
// The default is to keep commits which start empty, since creating such commits requires passing the --allow-empty override flag to git commit, signifying that a user is very intentionally creating such a commit and thus wants to keep it.
//
// Usage of this flag will probably be rare, since you can get rid of commits that start empty by just firing up an interactive rebase and removing the lines corresponding to the commits you don’t want.
// This flag exists as a convenient shortcut, such as for cases where
// external tools generate many empty commits and you want them all removed.
//
// For commits which do not start empty but become empty after rebasing, see the --empty flag.
// --no-keep-empty, --keep-empty
func KeepEmpty(g *types.Cmd) {
	g.AddOptions("--keep-empty")
}

// Merge Using merging strategies to rebase (default).
//
// Note that a rebase merge works by replaying each commit from the working branch on top of the <upstream> branch.
// Because of this, when a merge conflict happens, the side reported as ours is the so-far rebased series, starting with <upstream>, and theirs is the working branch.
// In other words, the sides are swapped.
// -m, --merge
func Merge(g *types.Cmd) {
	g.AddOptions("--merge")
}

// NoAutosquash When the commit log message begins with "squash! ..." or "fixup! ..." or "amend! ...", and there is already a commit in the todo list that matches the same ..., automatically modify the todo list of rebase -i, so that the commit marked for squashing comes right after the commit to be modified, and change the action of the moved commit from pick to squash or fixup or fixup -C respectively.
// A commit matches the ...  if the commit subject matches, or if the ...  refers to the commit’s hash.
// As a fall-back, partial matches of the commit subject work, too.
// The recommended way to create fixup/amend/squash commits is by using the --fixup, --fixup=amend: or --fixup=reword: and --squash options respectively of git-commit(1).
//
// If the --autosquash option is enabled by default using the configuration variable rebase.autoSquash, this option can be used to override and disable this setting.
// --autosquash, --no-autosquash
func NoAutosquash(g *types.Cmd) {
	g.AddOptions("--no-autosquash")
}

// NoAutostash Automatically create a temporary stash entry before the operation begins, and apply it after the operation ends.
// This means that you can run rebase on a dirty worktree.
// However, use with care: the final stash application after a successful rebase might result in non-trivial conflicts.
// --autostash, --no-autostash
func NoAutostash(g *types.Cmd) {
	g.AddOptions("--no-autostash")
}

// NoFf Individually replay all rebased commits instead of fast-forwarding over the unchanged ones.
// This ensures that the entire history of the rebased branch is composed of new commits.
//
// You may find this helpful after reverting a topic branch merge, as this option recreates the topic branch with fresh commits so it can be remerged successfully without needing to "revert the reversion" (see the revert-a-faulty-merge How-To[1] for details).
// --no-ff, --force-rebase, -f
func NoFf(g *types.Cmd) {
	g.AddOptions("--no-ff")
}

// NoForkPoint Use reflog to find a better common ancestor between <upstream> and <branch> when calculating which commits have been introduced by <branch>.
//
// When --fork-point is active, fork_point will be used instead of <upstream> to calculate the set of commits to rebase, where fork_point is the result of git merge-base --fork-point <upstream> <branch> command (see git-merge-base(1)).
// If fork_point ends up being empty, the <upstream> will be used as a fallback.
//
// If <upstream> is given on the command line, then the default is --no-fork-point, otherwise the default is --fork-point.
// See also rebase.forkpoint in git-config(1).
//
// If your branch was based on <upstream> but <upstream> was rewound and your branch contains commits which were dropped, this option can be used with --keep-base in order to drop those commits from your branch.
// --fork-point, --no-fork-point
func NoForkPoint(g *types.Cmd) {
	g.AddOptions("--no-fork-point")
}

// NoGpgSign GPG-sign commits.
// The keyid argument is optional and defaults to the committer identity; if specified, it must be stuck to the option without a space.
// --no-gpg-sign is useful to countermand both commit.gpgSign configuration variable, and earlier --gpg-sign.
// -S[<keyid>], --gpg-sign[=<keyid>], --no-gpg-sign
func NoGpgSign(g *types.Cmd) {
	g.AddOptions("--no-gpg-sign")
}

// NoKeepEmpty Do not keep commits that start empty before the rebase (i.e. that do not change anything from its parent) in the result.
// The default is to keep commits which start empty, since creating such commits requires passing the --allow-empty override flag to git commit, signifying that a user is very intentionally creating such a commit and thus wants to keep it.
//
// Usage of this flag will probably be rare, since you can get rid of commits that start empty by just firing up an interactive rebase and removing the lines corresponding to the commits you don’t want.
// This flag exists as a convenient shortcut, such as for cases where
// external tools generate many empty commits and you want them all removed.
//
// For commits which do not start empty but become empty after rebasing, see the --empty flag.
// --no-keep-empty, --keep-empty
func NoKeepEmpty(g *types.Cmd) {
	g.AddOptions("--no-keep-empty")
}

// NoReapplyCherryPicks Reapply all clean cherry-picks of any upstream commit instead of preemptively dropping them.
// (If these commits then become empty after rebasing, because they contain a subset of already upstream changes, the behavior towards them is controlled by the --empty flag.)
//
// By default (or if --no-reapply-cherry-picks is given), these commits will be automatically dropped.
// Because this necessitates reading all upstream commits, this can be expensive in repos with a large number of upstream commits that need to be read.
// When using the merge backend, warnings will be issued for each dropped commit (unless --quiet is given).
// Advice will also be issued unless advice.skippedCherryPicks is set to false (see git-config(1)).
// --reapply-cherry-picks allows rebase to forgo reading all upstream commits, potentially improving performance.
// --reapply-cherry-picks, --no-reapply-cherry-picks
func NoReapplyCherryPicks(g *types.Cmd) {
	g.AddOptions("--no-reapply-cherry-picks")
}

// NoRerereAutoupdate Allow the rerere mechanism to update the index with the result of auto-conflict resolution if possible.
// --rerere-autoupdate, --no-rerere-autoupdate
func NoRerereAutoupdate(g *types.Cmd) {
	g.AddOptions("--no-rerere-autoupdate")
}

// NoRescheduleFailedExec Automatically reschedule exec commands that failed.
// This only makes sense in interactive mode (or when an --exec option was provided).
//
// Even though this option applies once a rebase is started, it’s set for the whole rebase at the start based on either the rebase.rescheduleFailedExec configuration (see git-config(1) or "CONFIGURATION" below) or whether this option is provided.
// Otherwise an explicit --no-reschedule-failed-exec at the start would be overridden by the presence of rebase.rescheduleFailedExec=true configuration.
// --reschedule-failed-exec, --no-reschedule-failed-exec
func NoRescheduleFailedExec(g *types.Cmd) {
	g.AddOptions("--no-reschedule-failed-exec")
}

// NoStat Do not show a diffstat as part of the rebase process.
// -n, --no-stat
func NoStat(g *types.Cmd) {
	g.AddOptions("--no-stat")
}

// NoVerify This option bypasses the pre-rebase hook.
// See also githooks(5).
// --no-verify
func NoVerify(g *types.Cmd) {
	g.AddOptions("--no-verify")
}

// Onto Starting point at which to create the new commits.
// If the --onto option is not specified, the starting point is <upstream>.
// May be any valid commit, and not just an existing branch name.
// As a special case, you may use "A...B" as a shortcut for the merge base of A and B if there is exactly one merge base.
// You can leave out at most one of A and B, in which case it defaults to HEAD.
// --onto <newbase>
func Onto(newbase string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--onto")
		g.AddOptions(newbase)
	}
}

// Quiet Be quiet.
// Implies --no-stat.
// -q, --quiet
func Quiet(g *types.Cmd) {
	g.AddOptions("--quiet")
}

// Quit Abort the rebase operation but HEAD is not reset back to the original branch.
// The index and working tree are also left unchanged as a result.
// If a temporary stash entry was created using --autostash, it will be saved to the stash list.
// --quit
func Quit(g *types.Cmd) {
	g.AddOptions("--quit")
}

// ReapplyCherryPicks Reapply all clean cherry-picks of any upstream commit instead of preemptively dropping them.
// (If these commits then become empty after rebasing, because they contain a subset of already upstream changes, the behavior towards them is controlled by the --empty flag.)
//
// By default (or if --no-reapply-cherry-picks is given), these commits will be automatically dropped.
// Because this necessitates reading all upstream commits, this can be expensive in repos with a large number of upstream commits that need to be read.
// When using the merge backend, warnings will be issued for each dropped commit (unless --quiet is given).
// Advice will also be issued unless advice.skippedCherryPicks is set to false (see git-config(1)).
// --reapply-cherry-picks allows rebase to forgo reading all upstream commits, potentially improving performance.
// --reapply-cherry-picks, --no-reapply-cherry-picks
func ReapplyCherryPicks(g *types.Cmd) {
	g.AddOptions("--reapply-cherry-picks")
}

// RebaseMerges By default, a rebase will simply drop merge commits from the todo list, and put the rebased commits into a single, linear branch.
// With --rebase-merges, the rebase will instead try to preserve the branching structure within the commits that are to be rebased, by recreating the merge commits.
// Any resolved merge conflicts or manual amendments in these merge commits will have to be resolved/re-applied manually.
//
// By default, or when no-rebase-cousins was specified, commits which do not have <upstream> as direct ancestor will keep their original branch point, i.e. commits that would be excluded by git-log(1)'s --ancestry-path option will keep their original ancestry by default.
// If the rebase-cousins mode is turned on, such commits are instead rebased onto <upstream> (or <onto>, if specified).
//
// It is currently only possible to recreate the merge commits using the ort merge strategy; different merge strategies can be used only via explicit exec git merge -s <strategy> [...]  commands.
// -r, --rebase-merges[=(rebase-cousins|no-rebase-cousins)]
func RebaseMerges(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(value) == 0 {
			g.AddOptions("--rebase-merges")
		} else {
			g.AddOptions(fmt.Sprintf("--rebase-merges=%s", value))
		}
	}
}

// RerereAutoupdate Allow the rerere mechanism to update the index with the result of auto-conflict resolution if possible.
// --rerere-autoupdate, --no-rerere-autoupdate
func RerereAutoupdate(g *types.Cmd) {
	g.AddOptions("--rerere-autoupdate")
}

// RescheduleFailedExec Automatically reschedule exec commands that failed.
// This only makes sense in interactive mode (or when an --exec option was provided).
//
// Even though this option applies once a rebase is started, it’s set for the whole rebase at the start based on either the rebase.rescheduleFailedExec configuration (see git-config(1) or "CONFIGURATION" below) or whether this option is provided.
// Otherwise an explicit --no-reschedule-failed-exec at the start would be overridden by the presence of rebase.rescheduleFailedExec=true configuration.
// --reschedule-failed-exec, --no-reschedule-failed-exec
func RescheduleFailedExec(g *types.Cmd) {
	g.AddOptions("--reschedule-failed-exec")
}

// ResetAuthorDate Instead of using the author date of the original commit, use the current time as the author date of the rebased commit.
// This option implies --force-rebase.
// --ignore-date, --reset-author-date
func ResetAuthorDate(g *types.Cmd) {
	g.AddOptions("--reset-author-date")
}

// Root Rebase all commits reachable from <branch>, instead of limiting them with an <upstream>.
// This allows you to rebase the root commit(s) on a branch.
// When used with --onto, it will skip changes already contained in <newbase> (instead of <upstream>) whereas without
// --onto it will operate on every change.
// --root
func Root(g *types.Cmd) {
	g.AddOptions("--root")
}

// ShowCurrentPatch Show the current patch in an interactive rebase or when rebase is stopped because of conflicts.
// This is the equivalent of git show REBASE_HEAD.
// --show-current-patch
func ShowCurrentPatch(g *types.Cmd) {
	g.AddOptions("--show-current-patch")
}

// Signoff Add a Signed-off-by trailer to all the rebased commits.
// Note that if --interactive is given then only commits marked to be picked, edited or reworded will have the trailer added.
// --signoff
func Signoff(g *types.Cmd) {
	g.AddOptions("--signoff")
}

// Skip Restart the rebasing process by skipping the current patch.
// --skip
func Skip(g *types.Cmd) {
	g.AddOptions("--skip")
}

// Stat Show a diffstat of what changed upstream since the last rebase.
// The diffstat is also controlled by the configuration option rebase.stat.
// --stat
func Stat(g *types.Cmd) {
	g.AddOptions("--stat")
}

// Strategy Use the given merge strategy, instead of the default ort.
// This implies --merge.
//
// Because git rebase replays each commit from the working branch on top of the <upstream> branch using the given strategy, using the ours strategy simply empties all patches from the <branch>, which makes little sense.
// -s <strategy>, --strategy=<strategy>
func Strategy(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--strategy=%s", value))
	}
}

// StrategyOption Pass the <strategy-option> through to the merge strategy.
// This implies --merge and, if no strategy has been specified, -s ort.
// Note the reversal of ours and theirs as noted above for the -m option.
// -X <strategy-option>, --strategy-option=<strategy-option>
func StrategyOption(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--strategy-option=%s", value))
	}
}

// Verbose Be verbose.
// Implies --stat.
// -v, --verbose
func Verbose(g *types.Cmd) {
	g.AddOptions("--verbose")
}

// Verify Allows the pre-rebase hook to run, which is the default.
// This option can be used to override --no-verify.
// See also githooks(5).
// --verify
func Verify(g *types.Cmd) {
	g.AddOptions("--verify")
}

// Whitespace This flag is passed to the git apply program (see git-apply(1)) that applies the patch.
// Implies --apply.
// --whitespace=<option>
func Whitespace(option string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--whitespace=%s", option))
	}
}
