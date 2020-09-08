package main

import (
	"fmt"
	"strings"

	"github.com/matti/gommando"
)

func divider(title string) {
	fmt.Printf("\n\n--[ %s ]--%s\n", title, strings.Repeat("-", 50))
}

func main() {
	divider("arkade usage")
	usage()

	divider("arkade get")

	divider("kubectl")
	get("kubectl")
	divider("terraform")
	get("terraform")
}

func usage() {
	g := gommando.New("arkade")
	g.Output(false)

	g.Stdout().Once(func(haystack string) bool {
		return strings.Index(haystack, "the easy way") > -1
	}).Then(func(s string) {
		println(s)
		println("  ^--- it is the easy way")
	})

	g.Run()
}

func get(tool string) {
	g := gommando.New("arkade get " + tool)
	g.Output(false)

	g.Stdout().Once(func(haystack string) bool {
		return strings.Index(haystack, "Downloading "+tool) > -1
	}).Then(func(s string) {
		println("  Downloading ...")
	}).Once(func(haystack string) bool {
		return strings.Index(haystack, "Tool written to") > -1
	}).Then(func(s string) {
		println("  Download completed")
	})

	g.Run()
}
