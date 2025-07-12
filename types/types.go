package types

import (
	"context"
	"log"
	"os"
	"os/exec"
	"strings"
)

type logger interface {
	Fatal(args ...any)
	Fatalf(format string, args ...any)
	Print(args ...any)
	Println(args ...any)
	Printf(format string, args ...any)
}

// Executor The Git command call function.
type Executor func(ctx context.Context, name string, debug bool, args ...string) (string, error)

// Cmd Command.
type Cmd struct {
	Debug       bool
	Base        string
	BaseOptions []string
	Options     []string
	Logger      logger
	Executor    Executor
}

// NewCmd Creates a new Cmd.
func NewCmd(name string) *Cmd {
	g := &Cmd{
		Debug:   false,
		Base:    "git",
		Options: []string{name},
		Logger:  log.New(os.Stdout, "", 0),
	}
	g.Executor = defaultExecutor(g)

	return g
}

// Option Command option.
type Option func(g *Cmd)

// AddOptions Add one command option.
func (g *Cmd) AddOptions(option string) {
	g.Options = append(g.Options, option)
}

// AddBaseOptions Add one global command option.
func (g *Cmd) AddBaseOptions(option string) {
	g.BaseOptions = append(g.BaseOptions, option)
}

// ApplyOptions Apply command options.
func (g *Cmd) ApplyOptions(options ...Option) {
	for _, opt := range options {
		opt(g)
	}
}

// Exec Execute the Git command call.
func (g *Cmd) Exec(ctx context.Context, name string, debug bool, args ...string) (string, error) {
	return g.Executor(ctx, name, debug, args...)
}

func defaultExecutor(g *Cmd) Executor {
	return func(ctx context.Context, name string, debug bool, args ...string) (string, error) {
		if debug {
			g.Logger.Println(name, strings.Join(args, " "))
		}

		output, err := exec.CommandContext(ctx, name, args...).CombinedOutput()

		return string(output), err
	}
}
