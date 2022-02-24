package main

import "fmt"

func main() {

	var days, month, year int
	month = 2
	year = 2004
	days, _ = daysInMonth(month, year)
	fmt.Printf(" %s,%d has %d days\n", convertMonth(month), year, days)

	month = 2
	year = 2000
	days, _ = daysInMonth(month, year)
	fmt.Printf(" %s,%d has %d days\n", convertMonth(month), year, days)

	month = 7
	year = 2004
	days, _ = daysInMonth(month, year)
	fmt.Printf(" %s,%d has %d days\n", convertMonth(month), year, days)

}

func daysInMonth(month int, year int) (int, bool) {
	var daysInMth int

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		daysInMth = 31
	case 4, 6, 9, 11:
		daysInMth = 30
	case 2:
		if leapYear(year) {
			daysInMth = 29
		} else {
			daysInMth = 28
		}
	default:
		fmt.Println("Are you sure the month?")
	}
	return daysInMth, leapYear(year)
}

func leapYear(y int) bool {
	if y%4 == 0 && y%100 == 0 && y%400 == 0 {
		return true
	}
	return false
}
func convertMonth(m int) string {
	var month string = ""
	switch m {
	case 1:
		month = "January"
	case 2:
		month = "February"
	case 3:
		month = "March"
	case 4:
		month = "April"
	case 5:
		month = "May"
	case 6:
		month = "June"
	case 7:
		month = "July"
	case 8:
		month = "August"
	case 9:
		month = "September"
	case 10:
		month = "October"
	case 11:
		month = "November"
	case 12:
		month = "December"
	default:
		fmt.Println("No such month")
	}

	return month
}
