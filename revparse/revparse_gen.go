/*
Package revparse CODE GENERATED AUTOMATICALLY
THIS FILE MUST NOT BE EDITED BY HAND
*/
package revparse

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/types"
)

// AbbrevRef A non-ambiguous short name of the objects name. 
// The option core.warnAmbiguousRefs is used to select the strict abbreviation mode. 
// --abbrev-ref[=(strict|loose)]
func AbbrevRef(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(value) == 0 {
			g.AddOptions("--abbrev-ref")
		} else {
			g.AddOptions(fmt.Sprintf("--abbrev-ref=%s", value))
		}
	}
}

// AbsoluteGitDir Like --git-dir, but its output is always the canonicalized absolute path. 
// --absolute-git-dir
func AbsoluteGitDir(g *types.Cmd) {
	g.AddOptions("--absolute-git-dir")
}

// After Parse the date string, and output the corresponding --max-age= parameter for git rev-list. 
// --since=datestring, --after=datestring
func After(datestring string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("-after=%s", datestring))
	}
}

// All Show all refs found in refs/. 
// --all
func All(g *types.Cmd) {
	g.AddOptions("--all")
}

// Before Parse the date string, and output the corresponding --min-age= parameter for git rev-list. 
// --until=datestring, --before=datestring
func Before(datestring string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--before=%s", datestring))
	}
}

// Branches Show all branches, tags, or remote-tracking branches, respectively (i.e., refs found in refs/heads, refs/tags, or refs/remotes, respectively). 
// If a pattern is given, only refs matching the given shell glob are shown. 
// If the pattern does not contain a globbing character (?, *, or [), it is turned into a prefix match by appending /*. 
// --branches[=pattern], --tags[=pattern], --remotes[=pattern]
func Branches(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(value) == 0 {
			g.AddOptions("--branches")
		} else {
			g.AddOptions(fmt.Sprintf("--branches=%s", value))
		}
	}
}

// Default If there is no parameter given by the user, use <arg> instead. 
// --default <arg>
func Default(arg string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--default")
		g.AddOptions(arg)
	}
}

// Disambiguate Show every object whose name begins with the given prefix. 
// The <prefix> must be at least 4 hexadecimal digits long to avoid listing each and every object in the repository by mistake. 
// --disambiguate=<prefix>
func Disambiguate(prefix string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--disambiguate=%s", prefix))
	}
}

// Exclude Do not include refs matching <glob-pattern> that the next --all, --branches, --tags, --remotes, or --glob would otherwise consider. 
// Repetitions of this option accumulate exclusion patterns up to the next --all, --branches, --tags, --remotes, or --glob option (other options or arguments do not clear accumulated patterns). 
// The patterns given should not begin with refs/heads, refs/tags, or refs/remotes when applied to --branches, --tags, or --remotes, respectively, and they must begin with refs/ when applied to --glob or --all. 
// If a trailing /* is intended, it must be given explicitly. 
// --exclude=<glob-pattern>
func Exclude(globPattern string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--exclude=%s", globPattern))
	}
}

// Flags Do not output non-flag parameters. 
// --flags
func Flags(g *types.Cmd) {
	g.AddOptions("--flags")
}

// GitCommonDir Show $GIT_COMMON_DIR if defined, else $GIT_DIR. 
// --git-common-dir
func GitCommonDir(g *types.Cmd) {
	g.AddOptions("--git-common-dir")
}

// GitDir Show $GIT_DIR if defined. 
// Otherwise show the path to the .git directory. 
// The path shown, when relative, is relative to the current working directory. 
// If $GIT_DIR is not defined and the current directory is not detected to lie in a Git repository or work tree print a message to stderr and exit with nonzero status. 
// --git-dir
func GitDir(g *types.Cmd) {
	g.AddOptions("--git-dir")
}

// GitPath Resolve '$GIT_DIR/<path>' and takes other path relocation variables such as $GIT_OBJECT_DIRECTORY, $GIT_INDEX_FILE... into account. 
// For example, if $GIT_OBJECT_DIRECTORY is set to /foo/bar then 'git rev-parse --git-path objects/abc' returns /foo/bar/abc. 
// --git-path <path>
func GitPath(path string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--git-path")
		g.AddOptions(path)
	}
}

// Glob Show all refs matching the shell glob pattern pattern. 
// If the pattern does not start with refs/, this is automatically prepended. 
// If the pattern does not contain a globbing character (?, *, or [), it is turned into a prefix match by appending /*. 
// --glob=pattern
func Glob(pattern string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--glob=%s", pattern))
	}
}

// IsBareRepository When the repository is bare print 'true', otherwise 'false'. 
// --is-bare-repository
func IsBareRepository(g *types.Cmd) {
	g.AddOptions("--is-bare-repository")
}

// IsInsideGitDir When the current working directory is below the repository directory print 'true', otherwise 'false'. 
// --is-inside-git-dir
func IsInsideGitDir(g *types.Cmd) {
	g.AddOptions("--is-inside-git-dir")
}

// IsInsideWorkTree When the current working directory is inside the work tree of the repository print 'true', otherwise 'false'. 
// --is-inside-work-tree
func IsInsideWorkTree(g *types.Cmd) {
	g.AddOptions("--is-inside-work-tree")
}

// KeepDashdash Only meaningful in --parseopt mode. 
// Tells the option parser to echo out the first -- met instead of skipping it. 
// --keep-dashdash
func KeepDashdash(g *types.Cmd) {
	g.AddOptions("--keep-dashdash")
}

// LocalEnvVars List the GIT_* environment variables that are local to the repository (e.g. GIT_DIR or GIT_WORK_TREE, but not GIT_EDITOR). 
// Only the names of the variables are listed, not their value, even if they are set. 
// --local-env-vars
func LocalEnvVars(g *types.Cmd) {
	g.AddOptions("--local-env-vars")
}

// NoFlags Do not output flag parameters. 
// --no-flags
func NoFlags(g *types.Cmd) {
	g.AddOptions("--no-flags")
}

// NoRevs Do not output flags and parameters meant for git rev-list command. 
// --no-revs
func NoRevs(g *types.Cmd) {
	g.AddOptions("--no-revs")
}

// Not When showing object names, prefix them with ^ and strip ^ prefix from the object names that already have one. 
// --not
func Not(g *types.Cmd) {
	g.AddOptions("--not")
}

// Parseopt Use git rev-parse in option parsing mode (see PARSEOPT section below). 
// --parseopt
func Parseopt(g *types.Cmd) {
	g.AddOptions("--parseopt")
}

// Prefix Behave as if git rev-parse was invoked from the <arg> subdirectory of the working tree. 
// Any relative filenames are resolved as if they are prefixed by <arg> and will be printed in that form. 
// This can be used to convert arguments to a command run in a subdirectory so that they can still be used after moving to the top-level of the repository. 
// --prefix <arg>
func Prefix(arg string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--prefix")
		g.AddOptions(arg)
	}
}

// Quiet Only meaningful in --verify mode. 
// Do not output an error message if the first argument is not a valid object name; instead exit with non-zero status silently. 
// SHA-1s for valid object names are printed to stdout on success. 
// -q, --quiet
func Quiet(g *types.Cmd) {
	g.AddOptions("--quiet")
}

// Remotes Show all branches, tags, or remote-tracking branches, respectively (i.e., refs found in refs/heads, refs/tags, or refs/remotes, respectively). 
// If a pattern is given, only refs matching the given shell glob are shown. 
// If the pattern does not contain a globbing character (?, *, or [), it is turned into a prefix match by appending /*. 
// --branches[=pattern], --tags[=pattern], --remotes[=pattern]
func Remotes(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(value) == 0 {
			g.AddOptions("--remotes")
		} else {
			g.AddOptions(fmt.Sprintf("--remotes=%s", value))
		}
	}
}

// ResolveGitDir Check if <path> is a valid repository or a gitfile that points at a valid repository, and print the location of the repository. 
// If <path> is a gitfile then the resolved path to the real repository is printed. 
// --resolve-git-dir <path>
func ResolveGitDir(path string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--resolve-git-dir")
		g.AddOptions(path)
	}
}

// RevsOnly Do not output flags and parameters not meant for git rev-list command. 
// --revs-only
func RevsOnly(g *types.Cmd) {
	g.AddOptions("--revs-only")
}

// SharedIndexPath Show the path to the shared index file in split index mode, or empty if not in split-index mode. 
// --shared-index-path
func SharedIndexPath(g *types.Cmd) {
	g.AddOptions("--shared-index-path")
}

// Short Same as --verify but shortens the object name to a unique prefix with at least length characters. 
// The minimum length is 4, the default is the effective value of the core.abbrev configuration variable (see git-config(1)). 
// --short[=length]
func Short(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(value) == 0 {
			g.AddOptions("--short")
		} else {
			g.AddOptions(fmt.Sprintf("--short=%s", value))
		}
	}
}

// ShowCdup When the command is invoked from a subdirectory, show the path of the top-level directory relative to the current directory (typically a sequence of '../', or an empty string). 
// --show-cdup
func ShowCdup(g *types.Cmd) {
	g.AddOptions("--show-cdup")
}

// ShowPrefix When the command is invoked from a subdirectory, show the path of the current directory relative to the top-level directory. 
// --show-prefix
func ShowPrefix(g *types.Cmd) {
	g.AddOptions("--show-prefix")
}

// ShowSuperprojectWorkingTree Show the absolute path of the root of the superproject’s working tree (if exists) that uses the current repository as its submodule. 
// Outputs nothing if the current repository is not used as a submodule by any project. 
// --show-superproject-working-tree
func ShowSuperprojectWorkingTree(g *types.Cmd) {
	g.AddOptions("--show-superproject-working-tree")
}

// ShowToplevel Show the absolute path of the top-level directory. 
// --show-toplevel
func ShowToplevel(g *types.Cmd) {
	g.AddOptions("--show-toplevel")
}

// Since Parse the date string, and output the corresponding --max-age= parameter for git rev-list. 
// --since=datestring, --after=datestring
func Since(datestring string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--since=%s", datestring))
	}
}

// Sq Usually the output is made one line per flag and parameter. 
// This option makes output a single line, properly quoted for consumption by shell. 
// Useful when you expect your parameter to contain whitespaces and newlines (e.g. when using pickaxe -S with git diff-*). 
// In contrast to the --sq-quote option, the command input is still interpreted as usual. 
// --sq
func Sq(g *types.Cmd) {
	g.AddOptions("--sq")
}

// SqQuote Use git rev-parse in shell quoting mode (see SQ-QUOTE section below). 
// In contrast to the --sq option below, this mode does only quoting. 
// Nothing else is done to command input. 
// --sq-quote
func SqQuote(g *types.Cmd) {
	g.AddOptions("--sq-quote")
}

// StopAtNonOption Only meaningful in --parseopt mode. 
// Lets the option parser stop at the first non-option argument. 
// This can be used to parse sub-commands that take options themselves. 
// --stop-at-non-option
func StopAtNonOption(g *types.Cmd) {
	g.AddOptions("--stop-at-non-option")
}

// StuckLong Only meaningful in --parseopt mode. 
// Output the options in their long form if available, and with their arguments stuck. 
// --stuck-long
func StuckLong(g *types.Cmd) {
	g.AddOptions("--stuck-long")
}

// Symbolic Usually the object names are output in SHA-1 form (with possible ^ prefix); this option makes them output in a form as close to the original input as possible. 
// --symbolic
func Symbolic(g *types.Cmd) {
	g.AddOptions("--symbolic")
}

// SymbolicFullName This is similar to --symbolic, but it omits input that are not refs (i.e. branch or tag names; or more explicitly disambiguating 'heads/master' form, when you want to name the 'master' branch when there is an unfortunately named tag 'master'), and show them as full refnames (e.g. 'refs/heads/master'). 
// --symbolic-full-name
func SymbolicFullName(g *types.Cmd) {
	g.AddOptions("--symbolic-full-name")
}

// Tags Show all branches, tags, or remote-tracking branches, respectively (i.e., refs found in refs/heads, refs/tags, or refs/remotes, respectively). 
// If a pattern is given, only refs matching the given shell glob are shown. 
// If the pattern does not contain a globbing character (?, *, or [), it is turned into a prefix match by appending /*. 
// --branches[=pattern], --tags[=pattern], --remotes[=pattern]
func Tags(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(value) == 0 {
			g.AddOptions("--tags")
		} else {
			g.AddOptions(fmt.Sprintf("--tags=%s", value))
		}
	}
}

// Until Parse the date string, and output the corresponding --min-age= parameter for git rev-list. 
// --until=datestring, --before=datestring
func Until(datestring string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--until=%s", datestring))
	}
}

// Verify Verify that exactly one parameter is provided, and that it can be turned into a raw 20-byte SHA-1 that can be used to access the object database. 
// If so, emit it to the standard output; otherwise, error out. 
// --verify
func Verify(g *types.Cmd) {
	g.AddOptions("--verify")
}
