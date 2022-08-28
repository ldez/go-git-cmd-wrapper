/*
Package merge git-merge - Join two or more development histories together.

# SYNOPSIS

Reference: https://git-scm.com/docs/git-merge

	git merge [-n] [--stat] [--no-commit] [--squash] [--[no-]edit]
			[-s <strategy>] [-X <strategy-option>] [-S[<keyid>]]
			[--[no-]allow-unrelated-histories]
			[--[no-]rerere-autoupdate] [-m <msg>] [<commit>...]
	git merge --abort
	git merge --continue

# DESCRIPTION

Incorporates changes from the named commits (since the time their histories diverged from the current branch) into the current branch. This command is used by git pull to incorporate changes from another repository and can be used by hand to merge changes from one branch into another.

Assume the following history exists and the current branch is "master":

				 A---B---C topic
				/
	 D---E---F---G master

Then "git merge topic" will replay the changes made on the topic branch since it diverged from master (i.e., E) until its current commit (C) on top of master, and record the result in a new commit along with the names of the two parent commits and a log message from the user describing the changes.

				 A---B---C topic
				/         \
	 D---E---F---G---H master

The second syntax ("git merge --abort") can only be run after the merge has resulted in conflicts. git merge --abort will abort the merge process and try to reconstruct the pre-merge state. However, if there were uncommitted changes when the merge started (and especially if those changes were further modified after the merge was started), git merge --abort will in some cases be unable to reconstruct the original (pre-merge) changes. Therefore:

Warning: Running git merge with non-trivial uncommitted changes is discouraged: while possible, it may leave you in a state that is hard to back out of in the case of a conflict.

The fourth syntax ("git merge --continue") can only be run after the merge has resulted in conflicts.
*/
package merge
