package main

import (
	"fmt"
	"math"
)

type Square struct {
	a, b float64
}

type Circle struct {
	r float64
}

type Shape interface {
	Area() float64
}

type Shapes []Shape

func main() {
	sliceOfShapes := Shapes{NewCircle(3), NewCircle(5), NewSquare(2, 4),NewSquare(4,10)}
	fmt.Printf("Maximal area is :%.2f", sliceOfShapes.LargestArea())
}

func NewCircle(r float64) *Circle {
	return &Circle{r}
}

func NewSquare(aa, bb float64) *Square {
	return &Square{a: aa, b: bb}
}

func (s Square) Area() float64 {
	return s.a * s.b
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func (s Shapes) LargestArea() float64 {
	maxAreaShape := s[0]
	for _, v := range s {
		//fmt.Println(v.Area())
		if maxAreaShape.Area() < v.Area() {
			maxAreaShape = v
		}
	}
	return maxAreaShape.Area()
}
