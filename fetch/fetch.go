package fetch

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/git"
)

// Fetch https://git-scm.com/docs/git-fetch
// git fetch [<options>] [<repository> [<refspec>…​]]
// git fetch [<options>] <group>
// git fetch --multiple [<options>] [(<repository> | <group>)…​]
// git fetch --all [<options>]
//--all
//--append
//--depth=<depth>
//--deepen=<depth>
//--shallow-since=<date>
//--shallow-exclude=<revision>
//--unshallow
//--update-shallow
//--dry-run
//--force
//--keep
//--multiple
//--prune
//--no-tags
//--refmap=<refspec>
//--tags
//--recurse-submodules[=yes|on-demand|no]
//--jobs=<n>
//--no-recurse-submodules
//--submodule-prefix=<path>
//--recurse-submodules-default=[yes|on-demand]
//--update-head-ok
//--upload-pack <upload-pack>
//--quiet
//--verbose
//--progress
//--ipv4
//--ipv6
//
//<repository>
//<group>
//<refspec>

func All(g *git.Cmd) {
	g.AddOptions("--all")
}

func Append(g *git.Cmd) {
	g.AddOptions("--append")
}

func Depth(depth string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(fmt.Sprintf("--depth=%s", depth))
	}
}

func Deepen(depth string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(fmt.Sprintf("--deepen=%s", depth))
	}
}

func ShallowSince(date string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(fmt.Sprintf("--shallow-since=%s", date))
	}
}

func ShallowExclude(revision string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(fmt.Sprintf("--shallow-exclude=%s", revision))
	}
}

func Unshallow(g *git.Cmd) {
	g.AddOptions("--unshallow")
}

func UpdateShallow(g *git.Cmd) {
	g.AddOptions("--update-shallow")
}

func DryRun(g *git.Cmd) {
	g.AddOptions("--dry-run")
}

func Force(g *git.Cmd) {
	g.AddOptions("--force")
}

func Keep(g *git.Cmd) {
	g.AddOptions("--keep")
}

func Multiple(g *git.Cmd) {
	g.AddOptions("--multiple")
}

func Prune(g *git.Cmd) {
	g.AddOptions("--prune")
}

func NoTags(g *git.Cmd) {
	g.AddOptions("--no-tags")
}

func Tags(g *git.Cmd) {
	g.AddOptions("--tags")
}

func RefMap(refspec string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(fmt.Sprintf("--refmap=%s", refspec))
	}
}

func RecurseSubmodules(option string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(fmt.Sprintf("--recurse-submodules=%s", option))
	}
}

func Jobs(n string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(fmt.Sprintf("--jobs=%s", n))
	}
}

func NoRecurseSubmodules(g *git.Cmd) {
	g.AddOptions("--no-recurse-submodules")
}

func SubmodulePrefix(path string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(fmt.Sprintf("--submodule-prefix=%s", path))
	}
}

func RecurseSubmodulesDefault(option string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(fmt.Sprintf("--recurse-submodules-default=%s", option))
	}
}

func updateHeadOk(g *git.Cmd) {
	g.AddOptions("--update-head-ok")
}

func UploadPack(pack string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions("--upload-pack")
		g.AddOptions(pack)
	}
}

func Quiet(g *git.Cmd) {
	g.AddOptions("--quiet")
}

func Verbose(g *git.Cmd) {
	g.AddOptions("--verbose")
}

func Progress(g *git.Cmd) {
	g.AddOptions("--progress")
}

func IPv4(g *git.Cmd) {
	g.AddOptions("--ipv4")
}

func IPv6(g *git.Cmd) {
	g.AddOptions("--ipv6")
}

func Remote(name string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(name)
	}
}

func Group(name string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(name)
	}
}

func RefSpec(ref string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(ref)
	}
}
