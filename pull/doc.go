/*
Package pull git-pull - Fetch from and integrate with another repository or a local branch.

# SYNOPSIS

Reference:https://git-scm.com/docs/git-pull

	git pull [options] [<repository> [<refspec>...]]

# DESCRIPTION

Incorporates changes from a remote repository into the current branch. In its default mode, git pull is shorthand for git fetch followed by git merge FETCH_HEAD.

More precisely, git pull runs git fetch with the given parameters and calls git merge to merge the retrieved branch heads into the current branch. With --rebase, it runs git rebase instead of git merge.

<repository> should be the name of a remote repository as passed to git-fetch(1). <refspec> can name an arbitrary remote ref (for example, the name of a tag) or even a collection of refs with corresponding remote-tracking
branches (e.g., refs/heads/*:refs/remotes/origin/*), but usually it is the name of a branch in the remote repository.

Default values for <repository> and <branch> are read from the "remote" and "merge" configuration for the current branch as set by git-branch(1) --track.

Assume the following history exists and the current branch is "master":

				 A---B---C master on origin
				/
	 D---E---F---G master
			 ^
			 origin/master in your repository

Then "git pull" will fetch and replay the changes from the remote master branch since it diverged from the local master (i.e., E) until its current commit (C) on top of master and record the result in a new commit along with
the names of the two parent commits and a log message from the user describing the changes.

				 A---B---C origin/master
				/         \
	 D---E---F---G---H master

See git-merge(1) for details, including how conflicts are presented and handled.

In Git 1.7.0 or later, to cancel a conflicting merge, use git reset --merge. Warning: In older versions of Git, running git pull with uncommitted changes is discouraged: while possible, it leaves you in a state that may be
hard to back out of in the case of a conflict.

If any of the remote changes overlap with local uncommitted changes, the merge will be automatically canceled and the work tree untouched. It is generally best to get any local changes in working order before pulling or stash
them away with git-stash(1).
*/
package pull
