package main

import (
	"strings"

	"github.com/matti/gommando"
)

func main() {
	g := gommando.New("ping -c 5 google.com")

	g.Stdout().Once(func(haystack string) bool {
		return strings.Index(haystack, "google") > 0
	}).Then(func(s string) {
		println("found google, I'll never say this again because I'm after .Once, but I'll trigger the next...")
	}).Then(func(s string) {
		println("...Then and I'll trigger the next .Once which start to match now:")
	}).Once(func(haystack string) bool {
		return strings.Index(haystack, "ttl") > 0
	}).Then(func(s string) {
		println("'ttl' seen once and I'll never say this again because I'm after .Once")
	})

	g.Stdout().Every(func(haystack string) bool {
		return strings.Index(haystack, "icmp_seq=") > 0
	}).Then(func(s string) {
		println("'icmp_seq=' seen and I continue to say this because I'm after .Every")
	})

	g.Run()
}
