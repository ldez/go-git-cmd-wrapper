package types

type Cmd struct {
	Debug   bool
	Base    string
	Options []string
}

type Option func(g *Cmd)

func (g *Cmd) AddOptions(option string) {
	g.Options = append(g.Options, option)
}

