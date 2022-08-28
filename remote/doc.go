/*
Package remote git-remote - Manage set of tracked repositories.

# SYNOPSIS

Reference: https://git-scm.com/docs/git-remote

	git remote [-v | --verbose]
	git remote add [-t <branch>] [-m <master>] [-f] [--[no-]tags] [--mirror=<fetch|push>] <name> <url>
	git remote rename <old> <new>
	git remote remove <name>
	git remote set-head <name> (-a | --auto | -d | --delete | <branch>)
	git remote set-branches [--add] <name> <branch>...
	git remote get-url [--push] [--all] <name>
	git remote set-url [--push] <name> <newurl> [<oldurl>]
	git remote set-url --add [--push] <name> <newurl>
	git remote set-url --delete [--push] <name> <url>
	git remote [-v | --verbose] show [-n] <name>...
	git remote prune [-n | --dry-run] <name>...
	git remote [-v | --verbose] update [-p | --prune] [(<group> | <remote>)...]

# DESCRIPTION

Manage the set of repositories ("remotes") whose branches you track.
*/
package remote
