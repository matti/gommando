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

	stream *dynamicmultiwriter.DynamicMultiWriter
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
		stream: stream,
		prev:   prev,
		next:   next,
		writer: w,
		reader: r,
		wg:     sync.WaitGroup{},
	}
}

// Once ...
func (c *Chain) Once(needleFn func(haystack string) bool) *Chain {
	c.next = New(c.stream, c, nil)
	c.next.once = func() {
		c.next.stream.Add(c.next.writer)

		b := make([]byte, 4<<20)
		for {
			_, err := c.next.reader.Read(b)

			if err == io.EOF {
				return
			} else if err != nil {
				panic(err)
			}

			if needleFn(string(b)) {
				// no need to get any more data
				c.next.stream.Remove(c.next.writer)
				c.next.reader.Close()
				c.next.writer.Close()

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
	c.next = New(c.stream, c, nil)
	c.next.every = func() {
		wg := sync.WaitGroup{}

		defer func() {
			c.next.stream.Remove(c.next.writer)
			c.next.reader.Close()
			// In Every already closed, but just want to be explicit
			c.next.writer.Close()

			wg.Wait()
		}()

		c.next.stream.Add(c.next.writer)
		b := make([]byte, 4<<20)
		for {
			_, err := c.next.reader.Read(b)
			if err == io.EOF {
				return
			} else if err != nil {
				panic(err)
			}

			if needleFn(string(b)) {
				if c.next.next != nil {
					wg.Add(1)
					go func(b []byte) {
						c.next.next.fire(string(b))
						wg.Done()
					}(b)
				}
			}
		}
	}

	return c.next
}

// Then ...
func (c *Chain) Then(fn func(s string)) *Chain {
	c.next = New(c.stream, c, nil)
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
