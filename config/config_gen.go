// Code generated by pkg/commands/internal/migrate/cloner/cloner.go. DO NOT EDIT.

package config

import "github.com/ldez/go-git-cmd-wrapper/v2/types"

// Blob Similar to --file but use the given blob instead of a file. E.g. you can use master:.gitmodules to read values from the file .gitmodules in the master branch. See 'SPECIFYING REVISIONS' section in gitrevisions(7) for a more complete list of ways to spell blob names.
// --blob <blob>
func Blob(value string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("--blob")
		g.AddOptions(value)
	}
}

// Bool git config will ensure that the output is 'true' or 'false'
// --bool
func Bool(g *types.Cmd) {
	g.AddOptions("--bool")
}

// BoolOrInt git config will ensure that the output matches the format of either --bool or --int, as described above.
// --bool-or-int
func BoolOrInt(g *types.Cmd) {
	g.AddOptions("--bool-or-int")
}

// Edit Opens an editor to modify the specified config file; either --system, --global, or repository (default).
// -e, --edit
func Edit(g *types.Cmd) {
	g.AddOptions("--edit")
}

// File Use the given config file instead of the one specified by GIT_CONFIG.
// -f <config-file>, --file <config-file>
func File(configFile string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("--file")
		g.AddOptions(configFile)
	}
}

// Global For writing options: write to global ~/.gitconfig file rather than the repository .git/config, write to $XDG_CONFIG_HOME/git/config file if this file exists and the ~/.gitconfig file doesn’t.
// --global
func Global(g *types.Cmd) {
	g.AddOptions("--global")
}

// Int git config will ensure that the output is a simple decimal number. An optional value suffix of k, m, or g in the config file will cause the value to be multiplied by 1024, 1048576, or 1073741824 prior to output.
// --int
func Int(g *types.Cmd) {
	g.AddOptions("--int")
}

// List List all variables set in config file, along with their values.
// -l, --list
func List(g *types.Cmd) {
	g.AddOptions("--list")
}

// Local For writing options: write to the repository .git/config file. This is the default behavior.
// --local
func Local(g *types.Cmd) {
	g.AddOptions("--local")
}

// NameOnly Output only the names of config variables for --list or --get-regexp.
// --name-only
func NameOnly(g *types.Cmd) {
	g.AddOptions("--name-only")
}

// Null For all options that output values and/or keys, always end values with the null character (instead of a newline). Use newline instead as a delimiter between key and value. This allows for secure parsing of the output without getting confused e.g. by values that contain line breaks.
// -z, --null
func Null(g *types.Cmd) {
	g.AddOptions("--null")
}

// Path git-config will expand leading ~ to the value of $HOME, and ~user to the home directory for the specified user. This option has no effect when setting the value (but you can use git config bla ~/ from the command line to let your shell do the expansion).
// --path
func Path(g *types.Cmd) {
	g.AddOptions("--path")
}

// ShowOrigin Augment the output of all queried config options with the origin type (file, standard input, blob, command line) and the actual origin (config file path, ref, or blob id if applicable).
// --show-origin
func ShowOrigin(g *types.Cmd) {
	g.AddOptions("--show-origin")
}

// System For writing options: write to system-wide $(prefix)/etc/gitconfig rather than the repository .git/config.
// --system
func System(g *types.Cmd) {
	g.AddOptions("--system")
}
