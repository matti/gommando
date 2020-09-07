package main

import (
	"time"

	"github.com/matti/gommando"
)

func main() {
	g := gommando.New("echo out")
	g.Output(false)

	g.Stdout().Once(func(haystack string) bool {
		return true
	}).Then(func(s string) {
		println("found ", s)
	})

	g.Run()
	// TODO:
	time.Sleep(time.Second * 1)
}
