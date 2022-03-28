package main

import (
	"fmt"
	"sync"
)

func main() {

	inputs := []int{1, 17, 34, 56, 2, 8, 36, 3, 22, 55}

	evenCh := processEven(inputs)
	fmt.Println("Even numbers:")
	for v := range evenCh {
		fmt.Println(v)
	}

	oddCh := processOdd(inputs)
	fmt.Println("Odd numbers")

	for o := range oddCh {
		fmt.Println(o)
	}
}

func processEven(inputs []int) chan int {
	var wg1 sync.WaitGroup
	chEven := make(chan int)
	go func() {
		wg1.Add(len(inputs))
		for _, v := range inputs {

			go func(val int) {
				defer wg1.Done()
				if val%2 == 0 {
					chEven <- val
				}
			}(v)
		}

		wg1.Wait()
		close(chEven)
	}()

	return chEven
}

func processOdd(inputs []int) chan int {

	var wg sync.WaitGroup
	chOdd := make(chan int)
	go func() {
		wg.Add(len(inputs))
		for _, v := range inputs {

			go func(val int) {
				defer wg.Done()
				if val%2 != 0 {
					chOdd <- val
				}
			}(v)
		}

		wg.Wait()
		close(chOdd)
	}()

	return chOdd
}
