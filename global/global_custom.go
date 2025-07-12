package global

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/v2/types"
)

// UpperC Run as if git was started in <path> instead of the current working directory.
// When multiple -C options are given, each subsequent non-absolute -C <path> is interpreted relative to the preceding -C <path>.
// If <path> is present but empty, e.g. -C "", then the current working directory is left unchanged.
//
// This option affects options that expect path name like --git-dir and --work-tree in that their interpretations of the path names would be made relative to the working directory caused by the -C option.
// -C path
func UpperC(path string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddBaseOptions("-C")
		g.AddBaseOptions(path)
	}
}

// LowerC Pass a configuration parameter to the command.
// The value given will override values from configuration files.
// The <name> is expected in the same format as listed by git config (subkeys separated by dots).
// -c foo.bar=value
func LowerC(name, value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddBaseOptions("-c")
		g.AddBaseOptions(fmt.Sprintf("%s=%s", name, value))
	}
}

// ConfigEnv Like -c <name>=<value>, give configuration variable <name> a value,
// where <envvar> is the name of an environment variable from which to retrieve the value.
// Unlike -c there is no shortcut for directly setting the value to an empty string,
// instead the environment variable itself must be set to the empty string.
// It is an error if the <envvar> does not exist in the environment.
// <envvar> may not contain an equals sign to avoid ambiguity with <name> containing one.
// --config-env=<name>=<envvar>
func ConfigEnv(name, envvar string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddBaseOptions(fmt.Sprintf("--config-env=%s=%s", name, envvar))
	}
}

// ExecPath Path to wherever your core Git programs are installed.
// This can also be controlled by setting the GIT_EXEC_PATH environment variable.
// If no path is given, git will print the current setting and then exit.
// --exec-path[=<path>]
func ExecPath(path string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if path == "" {
			g.AddBaseOptions("--exec-path")
			return
		}

		g.AddBaseOptions(fmt.Sprintf("--exec-path=%s", path))
	}
}

// HTMLPath Print the path, without trailing slash,
// where Git’s HTML documentation is installed and exit.
// --html-path
func HTMLPath() func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddBaseOptions("--html-path")
	}
}

// ManPath Print the manpath (see man(1)) for the man pages for this version of Git and exit.
// --man-path
func ManPath() func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddBaseOptions("--man-path")
	}
}

// InfoPath Print the path where the Info files documenting this version of Git are installed and exit.
// --info-path
func InfoPath() func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddBaseOptions("--info-path")
	}
}

// Paginate Pipe all output into less (or if set, $PAGER) if standard output is a terminal.
// This overrides the pager.<cmd> configuration options (see the "Configuration Mechanism" section below).
// -p
// --paginate
func Paginate() func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddBaseOptions("--paginate")
	}
}

// NoPager Do not pipe Git output into a pager.
// -P
// --no-pager
func NoPager() func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddBaseOptions("--no-pager")
	}
}

// GitDir Set the path to the repository (".git" directory).
// This can also be controlled by setting the GIT_DIR environment variable.
// It can be an absolute path or relative path to current working directory.
//
// Specifying the location of the ".git" directory using this option (or GIT_DIR environment variable)
// turns off the repository discovery that tries to find a directory with ".git" subdirectory
// (which is how the repository and the top-level of the working tree are discovered),
// and tells Git that you are at the top level of the working tree.
// If you are not at the top-level directory of the working tree,
// you should tell Git where the top-level of the working tree is,
// with the --work-tree=<path> option (or GIT_WORK_TREE environment variable).
//
// If you just want to run git as if it was started in <path> then use git -C <path>.
// --git-dir=<path>
func GitDir(path string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddBaseOptions(fmt.Sprintf("--git-dir=%s", path))
	}
}

// WorkTree Set the path to the working tree.
// It can be an absolute path or a path relative to the current working directory.
// This can also be controlled by setting the GIT_WORK_TREE environment variable
// and the core.worktree configuration variable
// (see core.worktree in git-config[1] for a more detailed discussion).
// --work-tree=<path>
func WorkTree(path string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddBaseOptions(fmt.Sprintf("--work-tree=%s", path))
	}
}

// Namespace Set the Git namespace.
// See gitnamespaces[7] for more details.
// Equivalent to setting the GIT_NAMESPACE environment variable.
// --namespace=<path>
func Namespace(path string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddBaseOptions(fmt.Sprintf("--namespace=%s", path))
	}
}

// Bare Treat the repository as a bare repository.
// If GIT_DIR environment is not set, it is set to the current working directory.
// --bare
func Bare(g *types.Cmd) {
	g.AddBaseOptions("--bare")
}

// NoReplaceObjects Do not use replacement refs to replace Git objects.
// This is equivalent to exporting the GIT_NO_REPLACE_OBJECTS environment variable with any value.
// See git-replace[1] for more information.
// --no-replace-objects
func NoReplaceObjects(g *types.Cmd) {
	g.AddBaseOptions("--no-replace-objects")
}

// NoLazyFetch Do not fetch missing objects from the promisor remote on demand.
// Useful together with git cat-file -e <object> to see if the object is locally available.
// This is equivalent to setting the GIT_NO_LAZY_FETCH environment variable to 1.
// --no-lazy-fetch
func NoLazyFetch(g *types.Cmd) {
	g.AddBaseOptions("--no-lazy-fetch")
}

// NoOptionalLocks Do not perform optional operations that require locks.
// This is equivalent to setting the GIT_OPTIONAL_LOCKS to 0.
// --no-optional-locks
func NoOptionalLocks(g *types.Cmd) {
	g.AddBaseOptions("--no-optional-locks")
}

// NoAdvice Disable all advice hints from being printed.
// --no-advice
func NoAdvice(g *types.Cmd) {
	g.AddBaseOptions("--no-advice")
}

// LiteralPathSpecs Treat pathspecs literally (i.e. no globbing, no pathspec magic).
// This is equivalent to setting the GIT_LITERAL_PATHSPECS environment variable to 1.
// --literal-pathspecs
func LiteralPathSpecs(g *types.Cmd) {
	g.AddBaseOptions("--literal-pathspecs")
}

// GlobPathSpecs Add "glob" magic to all pathspec.
// This is equivalent to setting the GIT_GLOB_PATHSPECS environment variable to 1.
// Disabling globbing on individual pathspecs can be done using pathspec magic ":(literal)"
// --glob-pathspecs
func GlobPathSpecs(g *types.Cmd) {
	g.AddBaseOptions("--glob-pathspecs")
}

// NoGlobPathSpecs Add "literal" magic to all pathspec.
// This is equivalent to setting the GIT_NOGLOB_PATHSPECS environment variable to 1.
// Enabling globbing on individual pathspecs can be done using pathspec magic ":(glob)"
// --noglob-pathspecs
func NoGlobPathSpecs(g *types.Cmd) {
	g.AddBaseOptions("--noglob-pathspecs")
}

// ICasePathSpecs Add "icase" magic to all pathspec.
// This is equivalent to setting the GIT_ICASE_PATHSPECS environment variable to 1.
// --icase-pathspecs
func ICasePathSpecs(g *types.Cmd) {
	g.AddBaseOptions("--icase-pathspecs")
}

// AttrSource Read gitattributes from <tree-ish> instead of the worktree.
// See gitattributes[5].
// This is equivalent to setting the GIT_ATTR_SOURCE environment variable.
// --attr-source=<tree-ish>
func AttrSource(treeIsh string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddBaseOptions(fmt.Sprintf("--attr-source=%s", treeIsh))
	}
}
