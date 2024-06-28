package stash

// CODE GENERATED AUTOMATICALLY
// THIS FILE MUST NOT BE EDITED BY HAND

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/v2/types"
)

// All This option is only valid for push and save commands.
// All ignored and untracked files are also stashed and then cleaned up with git clean.
// -a, --all
func All(g *types.Cmd) {
	g.AddOptions("--all")
}

// IncludeUntracked When used with the push and save commands, all untracked files are also stashed and then cleaned up with git clean.
// When used with the show command, show the untracked files in the stash entry as part of the diff.
// -u, --include-untracked
func IncludeUntracked(g *types.Cmd) {
	g.AddOptions("--include-untracked")
}

// Index This option is only valid for pop and apply commands.
// Tries to reinstate not only the working tree’s changes, but also the index’s ones.
// However, this can fail, when you have conflicts (which are stored in the index, where you therefore can no longer apply the changes as they were originally).
// --index
func Index(g *types.Cmd) {
	g.AddOptions("--index")
}

// KeepIndex This option is only valid for push and save commands.
// All changes already added to the index are left intact.
// -k, --keep-index
func KeepIndex(g *types.Cmd) {
	g.AddOptions("--keep-index")
}

// NoIncludeUntracked Opposite of --include-untracked.
// --no-include-untracked
func NoIncludeUntracked(g *types.Cmd) {
	g.AddOptions("--no-include-untracked")
}

// NoKeepIndex Opposite of --keep-index.
// --no-keep-index
func NoKeepIndex(g *types.Cmd) {
	g.AddOptions("--no-keep-index")
}

// OnlyUntracked This option is only valid for the show command.
// Show only the untracked files in the stash entry as part of the diff.
// --only-untracked
func OnlyUntracked(g *types.Cmd) {
	g.AddOptions("--only-untracked")
}

// Patch This option is only valid for push and save commands.
// Interactively select hunks from the diff between HEAD and the working tree to be stashed.
// The stash entry is constructed such that its index state is the same as the index state of your repository, and its worktree contains only the changes you selected interactively.
// The selected changes are then rolled back from your worktree.
// See the “Interactive Mode” section of git-add(1) to learn how to operate the --patch mode.
// The --patch option implies --keep-index.
// You can use --no-keep-index to override this.
// -p, --patch
func Patch(g *types.Cmd) {
	g.AddOptions("--patch")
}

// PathspecFileNul This option is only valid for push command.
// Only meaningful with --pathspec-from-file. Pathspec elements are separated with NUL character and all other characters are taken literally (including newlines and quotes).
// --pathspec-file-nul
func PathspecFileNul(g *types.Cmd) {
	g.AddOptions("--pathspec-file-nul")
}

// PathspecFromFile This option is only valid for push command.
// Pathspec is passed in <file> instead of commandline args.
// If <file> is exactly - then standard input is used. Pathspec elements are separated by LF or CR/LF.
// Pathspec elements can be quoted as explained for the configuration variable core.quotePath (see git-config(1)).
// See also --pathspec-file-nul and global --literal-pathspecs.
// --pathspec-from-file=<file>
func PathspecFromFile(file string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--pathspec-from-file=%s", file))
	}
}

// Quiet This option is only valid for apply, drop, pop, push, save, store commands.
// Quiet, suppress feedback messages.
// -q, --quiet
func Quiet(g *types.Cmd) {
	g.AddOptions("--quiet")
}

// Staged This option is only valid for push and save commands.
// Stash only the changes that are currently staged.
// This is similar to basic git commit except the state is committed to the stash instead of current branch.
// The --patch option has priority over this one.
// -S, --staged
func Staged(g *types.Cmd) {
	g.AddOptions("--staged")
}
