package config

import (
	"strconv"

	"github.com/ldez/go-git-cmd-wrapper/types"
)

// Entry Adds a configuration entry.
func Entry(key string, value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(key)
		g.AddOptions(value)
	}
}

// Add Adds a new line to the option without altering any existing values. This is the same as providing ^$ as the value_regex in --replace-all.
// --add
func Add(name, value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--add")
		g.AddOptions(name)
		g.AddOptions(value)
	}
}

// ReplaceAll Default behavior is to replace at most one line. This replaces all lines matching the key (and optionally the value_regex).
// --replace-all
func ReplaceAll(name, value, valueRegex string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--replace-all")
		g.AddOptions(name)
		g.AddOptions(value)
		if len(valueRegex) != 0 {
			g.AddOptions(valueRegex)
		}
	}
}

// Get the value for a given key (optionally filtered by a regex matching the value). Returns error code 1 if the key was not found and the last value if multiple key values were found.
// --get
func Get(name, valueRegex string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--get")
		g.AddOptions(name)
		if len(valueRegex) != 0 {
			g.AddOptions(valueRegex)
		}
	}
}

// GetAll Like get, but returns all values for a multi-valued key.
// --get-all
func GetAll(name, valueRegex string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--get-all")
		g.AddOptions(name)
		if len(valueRegex) != 0 {
			g.AddOptions(valueRegex)
		}
	}
}

// GetRegexp Like --get-all, but interprets the name as a regular expression and writes out the key names. Regular expression matching is currently case-sensitive and done against a canonicalized version of the key in which section and variable names are lowercased, but subsection names are not.
//--get-regexp
func GetRegexp(nameRegexp, valueRegex string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--get-regexp")
		g.AddOptions(nameRegexp)
		if len(valueRegex) != 0 {
			g.AddOptions(valueRegex)
		}
	}
}

// GetURLMatch When given a two-part name section.key, the value for section.<url>.key whose <url> part matches the best to the given URL is returned (if no such key exists, the value for section.key is used as a fallback). When given just the section as name, do so for all the keys in the section and list them. Returns error code 1 if no value is found.
// --get-urlmatch name URL
func GetURLMatch(name, url string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--get-urlmatch")
		g.AddOptions(name)
		g.AddOptions(url)
	}
}

// Unset Remove the line matching the key from config file.
// --unset
func Unset(name, valueRegex string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--unset")
		g.AddOptions(name)
		if len(valueRegex) != 0 {
			g.AddOptions(valueRegex)
		}
	}
}

// UnsetAll Remove all lines matching the key from config file.
// --unset-all
func UnsetAll(name, valueRegex string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--unset-all")
		g.AddOptions(name)
		if len(valueRegex) != 0 {
			g.AddOptions(valueRegex)
		}
	}
}

// RenameSection Rename the given section to a new name.
// --rename-section
func RenameSection(oldName, newName string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--rename-section")
		g.AddOptions(oldName)
		g.AddOptions(newName)
	}
}

// RemoveSection Remove the given section from the configuration file.
// --remove-section
func RemoveSection(name string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--remove-section")
		g.AddOptions(name)
	}
}

// GetColor Find the color configured for name (e.g.  color.diff.new) and output it as the ANSI color escape sequence to the standard output. The optional default parameter is used instead, if there is no color configured for name.
// --get-color name [default]
func GetColor(name, defaultValue string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--get-color")
		g.AddOptions(name)
		if len(defaultValue) != 0 {
			g.AddOptions(defaultValue)
		}
	}
}

// GetColorBool Find the color setting for name (e.g.  color.diff) and output "true" or "false".  stdout-is-tty should be either "true" or "false", and is taken into account when configuration says "auto". If stdout-is-tty is missing, then checks the standard output of the command itself, and exits with status 0 if color is to be used, or exits with status 1 otherwise. When the color setting for name is undefined, the command uses color.ui as fallback.
// --get-colorbool name [stdout-is-tty]
func GetColorBool(name string, stdoutIsTTY bool) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--get-colorbool")
		g.AddOptions(name)
		g.AddOptions(strconv.FormatBool(stdoutIsTTY))
	}
}

//--[no-]includes
//Respect include.*  directives in config files when looking up values. Defaults to off when a specific file is given (e.g., using --file, --global, etc) and on when searching all config files.
