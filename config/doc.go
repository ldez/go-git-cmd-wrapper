/*
Package config git-config - Get and set repository or global options.

# SYNOPSIS

Reference: https://git-scm.com/docs/git-config

	git config [<file-option>] [type] [--show-origin] [-z|--null] name [value [value_regex]]
	git config [<file-option>] [type] --add name value
	git config [<file-option>] [type] --replace-all name value [value_regex]
	git config [<file-option>] [type] [--show-origin] [-z|--null] --get name [value_regex]
	git config [<file-option>] [type] [--show-origin] [-z|--null] --get-all name [value_regex]
	git config [<file-option>] [type] [--show-origin] [-z|--null] [--name-only] --get-regexp name_regex [value_regex]
	git config [<file-option>] [type] [-z|--null] --get-urlmatch name URL
	git config [<file-option>] --unset name [value_regex]
	git config [<file-option>] --unset-all name [value_regex]
	git config [<file-option>] --rename-section old_name new_name
	git config [<file-option>] --remove-section name
	git config [<file-option>] [--show-origin] [-z|--null] [--name-only] -l | --list
	git config [<file-option>] --get-color name [default]
	git config [<file-option>] --get-colorbool name [stdout-is-tty]
	git config [<file-option>] -e | --edit

# DESCRIPTION

You can query/set/replace/unset options with this command. The name is actually the section and the key separated by a dot, and the value will be escaped.

Multiple lines can be added to an option by using the --add option.
If you want to update or unset an option which can occur on multiple lines, a POSIX regexp value_regex needs to be given.
Only the existing values that match the regexp are updated or unset.
If you want to handle the lines that do not match the regex, just prepend a single exclamation mark in front (see also the section called “EXAMPLES”).

The type specifier can be either --int or --bool, to make git config ensure that the variable(s) are of the given type and convert the value to the canonical form (simple decimal number for int, a "true" or "false" string for bool), or --path, which does some path expansion (see --path below).
If no type specifier is passed, no checks or transformations are performed on the value.

When reading, the values are read from the system, global and repository local configuration files by default, and options --system, --global, --local and --file <filename> can be used to tell the command to read from only that location (see the section called “FILES”).

When writing, the new value is written to the repository local configuration file by default, and options --system, --global, --file <filename> can be used to tell the command to write to that location (you can say --local but that is the default).

This command will fail with non-zero status upon error. Some exit codes are:

- The section or key is invalid (ret=1),
- no section or name was provided (ret=2),
- the config file is invalid (ret=3),
- the config file cannot be written (ret=4),
- you try to unset an option which does not exist (ret=5),
- you try to unset/set an option for which multiple lines match (ret=5), or
- you try to use an invalid regexp (ret=6).

On success, the command returns the exit code 0.
*/
package config
