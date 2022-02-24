package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println("bool value", d)

	var e int
	fmt.Println("default iteger value is ",e)

	f := "apple"
	fmt.Printf("%v is of type %T",f,f)
}
