package main

import (
	"strings"

	"github.com/matti/gommando"
)

func main() {
	g := gommando.New("ping google.com")

	g.Stdout().Once(func(haystack string) bool {
		return strings.Index(haystack, "google") > 0
	}).Then(func(s string) {
		println("google")
	}).Then(func(s string) {
		println("found")
	}).Once(func(haystack string) bool {
		return strings.Index(haystack, "icmp") > 0
	}).Then(func(s string) {
		println("icmp seen")
	})

	g.Stdout().Every(func(haystack string) bool {
		return strings.Index(haystack, "icmp_seq=") > 0
	}).Then(func(s string) {
		println("icmp_seq=")
	})

	g.Run()
	println("exit")
}
