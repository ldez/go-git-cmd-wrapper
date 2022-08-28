/*
Package init git-init - Create an empty Git repository or reinitialize an existing one.

# SYNOPSIS

Reference: https://git-scm.com/docs/git-init

	git init [-q | --quiet] [--bare] [--template=<template_directory>]
					 [--separate-git-dir <git dir>]
					 [--shared[=<permissions>]] [directory]

# DESCRIPTION

This command creates an empty Git repository - basically a .git directory with subdirectories for objects, refs/heads, refs/tags, and template files. An initial HEAD file that references the HEAD of the master branch is also
created.

If the $GIT_DIR environment variable is set then it specifies a path to use instead of ./.git for the base of the repository.

If the object storage directory is specified via the $GIT_OBJECT_DIRECTORY environment variable then the sha1 directories are created underneath - otherwise the default $GIT_DIR/objects directory is used.

Running git init in an existing repository is safe. It will not overwrite things that are already there. The primary reason for rerunning git init is to pick up newly added templates (or to move the repository to another
place if --separate-git-dir is given).
*/
package init
