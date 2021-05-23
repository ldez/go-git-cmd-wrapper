package notes

// CODE GENERATED AUTOMATICALLY
// THIS FILE MUST NOT BE EDITED BY HAND

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/v2/types"
)

// Abort Abort/reset an in-progress git notes merge, i.e. a notes merge with conflicts.
// This simply removes all files related to the notes merge.
// --abort
func Abort(g *types.Cmd) {
	g.AddOptions("--abort")
}

// AllowEmpty Allow an empty note object to be stored. The default behavior is to automatically remove empty notes.
// --allow-empty
func AllowEmpty(g *types.Cmd) {
	g.AddOptions("--allow-empty")
}

// Commit Finalize an in-progress git notes merge.
// Use this option when you have resolved the conflicts that git notes merge stored in .git/NOTES_MERGE_WORKTREE.
// This amends the partial merge commit created by git notes merge (stored in .git/NOTES_MERGE_PARTIAL) by adding the notes in .git/NOTES_MERGE_WORKTREE.
// The notes ref stored in the .git/NOTES_MERGE_REF symref is updated to the resulting commit.
// --commit
func Commit(g *types.Cmd) {
	g.AddOptions("--commit")
}

// DryRun Do not remove anything; just report the object names whose notes would be removed.
// -n, --dry-run
func DryRun(g *types.Cmd) {
	g.AddOptions("--dry-run")
}

// File Take the note message from the given file.
// Use - to read the note message from the standard input.
// Lines starting with # and empty lines other than a single line between paragraphs will be stripped out.
// -F <file>, --file=<file>
func File(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--file=%s", value))
	}
}

// Force When adding notes to an object that already has notes, overwrite the existing notes (instead of aborting).
// -f, --force
func Force(g *types.Cmd) {
	g.AddOptions("--force")
}

// IgnoreMissing Do not consider it an error to request removing notes from an object that does not have notes attached to it.
// --ignore-missing
func IgnoreMissing(g *types.Cmd) {
	g.AddOptions("--ignore-missing")
}

// Message Use the given note message (instead of prompting).
// If multiple -m options are given, their values are concatenated as separate paragraphs.
// Lines starting with # and empty lines other than a single line between paragraphs will be stripped out.
// -m <msg>, --message=<msg>
func Message(msg string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--message=%s", msg))
	}
}

// Quiet When merging notes, operate quietly.
// -q, --quiet
func Quiet(g *types.Cmd) {
	g.AddOptions("--quiet")
}

// ReeditMessage Like -C, but with -c the editor is invoked, so that the user can further edit the note message.
// -c <object>, --reedit-message=<object>
func ReeditMessage(object string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--reedit-message=%s", object))
	}
}

// Ref Manipulate the notes tree in <ref>. This overrides GIT_NOTES_REF and the "core.notesRef" configuration.
// The ref specifies the full refname when it begins with refs/notes/;
// when it begins with notes/, refs/ and otherwise refs/notes/ is prefixed to form a full name of the ref.
// --ref <ref>
func Ref(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--ref")
		g.AddOptions(value)
	}
}

// ReuseMessage Take the given blob object (for example, another note) as the note message.
// (Use git notes copy <object> instead to copy notes between objects.)
// -C <object>, --reuse-message=<object>
func ReuseMessage(object string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--reuse-message=%s", object))
	}
}

// Stdin Also read the object names to remove notes from the standard input (there is no reason you cannot combine this with object names from the command line).
// --stdin
func Stdin(g *types.Cmd) {
	g.AddOptions("--stdin")
}

// Strategy When merging notes, resolve notes conflicts using the given strategy.
// The following strategies are recognized: "manual" (default), "ours", "theirs", "union" and "cat_sort_uniq".
// This option overrides the "notes.mergeStrategy" configuration setting.
// See the "NOTES MERGE STRATEGIES" section below for more information on each notes merge strategy.
// -s <strategy>, --strategy=<strategy>
func Strategy(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--strategy=%s", value))
	}
}

// Verbose When merging notes, be more verbose.
// When pruning notes, report all object names whose notes are removed.
// -v, --verbose
func Verbose(g *types.Cmd) {
	g.AddOptions("--verbose")
}
