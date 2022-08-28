/*
Package rebase git-rebase - Reapply commits on top of another base tip.

# SYNOPSIS

Reference: https://git-scm.com/docs/git-rebase

	 git rebase [-i | --interactive] [<options>] [--exec <cmd>]
					[--onto <newbase> | --keep-base] [<upstream> [<branch>]]
	 git rebase [-i | --interactive] [<options>] [--exec <cmd>] [--onto <newbase>]
					 --root [<branch>]
	 git rebase (--continue | --skip | --abort | --quit | --edit-todo | --show-current-patch)

# DESCRIPTION

If <branch> is specified, git rebase will perform an automatic git switch <branch> before doing anything else. Otherwise it remains on the current branch.

If <upstream> is not specified, the upstream configured in branch.<name>.remote and branch.<name>.merge options will be used (see git-config(1) for details) and the --fork-point option is assumed.
If you are currently not on any branch or if the current branch does not have a configured upstream, the rebase will abort.

All changes made by commits in the current branch but that are not in <upstream> are saved to a temporary area.
This is the same set of commits that would be shown by git log <upstream>..HEAD; or by git log 'fork_point'..HEAD, if --fork-point is active (see the description on --fork-point below); or by git log HEAD, if the --root option is specified.

The current branch is reset to <upstream>, or <newbase> if the --onto option was supplied.
This has the exact same effect as git reset --hard <upstream> (or <newbase>). ORIG_HEAD is set to point at the tip of the branch before the reset.

The commits that were previously saved into the temporary area are then reapplied to the current branch, one by one, in order.
Note that any commits in HEAD which introduce the same textual changes as a commit in HEAD..<upstream> are omitted (i.e., a patch already accepted upstream with a different commit message or timestamp will be skipped).

It is possible that a merge failure will prevent this process from being completely automatic.
You will have to resolve any such merge failure and run git rebase --continue. Another option is to bypass the commit that caused the merge failure with git rebase --skip.
To check out the original <branch> and remove the .git/rebase-apply working files, use the command git rebase --abort instead.
*/
package rebase
