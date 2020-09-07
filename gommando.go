package gommando

import (
	"os"
	"os/exec"

	"github.com/matti/dynamicmultiwriter"
	"github.com/matti/gommando/internal/chain"
)

// Gommando ...
type Gommando struct {
	chains []*chain.Chain
	out    *dynamicmultiwriter.DynamicMultiWriter
	err    *dynamicmultiwriter.DynamicMultiWriter
	both   *dynamicmultiwriter.DynamicMultiWriter

	cmd *exec.Cmd
}

// New ...
func New(cmd string) *Gommando {
	g := &Gommando{}

	g.out = dynamicmultiwriter.New()
	g.err = dynamicmultiwriter.New()
	g.both = dynamicmultiwriter.New()

	g.err.Add(g.both)
	g.out.Add(g.both)

	g.Output(true)

	g.cmd = exec.Command("/usr/bin/env", "sh", "-c", cmd)
	g.cmd.Stdout = g.out
	g.cmd.Stderr = g.err

	return g
}

// Output ...
func (g *Gommando) Output(enabled bool) {
	if enabled {
		g.out.Add(os.Stdout)
		g.err.Add(os.Stderr)
	} else {
		g.out.Remove(os.Stdout)
		g.err.Remove(os.Stderr)
	}
}

// Stdout ...
func (g *Gommando) Stdout() *chain.Chain {
	c := chain.New(g.out, nil, nil)
	g.chains = append(g.chains, c)
	return c
}

// Stderr ...
func (g *Gommando) Stderr() *chain.Chain {
	c := chain.New(g.err, nil, nil)
	g.chains = append(g.chains, c)

	return c
}

// Stdboth ...
func (g *Gommando) Stdboth() *chain.Chain {
	c := chain.New(g.both, nil, nil)
	g.chains = append(g.chains, c)

	return c
}

// Run ...
func (g *Gommando) Run() {
	for _, c := range g.chains {
		c.Start()
	}

	g.cmd.Run()

	for _, c := range g.chains {
		c.Close()
	}
}
