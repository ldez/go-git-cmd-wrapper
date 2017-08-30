/*
Package init CODE GENERATED AUTOMATICALLY
THIS FILE MUST NOT BE EDITED BY HAND
*/
package init

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/types"
)

// Bare Create a bare repository. 
// If GIT_DIR environment is not set, it is set to the current working directory. 
// --bare
func Bare(g *types.Cmd) {
	g.AddOptions("--bare")
}

// Quiet Only print error and warning messages; all other output will be suppressed. 
// -q, --quiet
func Quiet(g *types.Cmd) {
	g.AddOptions("--quiet")
}

// SeparateGitDir Instead of initializing the repository as a directory to either $GIT_DIR or ./.git/, create a text file there containing the path to the actual repository. 
// This file acts as filesystem-agnostic Git symbolic link to the repository. 
// --separate-git-dir=<git dir>
func SeparateGitDir(gitDir string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--separate-git-dir=%s", gitDir))
	}
}

// Shared Specify that the Git repository is to be shared amongst several users. 
// This allows users belonging to the same group to push into that repository. 
// When specified, the config variable 'core.sharedRepository' is set so that files and directories under $GIT_DIR are created with the requested permissions. 
// When not specified, Git will use permissions reported by umask(2). 
// --shared[=(false|true|umask|group|all|world|everybody|0xxx)]
func Shared(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(value) == 0 {
			g.AddOptions("--shared")
		} else {
			g.AddOptions(fmt.Sprintf("--shared=%s", value))
		}
	}
}

// Template Specify the directory from which templates will be used. 
// (See the 'TEMPLATE DIRECTORY' section below.) 
// --template=<template_directory>
func Template(templateDirectory string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--template=%s", templateDirectory))
	}
}
