package remote

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/git"
)

// Remote https://git-scm.com/docs/git-remote
//	git remote [-v | --verbose]
//	git remote add [-t <branch>] [-m <master>] [-f] [--[no-]tags] [--mirror=<fetch|push>] <name> <url>
//	git remote rename <old> <new>
//	git remote remove <name>
//	git remote set-head <name> (-a | --auto | -d | --delete | <branch>)
//	git remote set-branches [--add] <name> <branch>…​
//	git remote get-url [--push] [--all] <name>
//	git remote set-url [--push] <name> <newurl> [<oldurl>]
//	git remote set-url --add [--push] <name> <newurl>
//	git remote set-url --delete [--push] <name> <url>
//	git remote [-v | --verbose] show [-n] <name>…​
//	git remote prune [-n | --dry-run] <name>…​
//	git remote [-v | --verbose] update [-p | --prune] [(<group> | <remote>)…​]

func Verbose(g *git.Cmd) {
	g.AddOptions("--verbose")
}

// ----------

func Add(g *git.Cmd) {
	g.AddOptions("add")
}

func Track(branch string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions("-t")
		g.AddOptions(branch)
	}
}

func Master(symRef string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions("-m")
		g.AddOptions(symRef)
	}
}

func Fetch(g *git.Cmd) {
	g.AddOptions("-f")
}

func Tags(g *git.Cmd) {
	g.AddOptions("--tags")
}

func NoTags(g *git.Cmd) {
	g.AddOptions("--no-tags")
}

func Mirror(option string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(fmt.Sprintf("--mirror=%s", option))
	}
}

func Name(name string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(name)
	}
}

func URL(url string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(url)
	}
}
