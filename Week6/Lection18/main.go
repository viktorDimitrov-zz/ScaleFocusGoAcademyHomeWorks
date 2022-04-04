package main

import (
	"fib/hw/prime"
	"fmt"
)

func main() {
	p := prime.GoPrimesAndSleep(20, 10)

	fmt.Println(p)

}
