// --cleanup=<mode>
// This option determines how the supplied commit message should be cleaned up before committing. The <mode> can be strip, whitespace, verbatim, scissors or default.
package cleanup

const (
	// Strip leading and trailing empty lines, trailing whitespace, commentary and collapse consecutive empty lines.
	Strip = "strip"

	// Same as strip except #commentary is not removed.
	Whitespace = "whitespace"

	// Do not change the message at all.
	Verbatim = "verbatim"

	// Same as whitespace, except that everything from (and including) the line "# ------------------------ >8 ------------------------" is truncated if the message is to be edited. "#" can be customized with core.commentChar.
	Scissors = "scissors"

	// Same as strip if the message is to be edited. Otherwise whitespace.
	Default = "default"
)
