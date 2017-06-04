package config

import "github.com/ldez/go-git-cmd-wrapper/git"

// Config https://git-scm.com/docs/git-config

func Entry(key string, value string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(key)
		g.AddOptions(value)
	}
}
