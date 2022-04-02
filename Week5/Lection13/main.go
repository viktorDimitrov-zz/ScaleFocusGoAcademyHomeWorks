package main

import (
	"context"
	"fmt"
	"time"
)

type BufferedContext struct {
	context.Context
	buffer chan string
	context.CancelFunc
}

func NewBufferedContext(timeout time.Duration, bufferSize int) *BufferedContext {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	buff := make(chan string, bufferSize)
	newwBufferCTX := &BufferedContext{Context: ctx, buffer: buff, CancelFunc: cancel}
	return newwBufferCTX
}

func (bc *BufferedContext) Done() <-chan struct{} {
	fmt.Println("done")
	if len(bc.buffer) == cap(bc.buffer) {
		fmt.Println("buffer limit reached")
		bc.CancelFunc()
	}
	return bc.Context.Done()
}

func (bc *BufferedContext) Run(fn func(context.Context, chan string)) {
	fn(bc, bc.buffer)
}

func main() {

	ctx := NewBufferedContext(time.Second, 10)
	ctx.Run(func(ctx context.Context, buffer chan string) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("we are done, time out is reached")
				return
			case buffer <- "bar":
				time.Sleep(time.Millisecond * 200) // try different values here fmt.Println("bar")
				fmt.Println("bar")
			}
		}
	})

}
