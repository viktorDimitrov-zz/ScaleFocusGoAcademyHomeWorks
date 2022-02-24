package main

import "fmt"

func main() {

	fmt.Println("go" + "lang")
	fmt.Printf("version %f\n", 1.17)

	fmt.Println("1+1 =", 1+1)
	fmt.Println("float numbers 7.0/3.0 =", 7.0/3.0)
	fmt.Println("int numbers 7/3=", 7/3)

	fmt.Printf("true && false=%v\n", true && false)
	fmt.Printf("true || false=%v\n", true || false)

	fmt.Println(true || false)
	fmt.Println(!true)
}
