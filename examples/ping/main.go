package main

import (
	"strings"
	"syscall"

	"github.com/matti/gommando"
)

func main() {
	g := gommando.New("ping google.com")
	g.Output(true)
	g.Stdout().Once(func(haystack string) bool {
		return strings.Index(haystack, "google") > 0
	}).Then(func(s string) {
		println("  found 'google', I'll never say this again because I'm after .Once, but I'll trigger the next...")
	}).Then(func(s string) {
		println("  ... .Then and I'll trigger the next .Once which will start to match now:")
	}).Once(func(haystack string) bool {
		return strings.Index(haystack, "ttl") > 0
	}).Then(func(s string) {
		println("  found 'ttl' and I'll never say this again because I'm after .Once")
	})

	g.Stdout().Once(func(haystack string) bool {
		return strings.Index(haystack, "seq=3") > 0
	}).Then(func(s string) {
		println("  3 pings sent, exiting")
		g.Signal(syscall.SIGTERM)
		g.Wait()
		println("  ping exited")
	})

	g.Stdout().Every(func(haystack string) bool {
		return strings.Index(haystack, "seq=") > 0
	}).Then(func(s string) {
		println("  found 'seq=' seen and I continue to say this because I'm after .Every")
	})

	g.Run()
}
