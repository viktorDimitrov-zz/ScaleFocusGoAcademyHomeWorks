package story

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

//	"github.com/stretchr/testify/assert"
)

func handleTopStories(ids []int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(ids)
	}
}
func TestTopStoriesIds(t *testing.T) {
	router := http.NewServeMux()
	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	router.Handle("/v0/topstories.json", handleTopStories(ids))
	mockServer := httptest.NewServer(router)

	wont:=ids[:10]
	//Act
	ss := NewStoryService(mockServer.URL)
	got:=ss.GetTopStoriesIds()
	//Asssert
	if !reflect.DeepEqual(got,wont){
		t.Fatalf("Get %v,want %v",got,wont)
	}
}
