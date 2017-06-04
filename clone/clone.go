package clone

import (
	"fmt"
	"github.com/ldez/go-git-cmd-wrapper/git"
)

// Clone https://git-scm.com/docs/git-clone
// git clone [--template=<template_directory>]
//			[-l] [-s] [--no-hardlinks] [-q] [-n] [--bare] [--mirror]
//			[-o <name>] [-b <name>] [-u <upload-pack>] [--reference <repository>]
//			[--dissociate] [--separate-git-dir <git dir>]
//			[--depth <depth>] [--[no-]single-branch]
//			[--recurse-submodules] [--[no-]shallow-submodules]
//			[--jobs <n>] [--] <repository> [<directory>]

func Template(templateDirectory string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(fmt.Sprintf("--template=%s", templateDirectory))
	}
}

func Local(g *git.Cmd) {
	g.AddOptions("--local")
}

func Shared(g *git.Cmd) {
	g.AddOptions("--shared")
}

func NoHardlinks(g *git.Cmd) {
	g.AddOptions("--no-hardlinks")
}

func Quiet(g *git.Cmd) {
	g.AddOptions("--quiet")
}

func NoCheckout(g *git.Cmd) {
	g.AddOptions("--no-checkout")
}

func Mirror(g *git.Cmd) {
	g.AddOptions("--mirror")
}

func Bare(g *git.Cmd) {
	g.AddOptions("--bare")
}

func Origin(name string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions("--origin")
		g.AddOptions(name)
	}
}

func Branch(name string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions("--branch")
		g.AddOptions(name)
	}
}

func UploadPack(name string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions("--upload-pack")
		g.AddOptions(name)
	}
}

func Dissociate(g *git.Cmd) {
	g.AddOptions("--dissociate")
}

func SeparateGitDir(gitDir string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions("--separate-git-dir")
		g.AddOptions(gitDir)
	}
}

func Depth(depth string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions("--depth")
		g.AddOptions(depth)
	}
}

func SingleBranch(g *git.Cmd) {
	g.AddOptions("--single-branch")
}

func NoSingleBranch(g *git.Cmd) {
	g.AddOptions("--no-single-branch")
}

func RecurseSubmodules(g *git.Cmd) {
	g.AddOptions("--recurse-submodules")
}

func ShallowSubmodules(g *git.Cmd) {
	g.AddOptions("--shallow-submodules")
}

func NoShallowSubmodules(g *git.Cmd) {
	g.AddOptions("--no-shallow-submodules")
}

func Jobs(n string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions("--jobs")
		g.AddOptions(n)
	}
}

func Repository(url string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(url)
	}
}

func Directory(name string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(name)
	}
}
