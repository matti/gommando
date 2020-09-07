package main

import (
	"github.com/matti/gommando"
)

func main() {
	g := gommando.New("./examples/stdboth/helpers/print_stdout_stderr")
	g.Output(false)

	g.Stdboth().Every(func(haystack string) bool {
		return true
	}).Then(func(s string) {
		println("got ", s)
	})

	g.Run()
}
