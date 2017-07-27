// -u[<mode>], --untracked-files[=<mode>]
// The mode parameter is optional (defaults to all), and is used to specify the handling of untracked files; when -u is not used, the default is normal, i.e. show untracked files and directories.
package untracked

const (
	// Show no untracked files
	No = "no"

	// Shows untracked files and directories
	Normal = "normal"

	// Also shows individual files in untracked directories.
	All = "all"
)
