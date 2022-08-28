/*
Package clone git-clone - Clone a repository into a new directory.

# SYNOPSIS

Reference: https://git-scm.com/docs/git-clone

	git clone [--template=<template_directory>]
					 [-l] [-s] [--no-hardlinks] [-q] [-n] [--bare] [--mirror]
					 [-o <name>] [-b <name>] [-u <upload-pack>] [--reference <repository>]
					 [--dissociate] [--separate-git-dir <git dir>]
					 [--depth <depth>] [--[no-]single-branch]
					 [--recurse-submodules] [--[no-]shallow-submodules]
					 [--jobs <n>] [--] <repository> [<directory>]

# DESCRIPTION

Clones a repository into a newly created directory, creates remote-tracking branches for each branch in the cloned repository (visible using git branch -r), and creates and checks out an initial branch that is forked from the cloned repository’s currently active branch.

After the clone, a plain git fetch without arguments will update all the remote-tracking branches, and a git pull without arguments will in addition merge the remote master branch into the current master branch, if any (this is untrue when "--single-branch" is given; see below).

This default configuration is achieved by creating references to the remote branch heads under refs/remotes/origin and by initializing remote.origin.url and remote.origin.fetch configuration variables.
*/
package clone
