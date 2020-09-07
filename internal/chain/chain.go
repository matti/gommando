package chain

import (
	"io"
	"sync"

	"github.com/matti/dynamicmultiwriter"
)

// Chain ...
type Chain struct {
	then  func(s string)
	once  func()
	every func()

	prev *Chain
	next *Chain

	writer *io.PipeWriter
	reader *io.PipeReader

	Stream *dynamicmultiwriter.DynamicMultiWriter
	wg     sync.WaitGroup
}

// Start ...
func (c *Chain) Start() {
	defer func() {
		c.wg.Done()
	}()

	c.wg.Add(1)

	if c.next == nil {
		return
	}

	go c.next.fire("")
}

// Close ...
func (c *Chain) Close() {
	if c.next != nil {
		c.next.Close()
	}
	c.writer.Close()

	c.wg.Wait()
}

// New ...
func New(stream *dynamicmultiwriter.DynamicMultiWriter, prev *Chain, next *Chain) *Chain {
	r, w := io.Pipe()

	return &Chain{
		Stream: stream,
		prev:   prev,
		next:   next,
		writer: w,
		reader: r,
		wg:     sync.WaitGroup{},
	}
}

// Once ...
func (c *Chain) Once(needleFn func(haystack string) bool) *Chain {
	c.next = New(c.Stream, c, nil)
	c.next.once = func() {
		defer func() {
			c.next.Stream.Remove(c.next.writer)
		}()

		c.next.Stream.Add(c.next.writer)

		for {
			// TODO: stdboth fails unless inside of for, why?
			b := make([]byte, 4<<20)
			_, err := c.next.reader.Read(b)

			if err == io.EOF {
				return
			} else if err != nil {
				panic(err)
			}

			if needleFn(string(b)) {
				if c.next.next != nil {
					c.next.next.fire(string(b))
				}
				return
			}
		}
	}

	return c.next
}

// Every ...
func (c *Chain) Every(needleFn func(haystack string) bool) *Chain {
	c.next = New(c.Stream, c, nil)
	c.next.every = func() {
		wg := sync.WaitGroup{}

		defer func() {
			wg.Wait()
			c.next.Stream.Remove(c.next.writer)
		}()

		c.next.Stream.Add(c.next.writer)

		for {
			b := make([]byte, 4<<20)
			_, err := c.next.reader.Read(b)
			if err == io.EOF {
				return
			} else if err != nil {
				panic(err)
			}

			if needleFn(string(b)) {
				wg.Add(1)
				go func() {
					if c.next.next != nil {
						c.next.next.fire(string(b))
					}
					wg.Done()
				}()
			}
		}
	}

	return c.next
}

// Then ...
func (c *Chain) Then(fn func(s string)) *Chain {
	c.next = New(c.Stream, c, nil)
	c.next.then = fn

	return c.next
}

func (c *Chain) fire(s string) {
	defer func() {
		c.wg.Done()
	}()

	c.wg.Add(1)

	switch c.kind() {
	case "then":
		c.then(s)
		if c.next != nil {
			c.next.fire(s)
		}
	case "once":
		c.once()
	case "every":
		c.every()
	}
}

func (c *Chain) kind() string {
	if c.prev == nil {
		return "start"
	} else if c.then != nil {
		return "then"
	} else if c.once != nil {
		return "once"
	} else if c.every != nil {
		return "every"
	} else {
		panic("What am I?")
	}
}