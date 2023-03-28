package lsfiles

// CODE GENERATED AUTOMATICALLY
// THIS FILE MUST NOT BE EDITED BY HAND

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/v2/types"
)

// Abbrev Instead of showing the full 40-byte hexadecimal object lines, show the shortest prefix that is at least <n> hexdigits long that uniquely refers the object. Non default number of digits can be specified with --abbrev=<n>.
// --abbrev[=<n>]
func Abbrev(n string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(n) == 0 {
			g.AddOptions("--abbrev")
		} else {
			g.AddOptions(fmt.Sprintf("--abbrev=%s", n))
		}
	}
}

// Cached Show cached files in the output (default)
// -c, --cached
func Cached(g *types.Cmd) {
	g.AddOptions("--cached")
}

// Debug After each line that describes a file, add more data about its cache entry. This is intended to show as much information as possible for manual inspection; the exact format may change at any time.
// --debug
func Debug(g *types.Cmd) {
	g.AddOptions("--debug")
}

// Deduplicate When only filenames are shown, suppress duplicates that may come from having multiple stages during a merge, or giving --deleted and --modified option at the same time. When any of the -t, --unmerged, or --stage option is in use, this option has no effect.
// --deduplicate
func Deduplicate(g *types.Cmd) {
	g.AddOptions("--deduplicate")
}

// Deleted Show deleted files in the output
// -d, --deleted
func Deleted(g *types.Cmd) {
	g.AddOptions("--deleted")
}

// Directory If a whole directory is classified as "other", show just its name (with a trailing slash) and not its whole contents.
// --directory
func Directory(g *types.Cmd) {
	g.AddOptions("--directory")
}

// Eol Show <eolinfo> and <eolattr> of files. <eolinfo> is the file content identification used by Git when the "text" attribute is "auto" (or not set and core.autocrlf is not false). <eolinfo> is either "-text", "none", "lf", "crlf", "mixed" or "".\n"" means the file is not a regular file, it is not in the index or not accessible in the working tree.
// <eolattr> is the attribute that is used when checking out or committing, it is either "", "-text", "text", "text=auto", "text eol=lf", "text eol=crlf". Since Git 2.10 "text=auto eol=lf" and "text=auto eol=crlf" are supported.
// Both the <eolinfo> in the index ("i/<eolinfo>") and in the working tree ("w/<eolinfo>") are shown for regular files, followed by the ("attr/<eolattr>").
// --eol
func Eol(g *types.Cmd) {
	g.AddOptions("--eol")
}

// ErrorUnmatch If any <file> does not appear in the index, treat this as an error (return 1).
// --error-unmatch
func ErrorUnmatch(g *types.Cmd) {
	g.AddOptions("--error-unmatch")
}

// Exclude Skip untracked files matching pattern. Note that pattern is a shell wildcard pattern. See EXCLUDE PATTERNS below for more information.
// -x <pattern>, --exclude=<pattern>
func Exclude(pattern string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--exclude=%s", pattern))
	}
}

// ExcludeFrom Read exclude patterns from <file>; 1 per line.
// -X <file>, --exclude-from=<file>
func ExcludeFrom(file string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--exclude-from=%s", file))
	}
}

// ExcludePerDirectory Read additional exclude patterns that apply only to the directory and its subdirectories in <file>.
// --exclude-per-directory=<file>
func ExcludePerDirectory(file string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--exclude-per-directory=%s", file))
	}
}

// ExcludeStandard Add the standard Git exclusions: .git/info/exclude, .gitignore in each directory, and the user’s global exclusion file.
// --exclude-standard
func ExcludeStandard(g *types.Cmd) {
	g.AddOptions("--exclude-standard")
}

// F Similar to -t, but use lowercase letters for files that are marked as fsmonitor valid (see git-update-index(1)).
// -f
func F(g *types.Cmd) {
	g.AddOptions("-f")
}

// Format A string that interpolates %(fieldname) from the result being shown. It also interpolates %% to %, and %xx where xx are hex digits interpolates to character with hex code xx; for example %00 interpolates to \0 (NUL), %09 to \t (TAB) and %0a to \n (LF). --format cannot be combined with -s, -o, -k, -t, --resolve-undo and --eol.
// --format=<format>
func Format(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--format=%s", value))
	}
}

// FullName When run from a subdirectory, the command usually outputs paths relative to the current directory. This option forces paths to be output relative to the project top directory.
// --full-name
func FullName(g *types.Cmd) {
	g.AddOptions("--full-name")
}

// Ignored Show only ignored files in the output. When showing files in the index, print only those matched by an exclude pattern. When showing "other" files, show only those matched by an exclude pattern. Standard ignore rules are not automatically activated, therefore at least one of the --exclude* options is required.
// -i, --ignored
func Ignored(g *types.Cmd) {
	g.AddOptions("--ignored")
}

// Killed Show files on the filesystem that need to be removed due to file/directory conflicts for checkout-index to succeed.
// -k, --killed
func Killed(g *types.Cmd) {
	g.AddOptions("--killed")
}

// Modified Show modified files in the output
// -m, --modified
func Modified(g *types.Cmd) {
	g.AddOptions("--modified")
}

// NoEmptyDirectory Do not list empty directories. Has no effect without --directory.
// --no-empty-directory
func NoEmptyDirectory(g *types.Cmd) {
	g.AddOptions("--no-empty-directory")
}

// Others Show other (i.e. untracked) files in the output
// -o, --others
func Others(g *types.Cmd) {
	g.AddOptions("--others")
}

// RecurseSubmodules Recursively calls ls-files on each active submodule in the repository. Currently there is only support for the --cached and --stage modes.
// --recurse-submodules
func RecurseSubmodules(g *types.Cmd) {
	g.AddOptions("--recurse-submodules")
}

// Sparse If the index is sparse, show the sparse directories without expanding to the contained files. Sparse directories will be shown with a trailing slash, such as "x/" for a sparse directory •"x•".
// --sparse
func Sparse(g *types.Cmd) {
	g.AddOptions("--sparse")
}

// Stage Show staged contents' mode bits, object name and stage number in the output.
// -s, --stage
func Stage(g *types.Cmd) {
	g.AddOptions("--stage")
}

// T This feature is semi-deprecated. For scripting purpose, git-status(1) --porcelain and git-diff-files(1) --name-status are almost always superior alternatives, and users should look at git-status(1) --short or git-diff(1) --name-status for more user-friendly alternatives.
// -t
func T(g *types.Cmd) {
	g.AddOptions("-t")
}

// Unmerged Show unmerged files in the output (forces --stage)
// -u, --unmerged
func Unmerged(g *types.Cmd) {
	g.AddOptions("--unmerged")
}

// V Similar to -t, but use lowercase letters for files that are marked as assume unchanged (see git-update-index(1)).
// -v
func V(g *types.Cmd) {
	g.AddOptions("-v")
}

// WithTree When using --error-unmatch to expand the user supplied <file> (i.e. path pattern) arguments to paths, pretend that paths which were removed in the index since the named <tree-ish> are still present. Using this option with -s or -u options does not make any sense.
// --with-tree=<tree-ish>
func WithTree(treeIsh string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--with-tree=%s", treeIsh))
	}
}

// Z \0 line termination on output and do not quote filenames. See OUTPUT below for more information.
// -z
func Z(g *types.Cmd) {
	g.AddOptions("-z")
}
