package main

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

func main() {
	format := "Jan-02-2006"
	dates := []string{"Sep-14-2008", "Dec-03-2021", "Mar-18-2022"}

	fmt.Println(sortDates(format, dates...))

}

func sortDates(format string, dates ...string) ([]time.Time, error) {
	sortedDates := make([]time.Time, len(dates))
	for idx, val := range dates {
		currDate, err := time.Parse(format, val)
		if err != nil {
			return nil, errors.New("format-date error")
		} else {
			sortedDates[idx] = currDate
		}
	}

	sort.Slice(sortedDates, func(i, j int) bool {
		return sortedDates[i].Before(sortedDates[j])

	})
	return sortedDates, nil
}
