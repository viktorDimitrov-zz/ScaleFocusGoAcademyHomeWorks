package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	cities, prices := citiesAndPrices()
	fmt.Println(cities)
	fmt.Println(prices)

	/*
		citiesHard:=[]string{"Berlin","Tokio","Sofia","Sofia","Sofia","Sofia","Tokio","Tokio","Tokio","Tokio"}
		pricesHard:=[]int{1,2,3,4,5,6,7,8,9,10}
	*/

	m := countOccurancess(cities, prices)
	fmt.Println(m)
}

func citiesAndPrices() ([]string, []int) {
	rand.Seed(time.Now().UnixMilli())
	cityChoices := []string{"Berlin", "Moscow", "Chicago", "Tokyo", "London"}
	dataPointCount := 100

	// randomly choose cities
	cities := make([]string, dataPointCount)
	for i := range cities {
		cities[i] = cityChoices[rand.Intn(len(cityChoices))]
	}

	prices := make([]int, dataPointCount)
	for i := range prices {
		prices[i] = rand.Intn(100)
	}
	return cities, prices
}

func countOccurancess(s []string, p []int) map[string][]int {
	m := make(map[string][]int)

	for idx, v := range s {
		if m[v] == nil {
			m[v] = []int{p[idx]}
		} else {
			m[v] = append(m[v], p[idx])
		}
	}
	return m
}

/*
func countOccurances(s []string) map[string]int {
	m := make(map[string]int)
	for _, v := range s {
		m[v] = m[v] + 1
	}
	return m
}
*/
