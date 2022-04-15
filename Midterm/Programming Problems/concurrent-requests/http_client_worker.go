package httpClientWorker

import (
	"encoding/json"
	"log"
	"net/http"
)

type Worker struct {
	serverURL string
}

// Filter the numbers using the rule
func FilterNumbers(data []int) []int {
	result := []int{}

	if len(data) == 0 {
		return nil
	}

	for _, v := range data {

		if CheckSameNums(v) || CheckSumDivideByTen(v) {
			result = append(result, v)
		}
	}
	return result
}

func CheckSameNums(num int) bool {
	areSame := true
	last := num % 10
	rem := last
	for num != 0 {
		if last != rem {
			areSame = false
		}
		num /= 10
		rem = num % 10
	}
	return areSame
}

func CheckSumDivideByTen(num int) bool {
	sum := 0
	rem := num % 10
	for num != 0 {
		sum += rem
		num /= 10
		rem = num % 10
	}
	return sum%10 == 0
}

func NewWorker(serverUrl string) *Worker {
	return &Worker{serverURL: serverUrl}
}

// Get the data from the server
func (w Worker) GetNumbers() []int {

	url := w.serverURL + "/api/numbers"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	numbers := []int{}
	json.NewDecoder(resp.Body).Decode(&numbers)

	return numbers
}

// Get, Filter and Submit the "special" numbers
func (w Worker) ProcessAndSubmit() {
	//panic("Please implement the ProcessAndSubmit function")
}
