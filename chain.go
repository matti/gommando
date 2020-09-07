package gommando

import (
	"io"

	"github.com/matti/dynamicmultiwriter"
)

type Chain struct {
	then  func(s string)
	once  func()
	every func()

	stdout *dynamicmultiwriter.DynamicMultiWriter
	prev   *Chain
	next   *Chain
}

func (c *Chain) Once(needleFn func(haystack string) bool) *Chain {
	c.next = &Chain{
		stdout: c.stdout,
		prev:   c,
	}

	c.once = func() {
		r, w := io.Pipe()
		c.stdout.Add(w)

		defer w.Close()
		b := make([]byte, 4<<20)
		for {
			r.Read(b)
			if needleFn(string(b)) {
				c.next.Fire(string(b))
				c.stdout.Remove(w)
				return
			}
		}
	}

	// start this if we are the first in chain
	if c.prev == nil {
		go func() {
			c.once()
		}()
	}

	return c.next
}

func (c *Chain) Every(needleFn func(haystack string) bool) *Chain {
	c.next = &Chain{
		stdout: c.stdout,
		prev:   c,
	}

	c.every = func() {
		r, w := io.Pipe()
		c.stdout.Add(w)

		defer w.Close()
		b := make([]byte, 4<<20)
		for {
			r.Read(b)

			if needleFn(string(b)) {
				go c.next.Fire(string(b))
			}
		}
	}

	// start this if we are the first in chain
	if c.prev == nil {
		go func() {
			c.every()
		}()
	}

	return c.next
}

func (c *Chain) Fire(s string) {
	if c.then != nil {
		c.then(s)
		c.next.Fire(s)
	} else if c.once != nil {
		c.once()
	} else if c.every != nil {
		c.every()
	}
}

func (c *Chain) Then(fn func(s string)) *Chain {
	c.then = fn
	c.next = &Chain{
		stdout: c.stdout,
		prev:   c,
	}

	return c.next
}
