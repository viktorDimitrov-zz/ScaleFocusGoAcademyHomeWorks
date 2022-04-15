package tree

import (
	"fmt"
	"sort"
)

func main() {

	records := []Record{
		{ID: 5, Parent: 2},
		{ID: 3, Parent: 2},
		{ID: 2, Parent: 0},
		{ID: 4, Parent: 1},
		{ID: 1, Parent: 0},
		{ID: 0},
		{ID: 6, Parent: 2},
	}

	fmt.Println(records)

	sort.Slice(records, func(i, j int) bool {
		return records[i].Parent < records[j].Parent
	})
	fmt.Println(records)

}