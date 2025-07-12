package remote

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/v2/types"
)

// Add git remote add [-t <branch>] [-m <master>] [-f] [--[no-]tags] [--mirror=<fetch|push>] <name> <url>
func Add(name, url string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("add")
		g.AddOptions(name)
		g.AddOptions(url)
	}
}

// Rename git remote rename <old> <new>
func Rename(oldName, newName string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("rename")
		g.AddOptions(oldName)
		g.AddOptions(newName)
	}
}

// Remove git remote remove <name>
func Remove(name string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("remove")
		g.AddOptions(name)
	}
}

// SetHead git remote set-head <name> (-a | --auto | -d | --delete | <branch>)
func SetHead(name string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("set-head")
		g.AddOptions(name)
	}
}

// SetBranches git remote set-branches [--add] <name> <branch>…
func SetBranches(name, branch string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("set-branches")
		g.AddOptions(name)
		g.AddOptions(branch)
	}
}

// GetURL git remote get-url [--push] [--all] <name>
func GetURL(name string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("get-url")
		g.AddOptions(name)
	}
}

// SetURL git remote set-url [--push] <name> <newurl> [<oldurl>]
// git remote set-url --add [--push] <name> <newurl>
// git remote set-url --delete [--push] <name> <url>
func SetURL(name, newurl, oldurl string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("set-url")
		g.AddOptions(name)
		g.AddOptions(newurl)

		if oldurl != "" {
			g.AddOptions(oldurl)
		}
	}
}

// Show git remote [-v | --verbose] show [-n] <name>…
func Show(names ...string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("show")

		for _, name := range names {
			g.AddOptions(name)
		}
	}
}

// Prune git remote prune [-n | --dry-run] <name>…
func Prune(names ...string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("prune")

		for _, name := range names {
			g.AddOptions(name)
		}
	}
}

// Update Fetch updates for remotes or remote groups in the repository as defined by remotes.<group>.
func Update(remotes ...string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("update")

		for _, name := range remotes {
			g.AddOptions(name)
		}
	}
}

// Verbose [-v | --verbose]
func Verbose(g *types.Cmd) {
	g.AddOptions("--verbose")
}

// Tags [--tags]
func Tags(g *types.Cmd) {
	g.AddOptions("--tags")
}

// NoTags [--no-tags]
func NoTags(g *types.Cmd) {
	g.AddOptions("--no-tags")
}

// Fetch [-f]
func Fetch(g *types.Cmd) {
	g.AddOptions("-f")
}

// Auto -a | --auto
func Auto(g *types.Cmd) {
	g.AddOptions("--auto")
}

// Delete -d | --delete
func Delete(g *types.Cmd) {
	g.AddOptions("--delete")
}

// AddOpt [--add]
func AddOpt(g *types.Cmd) {
	g.AddOptions("--add")
}

// Push [--push]
func Push(g *types.Cmd) {
	g.AddOptions("--push")
}

// All [--all]
func All(g *types.Cmd) {
	g.AddOptions("--all")
}

// DryRun [-n | --dry-run]
func DryRun(g *types.Cmd) {
	g.AddOptions("--dry-run")
}

// PruneOpt [-p | --prune]
func PruneOpt(g *types.Cmd) {
	g.AddOptions("--prune")
}

// Master [-m <master>]
func Master(symRef string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("-m")
		g.AddOptions(symRef)
	}
}

// Track [-t <branch>]
func Track(branch string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("-t")
		g.AddOptions(branch)
	}
}

// Mirror [--mirror=<fetch|push>]
func Mirror(option string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--mirror=%s", option))
	}
}
