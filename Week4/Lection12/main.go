package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	out := generateThrottled("foo", 2, 3*time.Second)
	for f := range out {
		log.Println(f)
	}
}

func generateThrottled(data string, bufferLimit int, clearInterval time.Duration) <-chan string {

	channel := make(chan string, bufferLimit-1)
	go func() {
		timeoutCh := time.After(clearInterval)
		for {
			select {
			case channel <- data: // Put data in the channel unless it is full
			default:
				fmt.Printf("%d miliseconds \n",clearInterval)
				<-timeoutCh
				timeoutCh = time.After(clearInterval)
			}
		}

	}()

	return channel
}
