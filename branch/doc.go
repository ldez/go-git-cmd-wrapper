/*
Package branch git-branch - List, create, or delete branches.

# SYNOPSIS

Reference: https://git-scm.com/docs/git-branch

	git branch [--color[=<when>] | --no-color] [-r | -a]
		[--list] [-v [--abbrev=<length> | --no-abbrev]]
		[--column[=<options>] | --no-column] [--sort=<key>]
		[(--merged | --no-merged) [<commit>]]
		[--contains [<commit]] [--no-contains [<commit>]]
		[--points-at <object>] [--format=<format>] [<pattern>…​]
	git branch [--track | --no-track] [-l] [-f] <branchname> [<start-point>]
	git branch (--set-upstream-to=<upstream> | -u <upstream>) [<branchname>]
	git branch --unset-upstream [<branchname>]
	git branch (-m | -M) [<oldbranch>] <newbranch>
	git branch (-c | -C) [<oldbranch>] <newbranch>
	git branch (-d | -D) [-r] <branchname>…​
	git branch --edit-description [<branchname>]
*/
package branch
