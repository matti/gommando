package main

import (
	"io"
	"strings"
	"time"

	"github.com/matti/gommando"
)

func main() {
	g := gommando.New("./examples/input/helpers/ask")
	g.Output(true)

	g.Stdout().Once(func(haystack string) bool {
		return strings.Index(haystack, "name:") > 0
	}).Then(func(s string) {
		println("  ^-- now it asks my name")
		time.Sleep(1 * time.Second)
		println("      writing it (robot)")
		io.WriteString(g.Stdin(), "robot\n")
	}).Once(func(haystack string) bool {
		return strings.Index(haystack, "hello robot") >= 0
	}).Then(func(s string) {
		println("  ^-- now it greeted me", s)
	})

	g.Run()

	println("  Process exited with", g.ProcessState().ExitCode())
}
