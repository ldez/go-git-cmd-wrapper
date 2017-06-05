// Config https://git-scm.com/docs/git-config
package config

import "github.com/ldez/go-git-cmd-wrapper/types"


func Entry(key string, value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(key)
		g.AddOptions(value)
	}
}
