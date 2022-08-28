/*
Package push git-push - Update remote refs along with associated objects.

# SYNOPSIS

Reference: https://git-scm.com/docs/git-push

	git push [--all | --mirror | --tags] [--follow-tags] [--atomic] [-n | --dry-run] [--receive-pack=<git-receive-pack>]
		[--repo=<repository>] [-f | --force] [-d | --delete] [--prune] [-v | --verbose]
		[-u | --set-upstream] [--push-option=<string>]
		[--[no-]signed|--sign=(true|false|if-asked)]
		[--force-with-lease[=<refname>[:<expect>]]]
		[--no-verify] [<repository> [<refspec>...]]

# DESCRIPTION

Updates remote refs using local refs, while sending objects necessary to complete the given refs.

You can make interesting things happen to a repository every time you push into it, by setting up hooks there. See documentation for git-receive-pack(1).

When the command line does not specify where to push with the <repository> argument, branch.*.remote configuration for the current branch is consulted to determine where to push. If the configuration is missing, it defaults
to origin.

When the command line does not specify what to push with <refspec>... arguments or --all, --mirror, --tags options, the command finds the default <refspec> by consulting remote.*.push configuration, and if it is not found,
honors push.default configuration to decide what to push (See git-config(1) for the meaning of push.default).

When neither the command-line nor the configuration specify what to push, the default behavior is used, which corresponds to the simple value for push.default: the current branch is pushed to the corresponding upstream
branch, but as a safety measure, the push is aborted if the upstream branch does not have the same name as the local one.
*/
package push
