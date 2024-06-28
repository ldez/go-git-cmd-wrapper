/*
Package notes git-notes - Add or inspect object notes.

# SYNOPSIS

Reference: https://git-scm.com/docs/git-notes

	git notes [list [<object>]]
	git notes add [-f] [--allow-empty] [-F <file> | -m <msg> | (-c | -C) <object>] [<object>]
	git notes copy [-f] ( --stdin | <from-object> [<to-object>] )
	git notes append [--allow-empty] [-F <file> | -m <msg> | (-c | -C) <object>] [<object>]
	git notes edit [--allow-empty] [<object>]
	git notes show [<object>]
	git notes merge [-v | -q] [-s <strategy> ] <notes-ref>
	git notes merge --commit [-v | -q]
	git notes merge --abort [-v | -q]
	git notes remove [--ignore-missing] [--stdin] [<object>…]
	git notes prune [-n] [-v]
	git notes get-ref

# DESCRIPTION

Adds, removes, or reads notes attached to objects, without touching the objects themselves.

By default, notes are saved to and read from refs/notes/commits, but this default can be overridden. See the OPTIONS, CONFIGURATION, and ENVIRONMENT sections below.
If this ref does not exist, it will be quietly created when it is first needed to store a note.

A typical use of notes is to supplement a commit message without changing the commit itself. Notes can be shown by git log along with the original commit message.
To distinguish these notes from the message stored in the commit object, the notes are indented like the message, after an unindented line saying
"Notes (<refname>):" (or "Notes:" for refs/notes/commits).

Notes can also be added to patches prepared with git format-patch by using the --notes option. Such notes are added as a patch commentary after a three dash separator line.

To change which notes are shown by git log, see the "notes.displayRef" configuration in git-log[1].

See the "notes.rewrite.<command>" configuration for a way to carry notes across commands that rewrite commits.
*/
package notes
