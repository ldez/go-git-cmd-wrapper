/*
Package lsfiles git-ls-files - Show information about files in the index and the working tree.

# SYNOPSIS

Reference: https://git-scm.com/docs/git-ls-files

	git ls-files [-z] [-t] [-v] [-f]
								 [-c|--cached] [-d|--deleted] [-o|--others] [-i|--ignored]
								 [-s|--stage] [-u|--unmerged] [-k|--killed] [-m|--modified]
								 [--directory [--no-empty-directory]] [--eol]
								 [--deduplicate]
								 [-x <pattern>|--exclude=<pattern>]
								 [-X <file>|--exclude-from=<file>]
								 [--exclude-per-directory=<file>]
								 [--exclude-standard]
								 [--error-unmatch] [--with-tree=<tree-ish>]
								 [--full-name] [--recurse-submodules]
								 [--abbrev[=<n>]] [--format=<format>] [--] [<file>...]

# DESCRIPTION

This merges the file listing in the index with the actual working directory list, and shows different combinations of the two.
One or more of the options below may be used to determine the files shown:
*/
package lsfiles
