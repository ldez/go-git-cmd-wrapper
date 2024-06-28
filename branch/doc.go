/*
Package branch git-branch - List, create, or delete branches.

# SYNOPSIS

Reference: https://git-scm.com/docs/git-branch

	git branch [--color[=<when>] | --no-color] [--show-current]
		[-v [--abbrev=<n> | --no-abbrev]]
		[--column[=<options>] | --no-column] [--sort=<key>]
		[--merged [<commit>]] [--no-merged [<commit>]]
		[--contains [<commit>]] [--no-contains [<commit>]]
		[--points-at <object>] [--format=<format>]
		[(-r | --remotes) | (-a | --all)]
		[--list] [<pattern>...]
	git branch [--track[=(direct|inherit)] | --no-track] [-f]
		[--recurse-submodules] <branchname> [<start-point>]
	git branch (--set-upstream-to=<upstream> | -u <upstream>) [<branchname>]
	git branch --unset-upstream [<branchname>]
	git branch (-m | -M) [<oldbranch>] <newbranch>
	git branch (-c | -C) [<oldbranch>] <newbranch>
	git branch (-d | -D) [-r] <branchname>…
	git branch --edit-description [<branchname>]
*/
package branch
