/*
Package tag git-tag - git-tag - Create, list, delete or verify a tag object signed with GPG.

# SYNOPSIS

Reference: https://git-scm.com/docs/git-tag

	git tag [-a | -s | -u <keyid>] [-f] [-m <msg> | -F <file>] [-e]
		<tagname> [<commit> | <object>]
	git tag -d <tagname>…
	git tag [-n[<num>]] -l [--contains <commit>] [--no-contains <commit>]
		[--points-at <object>] [--column[=<options>] | --no-column]
		[--create-reflog] [--sort=<key>] [--format=<format>]
		[--[no-]merged [<commit>]] [<pattern>…]
	git tag -v [--format=<format>] <tagname>…

# DESCRIPTION

Add a tag reference in refs/tags/, unless -d/-l/-v is given to delete, list or verify tags.
*/
package tag
