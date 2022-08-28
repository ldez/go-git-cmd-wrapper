/*
Package worktree git-worktree - Manage multiple working trees.

# SYNOPSIS

Reference: https://git-scm.com/docs/git-worktree

	git worktree add [-f] [--detach] [--checkout] [--lock] [-b <new-branch>] <path> [<branch>]
	git worktree list [--porcelain]
	git worktree lock [--reason <string>] <worktree>
	git worktree prune [-n] [-v] [--expire <expire>]
	git worktree unlock <worktree>

# DESCRIPTION

Manage multiple working trees attached to the same repository.
*/
package worktree
