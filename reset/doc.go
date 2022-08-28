/*
Package reset git-reset - Reset current HEAD to the specified state.

# SYNOPSIS

Reference:  https://git-scm.com/docs/git-reset

	git reset [-q] [<tree-ish>] [--] <paths>...
	git reset (--patch | -p) [<tree-ish>] [--] [<paths>...]
	git reset [--soft | --mixed [-N] | --hard | --merge | --keep] [-q] [<commit>]

# DESCRIPTION

In the first and second form, copy entries from <tree-ish> to the index. In the third form, set the current branch head (HEAD) to <commit>, optionally modifying index and working tree to match. The <tree-ish>/<commit> defaults to HEAD in all forms.

	git reset [-q] [<tree-ish>] [--] <paths>...

This form resets the index entries for all <paths> to their state at <tree-ish>. (It does not affect the working tree or the current branch.)

This means that git reset <paths> is the opposite of git add <paths>.

After running git reset <paths> to update the index entry, you can use git-checkout(1) to check the contents out of the index to the working tree. Alternatively, using git-checkout(1) and specifying a commit, you can copy
the contents of a path out of a commit to the index and to the working tree in one go.

	git reset (--patch | -p) [<tree-ish>] [--] [<paths>...]

Interactively select hunks in the difference between the index and <tree-ish> (defaults to HEAD). The chosen hunks are applied in reverse to the index.

This means that git reset -p is the opposite of git add -p, i.e. you can use it to selectively reset hunks. See the “Interactive Mode” section of git-add(1) to learn how to operate the --patch mode.

	git reset [<mode>] [<commit>]

This form resets the current branch head to <commit> and possibly updates the index (resetting it to the tree of <commit>) and the working tree depending on <mode>. If <mode> is omitted, defaults to "--mixed".
*/
package reset
