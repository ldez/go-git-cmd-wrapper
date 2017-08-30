/*
Package fetch CODE GENERATED AUTOMATICALLY
THIS FILE MUST NOT BE EDITED BY HAND
*/
package fetch

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/types"
)

// All Fetch all remotes. 
// --all
func All(g *types.Cmd) {
	g.AddOptions("--all")
}

// Append Append ref names and object names of fetched refs to the existing contents of .git/FETCH_HEAD. 
// Without this option old data in .git/FETCH_HEAD will be overwritten. 
// -a, --append
func Append(g *types.Cmd) {
	g.AddOptions("--append")
}

// Deepen Similar to --depth, except it specifies the number of commits from the current shallow boundary instead of from the tip of each remote branch history. 
// --deepen=<depth>
func Deepen(depth string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--deepen=%s", depth))
	}
}

// Depth Limit fetching to the specified number of commits from the tip of each remote branch history. 
// If fetching to a shallow repository created by git clone with --depth=<depth> option (see git-clone(1)), deepen or shorten the history to the specified number of commits. 
// Tags for the deepened commits are not fetched. 
// --depth=<depth>
func Depth(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--depth=%s", value))
	}
}

// DryRun Show what would be done, without making any changes. 
// --dry-run
func DryRun(g *types.Cmd) {
	g.AddOptions("--dry-run")
}

// Force When git fetch is used with <rbranch>:<lbranch> refspec, it refuses to update the local branch <lbranch> unless the remote branch <rbranch> it fetches is a descendant of <lbranch>. 
// This option overrides that check. 
// -f, --force
func Force(g *types.Cmd) {
	g.AddOptions("--force")
}

// Ipv4 Use IPv4 addresses only, ignoring IPv6 addresses. 
// -4, --ipv4
func Ipv4(g *types.Cmd) {
	g.AddOptions("--ipv4")
}

// Ipv6 Use IPv6 addresses only, ignoring IPv4 addresses. 
// -6, --ipv6
func Ipv6(g *types.Cmd) {
	g.AddOptions("--ipv6")
}

// Jobs Number of parallel children to be used for fetching submodules. 
// Each will fetch from different submodules, such that fetching many submodules will be faster. 
// By default submodules will be fetched one at a time. 
// -j, --jobs=<n>
func Jobs(n string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--jobs=%s", n))
	}
}

// Keep Keep downloaded pack. 
// -k, --keep
func Keep(g *types.Cmd) {
	g.AddOptions("--keep")
}

// Multiple Allow several <repository> and <group> arguments to be specified. 
// No <refspec>s may be specified. 
// --multiple
func Multiple(g *types.Cmd) {
	g.AddOptions("--multiple")
}

// NoRecurseSubmodules Disable recursive fetching of submodules (this has the same effect as using the --recurse-submodules=no option). 
// --no-recurse-submodules
func NoRecurseSubmodules(g *types.Cmd) {
	g.AddOptions("--no-recurse-submodules")
}

// NoTags By default, tags that point at objects that are downloaded from the remote repository are fetched and stored locally. 
// This option disables this automatic tag following. 
// The default behavior for a remote may be specified with the remote.<name>.tagOpt setting. 
// -n, --no-tags
func NoTags(g *types.Cmd) {
	g.AddOptions("--no-tags")
}

// Progress Progress status is reported on the standard error stream by default when it is attached to a terminal, unless -q is specified. 
// This flag forces progress status even if the standard error stream is not directed to a terminal. 
// --progress
func Progress(g *types.Cmd) {
	g.AddOptions("--progress")
}

// Prune Before fetching, remove any remote-tracking references that no longer exist on the remote. 
// -p, --prune
func Prune(g *types.Cmd) {
	g.AddOptions("--prune")
}

// Quiet Pass --quiet to git-fetch-pack and silence any other internally used git commands. 
// Progress is not reported to the standard error stream. 
// -q, --quiet
func Quiet(g *types.Cmd) {
	g.AddOptions("--quiet")
}

// RecurseSubmodules This option controls if and under what conditions new commits of populated submodules should be fetched too. 
// --recurse-submodules[=yes|on-demand|no]
func RecurseSubmodules(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(value) == 0 {
			g.AddOptions("--recurse-submodules")
		} else {
			g.AddOptions(fmt.Sprintf("--recurse-submodules=%s", value))
		}
	}
}

// RecurseSubmodulesDefault This option is used internally to temporarily provide a non-negative default value for the --recurse-submodules option. 
// All other methods of configuring fetch’s submodule recursion (such as settings in gitmodules(5) and git-config(1)) override this option, as does specifying --[no-]recurse-submodules directly. 
// --recurse-submodules-default=[yes|on-demand]
func RecurseSubmodulesDefault(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--recurse-submodules-default=%s", value))
	}
}

// Refmap When fetching refs listed on the command line, use the specified refspec (can be given more than once) to map the refs to remote-tracking branches, instead of the values of remote.*.fetch configuration variables for the remote repository. 
// --refmap=<refspec>
func Refmap(refspec string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--refmap=%s", refspec))
	}
}

// ShallowExclude Deepen or shorten the history of a shallow repository to exclude commits reachable from a specified remote branch or tag. 
// This option can be specified multiple times. 
// --shallow-exclude=<revision>
func ShallowExclude(revision string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--shallow-exclude=%s", revision))
	}
}

// ShallowSince Deepen or shorten the history of a shallow repository to include all reachable commits after <date>. 
// --shallow-since=<date>
func ShallowSince(date string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--shallow-since=%s", date))
	}
}

// SubmodulePrefix Prepend <path> to paths printed in informative messages. 
// --submodule-prefix=<path>
func SubmodulePrefix(path string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--submodule-prefix=%s", path))
	}
}

// Tags Fetch all tags from the remote (i.e., fetch remote tags refs/tags/* into local tags with the same name), in addition to whatever else would otherwise be fetched. 
// -t, --tags
func Tags(g *types.Cmd) {
	g.AddOptions("--tags")
}

// Unshallow If the source repository is complete, convert a shallow repository to a complete one, removing all the limitations imposed by shallow repositories. 
// If the source repository is shallow, fetch as much as possible so that the current repository has the same history as the source repository. 
// --unshallow
func Unshallow(g *types.Cmd) {
	g.AddOptions("--unshallow")
}

// UpdateHeadOk By default git fetch refuses to update the head which corresponds to the current branch. 
// This flag disables the check. 
// This is purely for the internal use for git pull to communicate with git fetch, and unless you are implementing your own Porcelain you are not supposed to use it. 
// -u, --update-head-ok
func UpdateHeadOk(g *types.Cmd) {
	g.AddOptions("--update-head-ok")
}

// UpdateShallow By default when fetching from a shallow repository, git fetch refuses refs that require updating .git/shallow. 
// This option updates .git/shallow and accept such refs. 
// --update-shallow
func UpdateShallow(g *types.Cmd) {
	g.AddOptions("--update-shallow")
}

// UploadPack When given, and the repository to fetch from is handled by git fetch-pack, --exec=<upload-pack> is passed to the command to specify non-default path for the command run on the other end. 
// --upload-pack <upload-pack>
func UploadPack(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--upload-pack")
		g.AddOptions(value)
	}
}

// Verbose Be verbose. 
// -v, --verbose
func Verbose(g *types.Cmd) {
	g.AddOptions("--verbose")
}
