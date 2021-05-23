package notes

import "github.com/ldez/go-git-cmd-wrapper/v2/types"

// List the notes object for a given object.
// usage: git notes [list [<object>]]
func List(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("list")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Add notes for a given object (defaults to `HEAD`).
// usage: git notes add [<options>] [<object>]
//
//     -m, --message <message>
//                           note contents as a string
//     -F, --file <file>     note contents in a file
//     -c, --reedit-message <object>
//                           reuse and edit specified note object
//     -C, --reuse-message <object>
//                           reuse specified note object
//     --allow-empty         allow storing empty note
//     -f, --force           replace existing notes
func Add(object string, opts ...types.Option) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("add")

		for _, opt := range opts {
			opt(g)
		}

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Copy the notes for the first object onto the second object (defaults to `HEAD`).
// usage: git notes copy [<options>] <from-object> <to-object>
//    or: git notes copy --stdin [<from-object> <to-object>]...
//
//     -f, --force           replace existing notes
//     --stdin               read objects from stdin
//     --for-rewrite <command>
//                           load rewriting config for <command> (implies --stdin)
func Copy(opts ...types.Option) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("copy")

		for _, opt := range opts {
			opt(g)
		}
	}
}

// Append to the notes of an existing object (defaults to `HEAD`). Creates a new notes object if needed.
// usage: git notes append [<options>] [<object>]
//
//     -m, --message <message>
//                           note contents as a string
//     -F, --file <file>     note contents in a file
//     -c, --reedit-message <object>
//                           reuse and edit specified note object
//     -C, --reuse-message <object>
//                           reuse specified note object
//     --allow-empty         allow storing empty note
func Append(object string, opts ...types.Option) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("append")

		for _, opt := range opts {
			opt(g)
		}

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Edit the notes for a given object (defaults to `HEAD`).
// usage: git notes edit [<object>]
//
//     -m, --message <message>
//                           note contents as a string
//     -F, --file <file>     note contents in a file
//     -c, --reedit-message <object>
//                           reuse and edit specified note object
//     -C, --reuse-message <object>
//                           reuse specified note object
//     --allow-empty         allow storing empty note
func Edit(object string, opts ...types.Option) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("edit")

		for _, opt := range opts {
			opt(g)
		}

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Show the notes for a given object (defaults to `HEAD`).
// usage: git notes show [<object>]
func Show(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("show")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Merge the given notes ref into the current notes ref.
// usage: git notes merge [<options>] <notes-ref>
//    or: git notes merge --commit [<options>]
//    or: git notes merge --abort [<options>]
//
// General options
//     -v, --verbose         be more verbose
//     -q, --quiet           be more quiet
//
// Merge options
//     -s, --strategy <strategy>
//                           resolve notes conflicts using the given strategy (manual/ours/theirs/union/cat_sort_uniq)
//
// Committing unmerged notes
//     --commit              finalize notes merge by committing unmerged notes
//
// Aborting notes merge resolution
//     --abort               abort notes merge
func Merge(opts ...types.Option) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("merge")

		for _, opt := range opts {
			opt(g)
		}
	}
}

// Remove the notes for given objects (defaults to `HEAD`).
// usage: git notes remove [<object>]
//
//     --ignore-missing      attempt to remove non-existent note is not an error
//     --stdin               read object names from the standard input
func Remove(object string, opts ...types.Option) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("remove")

		for _, opt := range opts {
			opt(g)
		}

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Prune Remove all notes for non-existing/unreachable objects.
// usage: git notes prune [<options>]
//
//     -n, --dry-run         do not remove, show only
//     -v, --verbose         report pruned notes
func Prune(opts ...types.Option) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("prune")
	}
}

// GetRef Print the current notes ref. This provides an easy way to retrieve the current notes ref (e.g. from scripts).
// usage: git notes get-ref
func GetRef(opts ...types.Option) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("get-ref")
	}
}
