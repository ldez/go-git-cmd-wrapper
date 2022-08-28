/*
Package checkout git-checkout - Switch branches or restore working tree files.

# SYNOPSIS

Reference: https://git-scm.com/docs/git-checkout

	git checkout [-q] [-f] [-m] [<branch>]
	git checkout [-q] [-f] [-m] --detach [<branch>]
	git checkout [-q] [-f] [-m] [--detach] <commit>
	git checkout [-q] [-f] [-m] [[-b|-B|--orphan] <new_branch>] [<start_point>]
	git checkout [-f|--ours|--theirs|-m|--conflict=<style>] [<tree-ish>] [--] <paths>...
	git checkout [-p|--patch] [<tree-ish>] [--] [<paths>...]

# DESCRIPTION

Updates files in the working tree to match the version in the index or the specified tree. If no paths are given, git checkout will also update HEAD to set the specified branch as the current branch.
*/
package checkout
