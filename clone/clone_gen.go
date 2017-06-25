/*
* CODE GENERATED AUTOMATICALLY
* THIS FILE MUST NOT BE EDITED BY HAND
 */
package clone

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/types"
)

// Bare Make a bare Git repository. That is, instead of creating <directory> and placing the administrative files in <directory>/.git, make the <directory> itself the $GIT_DIR. This obviously implies the -n because there is nowhere to check out the working tree. Also the branch heads at the remote are copied directly to corresponding local branch heads, without mapping them to refs/remotes/origin/. When this option is used, neither remote-tracking branches nor the related configuration variables are created. 
// --bare
func Bare(g *types.Cmd) {
	g.AddOptions("--bare")
}

// Branch Instead of pointing the newly created HEAD to the branch pointed to by the cloned repository’s HEAD, point to <name> branch instead. In a non-bare repository, this is the branch that will be checked out.  --branch can also take tags and detaches the HEAD at that commit in the resulting repository. 
// --branch <name>, -b <name>
func Branch(name string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--branch")
		g.AddOptions(name)
	}
}

// Depth Create a shallow clone with a history truncated to the specified number of commits. Implies --single-branch unless --no-single-branch is given to fetch the histories near the tips of all branches. If you want to clone submodules shallowly, also pass --shallow-submodules. 
// --depth <depth>
func Depth(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--depth")
		g.AddOptions(value)
	}
}

// Dissociate Borrow the objects from reference repositories specified with the --reference options only to reduce network transfer, and stop borrowing from them after a clone is made by making necessary local copies of borrowed objects. This option can also be used when cloning locally from a repository that already borrows objects from another repository—the new repository will borrow objects from the same repository, and this option can be used to stop the borrowing. 
// --dissociate
func Dissociate(g *types.Cmd) {
	g.AddOptions("--dissociate")
}

// Jobs The number of submodules fetched at the same time. Defaults to the submodule.fetchJobs option. 
// -j <n>, --jobs <n>
func Jobs(n string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--jobs")
		g.AddOptions(n)
	}
}

// Local When the repository to clone from is on a local machine, this flag bypasses the normal 'Git aware' transport mechanism and clones the repository by making a copy of HEAD and everything under objects and refs directories. The files under .git/objects/ directory are hardlinked to save space when possible. 
// --local, -l
func Local(g *types.Cmd) {
	g.AddOptions("--local")
}

// Mirror Set up a mirror of the source repository. This implies --bare. Compared to --bare, --mirror not only maps local branches of the source to local branches of the target, it maps all refs (including remote-tracking branches, notes etc.) and sets up a refspec configuration such that all these refs are overwritten by a git remote update in the target repository. 
// --mirror
func Mirror(g *types.Cmd) {
	g.AddOptions("--mirror")
}

// NoCheckout No checkout of HEAD is performed after the clone is complete. 
// --no-checkout, -n
func NoCheckout(g *types.Cmd) {
	g.AddOptions("--no-checkout")
}

// NoHardlinks Force the cloning process from a repository on a local filesystem to copy the files under the .git/objects directory instead of using hardlinks. This may be desirable if you are trying to make a back-up of your repository. 
// --no-hardlinks
func NoHardlinks(g *types.Cmd) {
	g.AddOptions("--no-hardlinks")
}

// NoShallowSubmodules All submodules which are cloned will be shallow with a depth of 1. 
// --no-shallow-submodules
func NoShallowSubmodules(g *types.Cmd) {
	g.AddOptions("--no-shallow-submodules")
}

// NoSingleBranch Clone only the history leading to the tip of a single branch, either specified by the --branch option or the primary branch remote’s HEAD points at. Further fetches into the resulting repository will only update the remote-tracking branch for the branch this option was used for the initial cloning. If the HEAD at the remote did not point at any branch when --single-branch clone was made, no remote-tracking branch is created. 
// --no-single-branch
func NoSingleBranch(g *types.Cmd) {
	g.AddOptions("--no-single-branch")
}

// Origin Instead of using the remote name origin to keep track of the upstream repository, use <name>. 
// --origin <name>, -o <name>
func Origin(name string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--origin")
		g.AddOptions(name)
	}
}

// Progress Progress status is reported on the standard error stream by default when it is attached to a terminal, unless -q is specified. This flag forces progress status even if the standard error stream is not directed to a terminal. 
// --progress
func Progress(g *types.Cmd) {
	g.AddOptions("--progress")
}

// Quiet Operate quietly. Progress is not reported to the standard error stream. 
// --quiet, -q
func Quiet(g *types.Cmd) {
	g.AddOptions("--quiet")
}

// RecurseSubmodules After the clone is created, initialize and clone submodules within based on the provided pathspec. If no pathspec is provided, all submodules are initialized and cloned. Submodules are initialized and cloned using their default settings. The resulting clone has submodule.active set to the provided pathspec, or '.' (meaning all submodules) if no pathspec is provided. This is equivalent to running git submodule update --init --recursive immediately after the clone is finished. This option is ignored if the cloned repository does not have a worktree/checkout (i.e. if any of --no-checkout/-n, --bare, or --mirror is given) 
// --recurse-submodules[=<pathspec>]
func RecurseSubmodules(pathspec string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(pathspec) == 0 {
			g.AddOptions("--recurse-submodules")
		} else {
			g.AddOptions(fmt.Sprintf("--recurse-submodules=%s", pathspec))
		}
	}
}

// SeparateGitDir Instead of placing the cloned repository where it is supposed to be, place the cloned repository at the specified directory, then make a filesystem-agnostic Git symbolic link to there. The result is Git repository can be separated from working tree. 
// --separate-git-dir=<git dir>
func SeparateGitDir(gitDir string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--separate-git-dir=%s", gitDir))
	}
}

// ShallowExclude Create a shallow clone with a history, excluding commits reachable from a specified remote branch or tag. This option can be specified multiple times. 
// --shallow-exclude=<revision>
func ShallowExclude(revision string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--shallow-exclude=%s", revision))
	}
}

// ShallowSince Create a shallow clone with a history after the specified time. 
// --shallow-since=<date>
func ShallowSince(date string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--shallow-since=%s", date))
	}
}

// ShallowSubmodules All submodules which are cloned will be shallow with a depth of 1. 
// --shallow-submodules
func ShallowSubmodules(g *types.Cmd) {
	g.AddOptions("--shallow-submodules")
}

// Shared When the repository to clone is on the local machine, instead of using hard links, automatically setup .git/objects/info/alternates to share the objects with the source repository. The resulting repository starts out without any object of its own. 
// --shared, -s
func Shared(g *types.Cmd) {
	g.AddOptions("--shared")
}

// SingleBranch Clone only the history leading to the tip of a single branch, either specified by the --branch option or the primary branch remote’s HEAD points at. Further fetches into the resulting repository will only update the remote-tracking branch for the branch this option was used for the initial cloning. If the HEAD at the remote did not point at any branch when --single-branch clone was made, no remote-tracking branch is created. 
// --single-branch
func SingleBranch(g *types.Cmd) {
	g.AddOptions("--single-branch")
}

// Template Specify the directory from which templates will be used; (See the 'TEMPLATE DIRECTORY' section of git-init(1).) 
// --template=<template_directory>
func Template(template_directory string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--template=%s", template_directory))
	}
}

// UploadPack When given, and the repository to clone from is accessed via ssh, this specifies a non-default path for the command run on the other end. 
// --upload-pack <upload-pack>, -u <upload-pack>
func UploadPack(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--upload-pack")
		g.AddOptions(value)
	}
}

// Verbose Run verbosely. Does not affect the reporting of progress status to the standard error stream. 
// --verbose, -v
func Verbose(g *types.Cmd) {
	g.AddOptions("--verbose")
}
