package gommando

import (
	"os"
	"os/exec"

	"github.com/matti/dynamicmultiwriter"
	"github.com/matti/gommando/internal/chain"
)

// Gommando ...
type Gommando struct {
	stdout  *dynamicmultiwriter.DynamicMultiWriter
	stderr  *dynamicmultiwriter.DynamicMultiWriter
	stdboth *dynamicmultiwriter.DynamicMultiWriter

	cmd *exec.Cmd
}

// New ...
func New(cmd string) *Gommando {
	g := &Gommando{}
	g.stdout = dynamicmultiwriter.New()
	g.stderr = dynamicmultiwriter.New()
	g.stdboth = dynamicmultiwriter.New()

	g.stdout.Add(g.stdboth)
	g.stderr.Add(g.stdboth)

	g.Output(true)

	g.cmd = exec.Command("/usr/bin/env", "sh", "-c", cmd)
	g.cmd.Stdout = g.stdout
	g.cmd.Stderr = g.stderr

	return g
}

// Output ...
func (g *Gommando) Output(enabled bool) {
	if enabled {
		g.stdout.Add(os.Stdout)
		g.stderr.Add(os.Stderr)
	} else {
		g.stdout.Remove(os.Stdout)
		g.stderr.Remove(os.Stderr)
	}
}

// Stdout ...
func (g *Gommando) Stdout() *chain.Chain {
	return chain.New(g.stdout, nil, nil)
}

// Stderr ...
func (g *Gommando) Stderr() *chain.Chain {
	return chain.New(g.stderr, nil, nil)
}

// Stdboth ...
func (g *Gommando) Stdboth() *chain.Chain {
	return chain.New(g.stdboth, nil, nil)
}

// Run ...
func (g *Gommando) Run() {
	g.cmd.Run()
}
