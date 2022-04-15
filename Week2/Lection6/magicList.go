package main

import "fmt"

type MagicList struct {
	LastItem *Item
	length   int
}

type Item struct {
	Value    int
	PrevItem *Item
}

func main() {

	l := MagicList{}

	add(&l, 10)
	fmt.Printf("add %d - length - %d\n", l.LastItem.Value, l.length)
	add(&l, 22)
	fmt.Printf("add %d - length - %d\n", l.LastItem.Value, l.length)

	add(&l, 44)
	fmt.Printf("add %d - length - %d\n", l.LastItem.Value, l.length)

	add(&l, 120)
	fmt.Printf("add %d - length - %d\n", l.LastItem.Value, l.length)

	add(&l, 32)
	fmt.Printf("add %d - length - %d\n", l.LastItem.Value, l.length)

	add(&l, 44)
	fmt.Printf("add %d - length - %d\n", l.LastItem.Value, l.length)

	add(&l, 14)
	fmt.Printf("add %d - length - %d\n", l.LastItem.Value, l.length)

	add(&l, 22)
	fmt.Printf("add %d - length - %d\n", l.LastItem.Value, l.length)

	add(&l, 45)
	fmt.Printf("add %d - length - %d\n", l.LastItem.Value, l.length)

	add(&l, 45)
	fmt.Printf("add %d - length - %d\n", l.LastItem.Value, l.length)

	fmt.Println(l)
	fmt.Println(toSlice(&l))

}

func add(l *MagicList, value int) {
	if l == nil {
		l.LastItem = &Item{Value: value, PrevItem: nil}
	} else {
		second := l.LastItem
		l.LastItem = &Item{Value: value}
		l.LastItem.PrevItem = second
		l.length++
	}
}

func toSlice(l *MagicList) []int {
	result := []int{}

	if l.LastItem.PrevItem == nil {
		fmt.Println("It is an emty list!")
		return nil
	}
	currrent := l.LastItem
	for currrent.PrevItem != nil {
		//prepend
		result = append([]int{currrent.Value}, result...)
		currrent = currrent.PrevItem
	}
	//prepend
	result = append([]int{currrent.Value}, result...)
	return result
}
