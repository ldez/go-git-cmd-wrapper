package init

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/git"
)

// Init https://git-scm.com/docs/git-init
// git init [-q | --quiet] [--bare] [--template=<template_directory>]
// 			[--separate-git-dir <git dir>]
// 			[--shared[=<permissions>]] [directory]

func Quiet(g *git.Cmd) {
	g.AddOptions("--quiet")
}

func Bare(g *git.Cmd) {
	g.AddOptions("--bare")
}

func Template(templateDirectory string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(fmt.Sprintf("--template=%s", templateDirectory))
	}
}

func Shared(g *git.Cmd) {
	g.AddOptions("--shared")
}

func SeparateGitDir(gitDir string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions("--separate-git-dir")
		g.AddOptions(gitDir)
	}
}

func SharedWithPerms(permissions string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(fmt.Sprintf("--shared=%s", permissions))
	}
}

func Directory(directory string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(directory)
	}
}
