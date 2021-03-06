package main

import (
	"github.com/matti/gommando"
)

func main() {
	g := gommando.New("1>&2 echo err; echo out")
	g.Output(false)

	g.Stderr().Once(func(haystack string) bool {
		return true
	}).Then(func(s string) {
		println("found ", s)
	})

	g.Run()
}
