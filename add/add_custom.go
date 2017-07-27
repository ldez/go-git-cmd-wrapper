package add

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/types"
)

// HyphenHyphen add `--`
// This option can be used to separate command-line options from the list of files, (useful when filenames might be mistaken for command-line options).
func HyphenHyphen(g *types.Cmd) {
	g.AddOptions("--")
}

// PathSpec [<pathspec>...]
// Files to add content from.
// Fileglobs (e.g.  *.c) can be given to add all matching files.
// Also a leading directory name (e.g.  dir to add dir/file1 and dir/file2) can be given to update the index to match the current state of the directory as a whole (e.g. specifying dir will record not just a file dir/file1 modified in the working tree, a file dir/file2 added to the working tree, but also a file dir/file3 removed from the working tree.
func PathSpec(paths ...string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		for _, path := range paths {
			g.AddOptions(path)
		}
	}
}

// Chmod Override the executable bit of the added files.
// The executable bit is only changed in the index, the files on disk are left unchanged.
// --chmod=(+|-)x
func Chmod(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--chmod=%sx", value))
	}
}
