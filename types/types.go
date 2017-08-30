package types

// Cmd Command.
type Cmd struct {
	Debug   bool
	Base    string
	Options []string
}

// Option Command option.
type Option func(g *Cmd)

// AddOptions Add one command option.
func (g *Cmd) AddOptions(option string) {
	g.Options = append(g.Options, option)
}

// ApplyOptions Apply command options.
func (g *Cmd) ApplyOptions(options ...Option) {
	for _, opt := range options {
		opt(g)
	}
}
