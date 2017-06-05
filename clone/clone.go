// Clone https://git-scm.com/docs/git-clone
// git clone [--template=<template_directory>]
//			[-l] [-s] [--no-hardlinks] [-q] [-n] [--bare] [--mirror]
//			[-o <name>] [-b <name>] [-u <upload-pack>] [--reference <repository>]
//			[--dissociate] [--separate-git-dir <git dir>]
//			[--depth <depth>] [--[no-]single-branch]
//			[--recurse-submodules] [--[no-]shallow-submodules]
//			[--jobs <n>] [--] <repository> [<directory>]
package clone

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/types"
)

func Template(templateDirectory string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--template=%s", templateDirectory))
	}
}

func Local(g *types.Cmd) {
	g.AddOptions("--local")
}

func Shared(g *types.Cmd) {
	g.AddOptions("--shared")
}

func NoHardlinks(g *types.Cmd) {
	g.AddOptions("--no-hardlinks")
}

func Quiet(g *types.Cmd) {
	g.AddOptions("--quiet")
}

func NoCheckout(g *types.Cmd) {
	g.AddOptions("--no-checkout")
}

func Mirror(g *types.Cmd) {
	g.AddOptions("--mirror")
}

func Bare(g *types.Cmd) {
	g.AddOptions("--bare")
}

func Origin(name string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--origin")
		g.AddOptions(name)
	}
}

func Branch(name string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--branch")
		g.AddOptions(name)
	}
}

func UploadPack(name string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--upload-pack")
		g.AddOptions(name)
	}
}

func Dissociate(g *types.Cmd) {
	g.AddOptions("--dissociate")
}

func SeparateGitDir(gitDir string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--separate-git-dir")
		g.AddOptions(gitDir)
	}
}

func Depth(depth string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--depth")
		g.AddOptions(depth)
	}
}

func SingleBranch(g *types.Cmd) {
	g.AddOptions("--single-branch")
}

func NoSingleBranch(g *types.Cmd) {
	g.AddOptions("--no-single-branch")
}

func RecurseSubmodules(g *types.Cmd) {
	g.AddOptions("--recurse-submodules")
}

func ShallowSubmodules(g *types.Cmd) {
	g.AddOptions("--shallow-submodules")
}

func NoShallowSubmodules(g *types.Cmd) {
	g.AddOptions("--no-shallow-submodules")
}

func Jobs(n string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--jobs")
		g.AddOptions(n)
	}
}

func Repository(url string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(url)
	}
}

func Directory(name string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(name)
	}
}
