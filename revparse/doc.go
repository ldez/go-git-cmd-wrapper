/*
Package revparse git-rev-parse - Pick out and massage parameters.

# SYNOPSIS

Reference: https://git-scm.com/docs/git-rev-parse

	git rev-parse [ --option ] <args>...

# DESCRIPTION

Many Git porcelainish commands take mixture of flags (i.e. parameters that begin with a dash -) and parameters meant for the underlying git rev-list command they use internally and flags and parameters for the other commands
they use downstream of git rev-list. This command is used to distinguish between them.
*/
package revparse
