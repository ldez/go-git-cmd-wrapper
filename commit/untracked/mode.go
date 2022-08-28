/*
Package untracked The mode parameter is optional (defaults to all), and is used to specify the handling of untracked files;
when -u is not used, the default is normal, i.e. show untracked files and directories.

	-u[<mode>], --untracked-files[=<mode>]
*/
package untracked

const (
	// No Show no untracked files.
	No = "no"

	// Normal Shows untracked files and directories.
	Normal = "normal"

	// All Also shows individual files in untracked directories.
	All = "all"
)
