package merge

// CODE GENERATED AUTOMATICALLY
// THIS FILE MUST NOT BE EDITED BY HAND

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/types"
)

// Abort Abort the current conflict resolution process, and try to reconstruct the pre-merge state.
// If there were uncommitted worktree changes present when the merge started, git merge --abort will in some cases be unable to reconstruct these changes.
// It is therefore recommended to always commit or stash your changes before running git merge.
// git merge --abort is equivalent to git reset --merge when MERGE_HEAD is present.
// --abort
func Abort(g *types.Cmd) {
	g.AddOptions("--abort")
}

// AllowUnrelatedHistories By default, git merge command refuses to merge histories that do not share a common ancestor.
// This option can be used to override this safety when merging histories of two projects that started their lives independently.
// As that is a very rare occasion, no configuration variable to enable this by default exists and will not be added.
// --allow-unrelated-histories
func AllowUnrelatedHistories(g *types.Cmd) {
	g.AddOptions("--allow-unrelated-histories")
}

// Commit Perform the merge and commit the result.
// This option can be used to override --no-commit.
// --commit, --no-commit
func Commit(g *types.Cmd) {
	g.AddOptions("--commit")
}

// Continue After a git merge stops due to conflicts you can conclude the merge by running git merge --continue (see 'HOW TO RESOLVE CONFLICTS' section below).
// --continue
func Continue(g *types.Cmd) {
	g.AddOptions("--continue")
}

// Edit Invoke an editor before committing successful mechanical merge to further edit the auto-generated merge message, so that the user can explain and justify the merge.
// The --no-edit option can be used to accept the auto-generated message (this is generally discouraged).
// The --edit (or -e) option is still useful if you are giving a draft message with the -m option from the command line and want to edit it in the editor.
// Older scripts may depend on the historical behaviour of not allowing the user to edit the merge log message.
// They will see an editor opened when they run git merge.
// To make it easier to adjust such scripts to the updated behaviour, the environment variable GIT_MERGE_AUTOEDIT can be set to no at the beginning of them.
// --edit, -e, --no-edit
func Edit(g *types.Cmd) {
	g.AddOptions("--edit")
}

// Ff When the merge resolves as a fast-forward, only update the branch pointer, without creating a merge commit.
// This is the default behavior.
// --ff
func Ff(g *types.Cmd) {
	g.AddOptions("--ff")
}

// FfOnly Refuse to merge and exit with a non-zero status unless the current HEAD is already up-to-date or the merge can be resolved as a fast-forward.
// --ff-only
func FfOnly(g *types.Cmd) {
	g.AddOptions("--ff-only")
}

// GpgSign GPG-sign the resulting merge commit.
// The keyid argument is optional and defaults to the committer identity; if specified, it must be stuck to the option without a space.
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

// Log In addition to branch names, populate the log message with one-line descriptions from at most <n> actual commits that are being merged.
// See also git-fmt-merge-msg(1).
// With --no-log do not list one-line descriptions from the actual commits being merged.
// --log[=<n>], --no-log
func Log(n string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(n) == 0 {
			g.AddOptions("--log")
		} else {
			g.AddOptions(fmt.Sprintf("--log=%s", n))
		}
	}
}

// M Set the commit message to be used for the merge commit (in case one is created).
// If --log is specified, a shortlog of the commits being merged will be appended to the specified message.
// The git fmt-merge-msg command can be used to give a good default for automated git merge invocations.
// The automated message can include the branch description.
// -m <msg>
func M(msg string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("-m")
		g.AddOptions(msg)
	}
}

// NoCommit With --no-commit perform the merge but pretend the merge failed and do not autocommit, to give the user a chance to inspect and further tweak the merge result before committing.
// --commit, --no-commit
func NoCommit(g *types.Cmd) {
	g.AddOptions("--no-commit")
}

// NoEdit Invoke an editor before committing successful mechanical merge to further edit the auto-generated merge message, so that the user can explain and justify the merge.
// The --no-edit option can be used to accept the auto-generated message (this is generally discouraged).
// The --edit (or -e) option is still useful if you are giving a draft message with the -m option from the command line and want to edit it in the editor.
// Older scripts may depend on the historical behaviour of not allowing the user to edit the merge log message.
// They will see an editor opened when they run git merge.
// To make it easier to adjust such scripts to the updated behaviour, the environment variable GIT_MERGE_AUTOEDIT can be set to no at the beginning of them.
// --edit, -e, --no-edit
func NoEdit(g *types.Cmd) {
	g.AddOptions("--no-edit")
}

// NoFf Create a merge commit even when the merge resolves as a fast-forward.
// This is the default behaviour when merging an annotated (and possibly signed) tag.
// --no-ff
func NoFf(g *types.Cmd) {
	g.AddOptions("--no-ff")
}

// NoLog In addition to branch names, populate the log message with one-line descriptions from at most <n> actual commits that are being merged.
// See also git-fmt-merge-msg(1).
// With --no-log do not list one-line descriptions from the actual commits being merged.
// --log[=<n>], --no-log
func NoLog(g *types.Cmd) {
	g.AddOptions("--no-log")
}

// NoProgress Turn progress on/off explicitly.
// If neither is specified, progress is shown if standard error is connected to a terminal.
// Note that not all merge strategies may support progress reporting.
// --progress, --no-progress
func NoProgress(g *types.Cmd) {
	g.AddOptions("--no-progress")
}

// NoRerereAutoupdate Allow the rerere mechanism to update the index with the result of auto-conflict resolution if possible.
// --[no-]rerere-autoupdate
func NoRerereAutoupdate(g *types.Cmd) {
	g.AddOptions("--no-rerere-autoupdate")
}

// NoSquash Produce the working tree and index state as if a real merge happened (except for the merge information), but do not actually make a commit, move the HEAD, or record $GIT_DIR/MERGE_HEAD (to cause the next git commit command to create a merge commit).
// This allows you to create a single commit on top of the current branch whose effect is the same as merging another branch (or more in case of an octopus).
// With --no-squash perform the merge and commit the result.
// This option can be used to override --squash.
// --squash, --no-squash
func NoSquash(g *types.Cmd) {
	g.AddOptions("--no-squash")
}

// NoStat Show a diffstat at the end of the merge.
// The diffstat is also controlled by the configuration option merge.stat.
// With -n or --no-stat do not show a diffstat at the end of the merge.
// --stat, -n, --no-stat
func NoStat(g *types.Cmd) {
	g.AddOptions("--no-stat")
}

// NoSummary Synonyms to --stat and --no-stat; these are deprecated and will be removed in the future.
// --summary, --no-summary
func NoSummary(g *types.Cmd) {
	g.AddOptions("--no-summary")
}

// NoVerifySignatures Verify that the tip commit of the side branch being merged is signed with a valid key, i.e. a key that has a valid uid: in the default trust model, this means the signing key has been signed by a trusted key.
// If the tip commit of the side branch is not signed with a valid key, the merge is aborted.
// --verify-signatures, --no-verify-signatures
func NoVerifySignatures(g *types.Cmd) {
	g.AddOptions("--no-verify-signatures")
}

// Progress Turn progress on/off explicitly.
// If neither is specified, progress is shown if standard error is connected to a terminal.
// Note that not all merge strategies may support progress reporting.
// --progress, --no-progress
func Progress(g *types.Cmd) {
	g.AddOptions("--progress")
}

// Quiet Operate quietly.
// Implies --no-progress.
// -q, --quiet
func Quiet(g *types.Cmd) {
	g.AddOptions("--quiet")
}

// RerereAutoupdate Allow the rerere mechanism to update the index with the result of auto-conflict resolution if possible.
// --[no-]rerere-autoupdate
func RerereAutoupdate(g *types.Cmd) {
	g.AddOptions("--rerere-autoupdate")
}

// Squash Produce the working tree and index state as if a real merge happened (except for the merge information), but do not actually make a commit, move the HEAD, or record $GIT_DIR/MERGE_HEAD (to cause the next git commit command to create a merge commit).
// This allows you to create a single commit on top of the current branch whose effect is the same as merging another branch (or more in case of an octopus).
// With --no-squash perform the merge and commit the result.
// This option can be used to override --squash.
// --squash, --no-squash
func Squash(g *types.Cmd) {
	g.AddOptions("--squash")
}

// Stat Show a diffstat at the end of the merge.
// The diffstat is also controlled by the configuration option merge.stat.
// With -n or --no-stat do not show a diffstat at the end of the merge.
// --stat, -n, --no-stat
func Stat(g *types.Cmd) {
	g.AddOptions("--stat")
}

// Strategy Use the given merge strategy; can be supplied more than once to specify them in the order they should be tried.
// If there is no -s option, a built-in list of strategies is used instead (git merge-recursive when merging a single head, git merge-octopus otherwise).
// -s <strategy>, --strategy=<strategy>
func Strategy(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--strategy=%s", value))
	}
}

// StrategyOption Pass merge strategy specific option through to the merge strategy.
// -X <option>, --strategy-option=<option>
func StrategyOption(option string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--strategy-option=%s", option))
	}
}

// Summary Synonyms to --stat and --no-stat; these are deprecated and will be removed in the future.
// --summary, --no-summary
func Summary(g *types.Cmd) {
	g.AddOptions("--summary")
}

// Verbose Be verbose.
// -v, --verbose
func Verbose(g *types.Cmd) {
	g.AddOptions("--verbose")
}

// VerifySignatures Verify that the tip commit of the side branch being merged is signed with a valid key, i.e. a key that has a valid uid: in the default trust model, this means the signing key has been signed by a trusted key.
// If the tip commit of the side branch is not signed with a valid key, the merge is aborted.
// --verify-signatures, --no-verify-signatures
func VerifySignatures(g *types.Cmd) {
	g.AddOptions("--verify-signatures")
}
