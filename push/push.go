package push

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/git"
)

// Push https://git-scm.com/docs/git-push
// git push [--all | --mirror | --tags] [--follow-tags] [--atomic] [-n | --dry-run] [--receive-pack=<git-receive-pack>]
// 			[--repo=<repository>] [-f | --force] [-d | --delete] [--prune] [-v | --verbose]
// 			[-u | --set-upstream] [--push-option=<string>]
// 			[--[no-]signed|--sign=(true|false|if-asked)]
// 			[--force-with-lease[=<refname>[:<expect>]]]
// 			[--no-verify] [<repository> [<refspec>...]]

func All(g *git.Cmd) {
	g.AddOptions("--all")
}

func Mirror(g *git.Cmd) {
	g.AddOptions("--mirror")
}

func Tags(g *git.Cmd) {
	g.AddOptions("--tags")
}

func FollowTags(g *git.Cmd) {
	g.AddOptions("--follow-tags")
}

func Atomic(g *git.Cmd) {
	g.AddOptions("--atomic")
}

func DryRun(g *git.Cmd) {
	g.AddOptions("--dry-run")
}

func ReceivePack(gitReceivePack string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(fmt.Sprintf("--receive-pack=%s", gitReceivePack))
	}
}

func Repo(repository string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(fmt.Sprintf("--repo=%s", repository))
	}
}

func Delete(g *git.Cmd) {
	g.AddOptions("--delete")
}

func Prune(g *git.Cmd) {
	g.AddOptions("--prune")
}

func Verbose(g *git.Cmd) {
	g.AddOptions("--verbose")
}

func SetUpstream(g *git.Cmd) {
	g.AddOptions("--set-upstream")
}

func PushOption(option string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(fmt.Sprintf("--push-option=%s", option))
	}
}

func Force(g *git.Cmd) {
	g.AddOptions("--force")
}

func ForceWithLease(g *git.Cmd) {
	g.AddOptions("--force-with-lease")
}

func Remote(repository string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(repository)
	}
}

func RefSpec(refspec string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(refspec)
	}
}
