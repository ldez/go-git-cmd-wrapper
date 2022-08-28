package tag

// CODE GENERATED AUTOMATICALLY
// THIS FILE MUST NOT BE EDITED BY HAND

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/v2/types"
)

// Annotate Make an unsigned, annotated tag object.
// -a, --annotate
func Annotate(g *types.Cmd) {
	g.AddOptions("--annotate")
}

// Cleanup This option sets how the tag message is cleaned up. The <mode> can be one of verbatim, whitespace and strip. The strip mode is default. The verbatim mode does not change message at all, whitespace
//
//	removes just leading/trailing whitespace lines and strip removes both whitespace and commentary.
//
// --cleanup=<mode>
func Cleanup(mode string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--cleanup=%s", mode))
	}
}

// Color Respect any colors specified in the --format option. The <when> field must be one of always, never, or auto (if <when> is absent, behave as if always was given).
// --color[=<when>]
func Color(when string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(when) == 0 {
			g.AddOptions("--color")
		} else {
			g.AddOptions(fmt.Sprintf("--color=%s", when))
		}
	}
}

// Column Display tag listing in columns. See configuration variable column.tag for option syntax.--column and --no-column without options are equivalent to always and never respectively.
//
//	This option is only applicable when listing tags without annotation lines.
//
// --column[=<options>]
func Column(options string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(options) == 0 {
			g.AddOptions("--column")
		} else {
			g.AddOptions(fmt.Sprintf("--column=%s", options))
		}
	}
}

// Contains Only list tags which contain the specified commit (HEAD if not specified). Implies --list.
// --contains [<commit>]
func Contains(commit string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(commit) == 0 {
			g.AddOptions("--contains")
		} else {
			g.AddOptions(fmt.Sprintf("--contains=%s", commit))
		}
	}
}

// CreateReflog Create a reflog for the tag. To globally enable reflogs for tags, see core.logAllRefUpdates in git-config(1). The negated form --no-create-reflog only overrides an earlier --create-reflog, but
//
//	currently does not negate the setting of core.logAllRefUpdates.
//
// --create-reflog
func CreateReflog(g *types.Cmd) {
	g.AddOptions("--create-reflog")
}

// Edit The message taken from file with -F and command line with -m are usually used as the tag message unmodified.
// This option lets you further edit the message taken from these sources.
// -e, --edit
func Edit(g *types.Cmd) {
	g.AddOptions("--edit")
}

// File Take the tag message from the given file. Use - to read the message from the standard input. Implies -a if none of -a, -s, or -u <keyid> is given.
// -F <file>, --file=<file>
func File(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--file=%s", value))
	}
}

// Force Replace an existing tag with the given name (instead of failing).
// -f, --force
func Force(g *types.Cmd) {
	g.AddOptions("--force")
}

// IgnoreCase Sorting and filtering tags are case insensitive.
// -i, --ignore-case
func IgnoreCase(g *types.Cmd) {
	g.AddOptions("--ignore-case")
}

// List List tags. With optional <pattern>..., e.g.  git tag --list 'v-*', list only the tags that match the pattern(s).
//
// Running "git tag" without arguments also lists all tags. The pattern is a shell wildcard (i.e., matched using fnmatch(3)). Multiple patterns may be given; if any of them matches, the tag is shown.
//
// This option is implicitly supplied if any other list-like option such as --contains is provided. See the documentation for each of those options for details.
// -l, --list
func List(g *types.Cmd) {
	g.AddOptions("--list")
}

// LocalUser Make a GPG-signed tag, using the given key.
// -u <keyid>, --local-user=<keyid>
func LocalUser(keyID string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--local-user=%s", keyID))
	}
}

// Merged Only list tags whose commits are reachable from the specified commit (HEAD if not specified), incompatible with --no-merged.
// --merged [<commit>]
func Merged(commit string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(commit) == 0 {
			g.AddOptions("--merged")
		} else {
			g.AddOptions(fmt.Sprintf("--merged=%s", commit))
		}
	}
}

// Message Use the given tag message (instead of prompting).
// If multiple -m options are given, their values are concatenated as separate paragraphs. Implies -a if none of -a, -s, or -u <keyid> is given.
// -m <msg>, --message=<msg>
func Message(msg string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--message=%s", msg))
	}
}

// N <num> specifies how many lines from the annotation, if any, are printed when using -l. Implies --list.
//
// The default is not to print any annotation lines. If no number is given to -n, only the first line is printed. If the tag is not annotated, the commit message is displayed instead.
// -n<num>
func N(num string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("-n")
		g.AddOptions(num)
	}
}

// NoColumn Display tag listing in columns. See configuration variable column.tag for option syntax.--column and --no-column without options are equivalent to always and never respectively.
//
//	This option is only applicable when listing tags without annotation lines.
//
// --no-column
func NoColumn(g *types.Cmd) {
	g.AddOptions("--no-column")
}

// NoContains Only list tags which don’t contain the specified commit (HEAD if not specified). Implies --list.
// --no-contains [<commit>]
func NoContains(commit string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(commit) == 0 {
			g.AddOptions("--no-contains")
		} else {
			g.AddOptions(fmt.Sprintf("--no-contains=%s", commit))
		}
	}
}

// NoMerged Only list tags whose commits are not reachable from the specified commit (HEAD if not specified), incompatible with --merged.
// --no-merged [<commit>]
func NoMerged(commit string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(commit) == 0 {
			g.AddOptions("--no-merged")
		} else {
			g.AddOptions(fmt.Sprintf("--no-merged=%s", commit))
		}
	}
}

// PointsAt Only list tags of the given object (HEAD if not specified). Implies --list.
// --points-at <object>
func PointsAt(object string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--points-at")
		g.AddOptions(object)
	}
}

// Sign Make a GPG-signed tag, using the default e-mail address’s key.
// -s, --sign
func Sign(g *types.Cmd) {
	g.AddOptions("--sign")
}

// Sort Sort based on the key given. Prefix - to sort in descending order of the value. You may use the --sort=<key> option multiple times, in which case the last key becomes the primary key. Also
// supports "version:refname" or "v:refname" (tag names are treated as versions). The "version:refname" sort order can also be affected by the "versionsort.suffix" configuration variable. The keys
// supported are the same as those in git for-each-ref. Sort order defaults to the value configured for the tag.sort variable if it exists, or lexicographic order otherwise. See git-config(1).
// --sort=<key>
func Sort(key string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--sort=%s", key))
	}
}

// Verify Verify the GPG signature of the given tag names.
// -v, --verify
func Verify(g *types.Cmd) {
	g.AddOptions("--verify")
}
