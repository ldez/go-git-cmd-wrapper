package git

import (
	"log"
	"os/exec"
	"strings"

	"github.com/ldez/go-git-cmd-wrapper/types"
)

var cmdExecutor = commander

// Init https://git-scm.com/docs/git-init
func Init(options ...types.Option) (string, error) {
	return command("init", options...)
}

// Push https://git-scm.com/docs/git-push
func Push(options ...types.Option) (string, error) {
	return command("push", options...)
}

// Clone https://git-scm.com/docs/git-clone
func Clone(options ...types.Option) (string, error) {
	return command("clone", options...)
}

// Remote https://git-scm.com/docs/git-remote
func Remote(options ...types.Option) (string, error) {
	return command("remote", options...)
}

// Fetch https://git-scm.com/docs/git-fetch
func Fetch(options ...types.Option) (string, error) {
	return command("fetch", options...)
}

// Rebase https://git-scm.com/docs/git-rebase
func Rebase(options ...types.Option) (string, error) {
	return command("rebase", options...)
}

// Checkout https://git-scm.com/docs/git-checkout
func Checkout(options ...types.Option) (string, error) {
	return command("checkout", options...)
}

// Config https://git-scm.com/docs/git-config
func Config(options ...types.Option) (string, error) {
	return command("config", options...)
}

// Branch https://git-scm.com/docs/git-branch
func Branch(options ...types.Option) (string, error) {
	return command("branch", options...)
}

// Debug display command line
func Debug(g *types.Cmd) {
	g.Debug = true
}

// DebugBool display command line
func Debugger(debug bool) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.Debug = debug
	}
}

// Cond apply conditionally an option
func Cond(apply bool, option types.Option) func(*types.Cmd) {
	if apply {
		return option
	}
	return NoOp
}

// NoOp do nothing.
func NoOp(g *types.Cmd) {}

func command(name string, options ...types.Option) (string, error) {
	g := &types.Cmd{Base: "git", Options: []string{name}}
	for _, opt := range options {
		opt(g)
	}
	return cmdExecutor(g.Base, g.Debug, g.Options...)
}

func commander(name string, debug bool, args ...string) (string, error) {
	if debug {
		log.Println(name, strings.Join(args, " "))
	}

	cmd := exec.Command(name, args...)

	output, err := cmd.CombinedOutput()

	return string(output), err
}
