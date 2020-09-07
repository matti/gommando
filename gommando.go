package gommando

import (
	"os"
	"os/exec"

	"github.com/matti/dynamicmultiwriter"
)

type Gommando struct {
	stdout *dynamicmultiwriter.DynamicMultiWriter
	cmd    *exec.Cmd
}

func (g *Gommando) Stdout() *Chain {
	return &Chain{
		stdout: g.stdout,
		prev:   nil,
	}
}

func New(cmd string) *Gommando {
	g := &Gommando{}
	g.stdout = dynamicmultiwriter.New()
	g.stdout.Add(os.Stdout)

	g.cmd = exec.Command("/usr/bin/env", "sh", "-c", cmd)
	g.cmd.Stdout = g.stdout

	return g
}

func (g *Gommando) Run() {
	g.cmd.Run()
}
