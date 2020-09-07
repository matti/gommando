package chain

import (
	"io"

	"github.com/matti/dynamicmultiwriter"
)

// Chain ...
type Chain struct {
	then  func(s string)
	once  func()
	every func()

	stream *dynamicmultiwriter.DynamicMultiWriter
	prev   *Chain
	next   *Chain
}

// New ...
func New(stream *dynamicmultiwriter.DynamicMultiWriter, prev *Chain, next *Chain) *Chain {
	return &Chain{
		stream: stream,
		prev:   prev,
		next:   next,
	}
}

// Once ...
func (c *Chain) Once(needleFn func(haystack string) bool) *Chain {
	c.next = New(c.stream, c, nil)

	c.once = func() {
		r, w := io.Pipe()
		defer w.Close()

		c.stream.Add(w)

		b := make([]byte, 4<<20)
		for {
			r.Read(b)
			if needleFn(string(b)) {
				c.next.fire(string(b))
				c.stream.Remove(w)
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

// Every ...
func (c *Chain) Every(needleFn func(haystack string) bool) *Chain {
	c.next = New(c.stream, c, nil)

	c.every = func() {
		r, w := io.Pipe()
		defer w.Close()

		c.stream.Add(w)

		b := make([]byte, 4<<20)
		for {
			r.Read(b)

			if needleFn(string(b)) {
				go c.next.fire(string(b))
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

// Then ...
func (c *Chain) Then(fn func(s string)) *Chain {
	c.then = fn
	c.next = New(c.stream, c, nil)

	return c.next
}

func (c *Chain) fire(s string) {
	if c.then != nil {
		c.then(s)
		c.next.fire(s)
	} else if c.once != nil {
		c.once()
	} else if c.every != nil {
		c.every()
	}
}
