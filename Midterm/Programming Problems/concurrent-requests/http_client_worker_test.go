package httpClientWorker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestFilterNumbers(t *testing.T) {

	want := []int{}
	got := FilterNumbers(want)
	sort.Ints(got)
	if reflect.DeepEqual(got, want) {
		t.Fatalf(`Got %v, want %v.`, got, want)
	}

	want = []int{1, 22, 33, 222}
	got = FilterNumbers([]int{1, 17, 22, 30, 33, 222})
	sort.Ints(got)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`Got %v, want %v.`, got, want)
	}

	want = []int{64, 91}
	got = FilterNumbers([]int{10, 20, 64, 70, 91})
	sort.Ints(got)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`Got %v, want %v.`, got, want)
	}

	want = []int{22, 64, 91}
	got = FilterNumbers([]int{17, 22, 30, 64, 91})
	sort.Ints(got)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`Got %v, want %v.`, got, want)
	}
}
func TestGetNumbers(t *testing.T) {
	want := []int{17, 22, 30, 64, 91}
	router := http.NewServeMux()
	router.Handle("/api/numbers", mockGetNumbers(want))

	mockServer := httptest.NewServer(router)
	worker := NewWorker(mockServer.URL)
	got := worker.GetNumbers()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`Got %v, want %v.`, got, want)
	}
}
func TestProcessAndSubmit(t *testing.T) {
	router := http.NewServeMux()
	data := []int{17, 22, 30, 64, 91}
	resultStream := make(chan int, 10)
	router.Handle("/api/numbers", mockGetNumbers(data))
	router.Handle("/api/numbers/special", mockPostNumber(resultStream))
	mockServer := httptest.NewServer(router)
	worker := NewWorker(mockServer.URL)
	worker.ProcessAndSubmit()
	close(resultStream)
	var got []int
	for v := range resultStream {
		got = append(got, v)
	}
	want := []int{22, 64, 91}
	sort.Ints(got)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`Got %v, want %v.`, got, want)
	}
}
func mockGetNumbers(numbers []int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(numbers)
	}
}

func mockPostNumber(dataStream chan int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, _ := io.ReadAll(r.Body)
		spetialNumber, _ := strconv.Atoi(string(data))
		select {
		case dataStream <- spetialNumber:
		default:
			fmt.Println("Chanel is closed or full")
		}

	}
}
