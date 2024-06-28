/*
Package stash git-stash - Stash the changes in a dirty working directory away.

# SYNOPSIS

Reference: https://git-scm.com/docs/git-stash

	git stash list [<log-options>]
	git stash show [-u | --include-untracked | --only-untracked] [<diff-options>] [<stash>]
	git stash drop [-q | --quiet] [<stash>]
	git stash pop [--index] [-q | --quiet] [<stash>]
	git stash apply [--index] [-q | --quiet] [<stash>]
	git stash branch <branchname> [<stash>]
	git stash [push [-p | --patch] [-S | --staged] [-k | --[no-]keep-index] [-q | --quiet]
				 [-u | --include-untracked] [-a | --all] [(-m | --message) <message>]
				 [--pathspec-from-file=<file> [--pathspec-file-nul]]
				 [--] [<pathspec>…​]]
	git stash save [-p | --patch] [-S | --staged] [-k | --[no-]keep-index] [-q | --quiet]
				 [-u | --include-untracked] [-a | --all] [<message>]
	git stash clear
	git stash create [<message>]
	git stash store [(-m | --message) <message>] [-q | --quiet] <commit>

# DESCRIPTION

Use `git stash` when you want to record the current state of the working directory and the index, but want to go back to a clean working directory. The command saves your local modifications away and reverts the working directory to match the `HEAD` commit.

The modifications stashed away by this command can be listed with `git stash list`, inspected with `git stash show`, and restored (potentially on top of a different commit) with `git stash apply`. Calling `git stash` without any arguments is equivalent to `git stash push`. A stash is by default listed as "WIP on branchname …", but you can give a more descriptive message on the command line when you create one.

The latest stash you created is stored in `refs/stash`; older stashes are found in the reflog of this reference and can be named using the usual reflog syntax (e.g. `stash@{0}` is the most recently created stash, `stash@{1}` is the one before it, `stash@{2.hours.ago}` is also possible). Stashes may also be referenced by specifying just the stash index (e.g. the integer `n` is equivalent to `stash@{n}`).
*/
package stash
