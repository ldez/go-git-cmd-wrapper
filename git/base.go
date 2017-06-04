package git

import (
	"log"
	"os/exec"
	"strings"
)

type Cmd struct {
	Debug   bool
	base    string
	options []string
}

func (g *Cmd) AddOptions(option string) {
	g.options = append(g.options, option)
}

// Init https://git-scm.com/docs/git-init
func Init(options ...func(*Cmd)) (string, error) {
	return command("init", options...)
}

// Push https://git-scm.com/docs/git-push
func Push(options ...func(*Cmd)) (string, error) {
	return command("push", options...)
}

// Clone https://git-scm.com/docs/git-clone
func Clone(options ...func(*Cmd)) (string, error) {
	return command("clone", options...)
}

// Remote https://git-scm.com/docs/git-remote
func Remote(options ...func(*Cmd)) (string, error) {
	return command("remote", options...)
}

// Fetch https://git-scm.com/docs/git-fetch
func Fetch(options ...func(*Cmd)) (string, error) {
	return command("fetch", options...)
}

// Rebase https://git-scm.com/docs/git-rebase
func Rebase(options ...func(*Cmd)) (string, error) {
	return command("rebase", options...)
}

// Checkout https://git-scm.com/docs/git-checkout
func Checkout(options ...func(*Cmd)) (string, error) {
	return command("checkout", options...)
}

// Config https://git-scm.com/docs/git-config
func Config(options ...func(*Cmd)) (string, error) {
	return command("config", options...)
}

// Debug display command line
func Debug(g *Cmd) {
	g.Debug = true
}

// Debugger display command line
func Debugger(debug bool) func(*Cmd) {
	return func(g *Cmd) {
		g.Debug = debug
	}
}

func command(name string, options ...func(*Cmd)) (string, error) {
	g := &Cmd{base: "git", options: []string{name}}
	for _, opt := range options {
		opt(g)
	}
	return commander(g.base, g.Debug, g.options...)
}

func commander(name string, debug bool, args ...string) (string, error) {
	if debug {
		log.Println(name, strings.Join(args, " "))
	}

	cmd := exec.Command(name, args...)

	output, err := cmd.CombinedOutput()

	return string(output), err
}
