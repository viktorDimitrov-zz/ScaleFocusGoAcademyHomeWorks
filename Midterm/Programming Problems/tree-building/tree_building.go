package tree

import (
	//	"errors"
	"errors"
	"sort"
)

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

	relationMap := make(map[int][]*Node)
	sortedRecords := records
	sort.Slice(records, func(i, j int) bool {
		return sortedRecords[i].ID < sortedRecords[j].ID
	})

	if len(records) == 0 {
		return nil, nil
	}

	root := &Node{}

	if sortedRecords[0].Parent != 0 {
		return nil, errors.New("Root shoud not have parent")
		//panic("Root shoud not have parent")
	}

	if sortedRecords[0].ID != 0 {
		return nil, errors.New("No root in this tree")
	}

	if len(records) == 1 {
		root = &Node{ID: 0}
		return root, nil
	}

	//adding root to map
	relationMap[root.ID] = root.Children

	for i := 1; i < len(sortedRecords); i++ {
		// check for duplicate node
		if _, ok := relationMap[sortedRecords[i].ID]; ok {
			// the key  exists within the map
			return nil, errors.New("duplicate key ")
		}

		currentNode := Node{ID: sortedRecords[i].ID}

		//check not continous
		if currentNode.ID != i {
			return nil, errors.New("Non-continous")
		}
		//adding new element
		relationMap[sortedRecords[i].ID] = []*Node{}
		//update parent node
		relationMap[sortedRecords[i].Parent] = append(relationMap[sortedRecords[i].Parent], &currentNode)

	}

	return &Node{ID: 0, Children: relationMap[0]}, nil

}
