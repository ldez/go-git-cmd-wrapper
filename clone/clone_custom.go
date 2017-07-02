package clone

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/types"
)

// Config Set a configuration variable in the newly-created repository; this takes effect immediately after the repository is initialized, but before the remote history is fetched or any files checked out. The key is in the same format as expected by git-config(1) (e.g., core.eol=true). If multiple values are given for the same key, each value will be written to the config file. This makes it safe, for example, to add additional fetch refspecs to the origin remote.
// --config <key>=<value>, -c <key>=<value>
func Config(key, value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--config")
		g.AddOptions(fmt.Sprintf("%s=%s", key, value))
	}
}

// Repository The (possibly remote) repository to clone from. See the URLS section below for more information on specifying repositories.
// <repository>
func Repository(url string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(url)
	}
}

// Directory The name of a new directory to clone into. The "humanish" part of the source repository is used if no directory is explicitly given (repo for /path/to/repo.git and foo for host.xz:foo/.git). Cloning into an existing directory is only allowed if the directory is empty.
// <directory>
func Directory(name string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(name)
	}
}

// --reference[-if-able] <repository>
// If the reference repository is on the local machine, automatically setup .git/objects/info/alternates to obtain objects from the reference repository. Using an already existing repository as an alternate will require fewer objects to be copied from the repository being cloned, reducing network and local storage costs. When using the --reference-if-able, a non existing directory is skipped with a warning instead of aborting the clone.
