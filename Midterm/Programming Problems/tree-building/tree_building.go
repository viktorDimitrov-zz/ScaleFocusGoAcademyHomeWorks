package tree

import "fmt"

// Define the Record type
type Record struct {
	ID     int
	Parent int
}

// Define the Node type
type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {

  
	for _, r := range records {
		fmt.Println(r)
	}

	panic("Please implement the Build function")
}
