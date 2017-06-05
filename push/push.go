// Push https://git-scm.com/docs/git-push
// git push [--all | --mirror | --tags] [--follow-tags] [--atomic] [-n | --dry-run] [--receive-pack=<git-receive-pack>]
// 			[--repo=<repository>] [-f | --force] [-d | --delete] [--prune] [-v | --verbose]
// 			[-u | --set-upstream] [--push-option=<string>]
// 			[--[no-]signed|--sign=(true|false|if-asked)]
// 			[--force-with-lease[=<refname>[:<expect>]]]
// 			[--no-verify] [<repository> [<refspec>...]]
package push

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/types"
)

func All(g *types.Cmd) {
	g.AddOptions("--all")
}

func Mirror(g *types.Cmd) {
	g.AddOptions("--mirror")
}

func Tags(g *types.Cmd) {
	g.AddOptions("--tags")
}

func FollowTags(g *types.Cmd) {
	g.AddOptions("--follow-tags")
}

func Atomic(g *types.Cmd) {
	g.AddOptions("--atomic")
}

func DryRun(g *types.Cmd) {
	g.AddOptions("--dry-run")
}

func ReceivePack(gitReceivePack string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--receive-pack=%s", gitReceivePack))
	}
}

func Repo(repository string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--repo=%s", repository))
	}
}

func Delete(g *types.Cmd) {
	g.AddOptions("--delete")
}

func Prune(g *types.Cmd) {
	g.AddOptions("--prune")
}

func Verbose(g *types.Cmd) {
	g.AddOptions("--verbose")
}

func SetUpstream(g *types.Cmd) {
	g.AddOptions("--set-upstream")
}

func PushOption(option string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--push-option=%s", option))
	}
}

func Force(g *types.Cmd) {
	g.AddOptions("--force")
}

func ForceWithLease(g *types.Cmd) {
	g.AddOptions("--force-with-lease")
}

func Remote(repository string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(repository)
	}
}

func RefSpec(refspec string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(refspec)
	}
}
