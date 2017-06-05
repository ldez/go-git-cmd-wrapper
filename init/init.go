// Init https://git-scm.com/docs/git-init
// git init [-q | --quiet] [--bare] [--template=<template_directory>]
// 			[--separate-git-dir <git dir>]
// 			[--shared[=<permissions>]] [directory]
package init

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/types"
)

func Quiet(g *types.Cmd) {
	g.AddOptions("--quiet")
}

func Bare(g *types.Cmd) {
	g.AddOptions("--bare")
}

func Template(templateDirectory string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--template=%s", templateDirectory))
	}
}

func Shared(g *types.Cmd) {
	g.AddOptions("--shared")
}

func SeparateGitDir(gitDir string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--separate-git-dir")
		g.AddOptions(gitDir)
	}
}

func SharedWithPerms(permissions string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--shared=%s", permissions))
	}
}

func Directory(directory string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(directory)
	}
}
