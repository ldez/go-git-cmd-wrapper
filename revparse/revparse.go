// git rev-parse [ --option ] <args>...
package revparse

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/types"
)

// Use git rev-parse in option parsing mode (see PARSEOPT section below).
// --parseopt
func Parseopt(g *types.Cmd) {
	g.AddOptions("--parseopt")
}

// Use git rev-parse in shell quoting mode (see SQ-QUOTE section below).
// In contrast to the --sq option below, this mode does only quoting.
// Nothing else is done to command input.
// --sq-quote
func SqQuote(g *types.Cmd) {
	g.AddOptions("--sq-quote")
}

func KeepDashdash(g *types.Cmd) {
	g.AddOptions("--keep-dashdash")
}

func StopAtNonOption(g *types.Cmd) {
	g.AddOptions("--stop-at-non-option")
}

func StuckLong(g *types.Cmd) {
	g.AddOptions("--stuck-long")
}

func RevsOnly(g *types.Cmd) {
	g.AddOptions("--revs-only")
}

func NoRevs(g *types.Cmd) {
	g.AddOptions("--no-revs")
}

func flags(g *types.Cmd) {
	g.AddOptions("--flags")
}

func NoFlags(g *types.Cmd) {
	g.AddOptions("--no-flags")
}

func verify(g *types.Cmd) {
	g.AddOptions("--verify")
}

func quiet(g *types.Cmd) {
	g.AddOptions("--quiet")
}

func sq(g *types.Cmd) {
	g.AddOptions("--sq")
}

func not(g *types.Cmd) {
	g.AddOptions("--not")
}

func symbolic(g *types.Cmd) {
	g.AddOptions("--symbolic")
}

func SymbolicFullName(g *types.Cmd) {
	g.AddOptions("--symbolic-full-name")
}

func All(g *types.Cmd) {
	g.AddOptions("--all")
}

func LocalEnvVars(g *types.Cmd) {
	g.AddOptions("--local-env-vars")
}

func GitDir(g *types.Cmd) {
	g.AddOptions("--git-dir")
}

func AbsoluteGitDir(g *types.Cmd) {
	g.AddOptions("--absolute-git-dir")
}

func GitCommonDir(g *types.Cmd) {
	g.AddOptions("--git-common-dir")
}

func IsInsideGitDir(g *types.Cmd) {
	g.AddOptions("--is-inside-git-dir")
}

func IsInsideWorkTree(g *types.Cmd) {
	g.AddOptions("--is-inside-work-tree")
}

func IsBareRepository(g *types.Cmd) {
	g.AddOptions("--is-bare-repository")
}

func ShowCdup(g *types.Cmd) {
	g.AddOptions("--show-cdup")
}

func ShowPrefix(g *types.Cmd) {
	g.AddOptions("--show-prefix")
}

func ShowToplevel(g *types.Cmd) {
	g.AddOptions("--show-toplevel")
}

func ShowSuperprojectWorkingTree(g *types.Cmd) {
	g.AddOptions("--show-superproject-working-tree")
}

func SharedIndexPath(g *types.Cmd) {
	g.AddOptions("--shared-index-path")
}

func Exclude(globPattern string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--exclude=%s", globPattern))
	}
}

func Disambiguate(prefix string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--disambiguate=%s", prefix))
	}
}

func Glob(pattern string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--glob=%s", pattern))
	}
}

func Since(date string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--since=%s", date))
	}
}

func After(date string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--after=%s", date))
	}
}

func Until(date string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--until=%s", date))
	}
}

func Before(date string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--before=%s", date))
	}
}

// --abbrev-ref[=(strict|loose)]
func AbbrevRef(mode string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(mode) == 0 {
			g.AddOptions("--abbrev-ref")
		} else {
			g.AddOptions(fmt.Sprintf("--abbrev-ref=%s", mode))
		}
	}
}

// --branches[=pattern]
func Branches(pattern string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(pattern) == 0 {
			g.AddOptions("--branches")
		} else {
			g.AddOptions(fmt.Sprintf("--branches=%s", pattern))
		}
	}
}

// --tags[=pattern]
func Tags(pattern string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(pattern) == 0 {
			g.AddOptions("--tags")
		} else {
			g.AddOptions(fmt.Sprintf("--tags=%s", pattern))
		}
	}
}

// --remotes[=pattern]
func Remotes(pattern string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(pattern) == 0 {
			g.AddOptions("--remotes")
		} else {
			g.AddOptions(fmt.Sprintf("--remotes=%s", pattern))
		}
	}
}

// --short, --short=number
func Short(number string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(number) == 0 {
			g.AddOptions("--short")
		} else {
			g.AddOptions(fmt.Sprintf("--short=%s", number))
		}
	}
}

func Args(args ...string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		for _, arg := range args {
			g.AddOptions(arg)
		}
	}
}

//--default <arg>
//--prefix <arg>
//--resolve-git-dir <path>
//--git-path <path>
