/*
Package status git-status - Show the working tree status

# SYNOPSIS

Reference: https://git-scm.com/docs/git-status

	git status [<options>...] [--] [<pathspec>...]

# DESCRIPTION

Displays paths that have differences between the index file and the current HEAD commit, paths that have differences between the working tree and the index file, and paths in the working tree
that are not tracked by Git (and are not ignored by gitignore(5)). The first are what you would commit by running git commit; the second and third are what you could commit by running git add
before running git commit.
*/
package status
