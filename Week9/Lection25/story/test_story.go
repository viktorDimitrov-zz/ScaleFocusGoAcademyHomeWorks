package story

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func handleTopStories(ids []int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(ids)
	}
}

func handleGetStory(stories []Story) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//try to find the id of querry
		var id int
		var resultStory Story

		for _, v := range stories {
			if v.Id == id {
				resultStory = v
				break
			}
		}
		//return story
		json.NewEncoder(w).Encode(resultStory)
	}
}

func TestTopStoriesIds(t *testing.T) {

	router := http.NewServeMux()
	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	router.Handle("/v0/topstories.json", handleTopStories(ids))
	mockServer := httptest.NewServer(router)

	wont := ids[:10]
	//Act
	ss := NewStoryService(mockServer.URL)
	got := ss.getTopStoriesIds(10)

	//Asssert
	if !reflect.DeepEqual(got, wont) {
		t.Fatalf("Get %v,want %v", got, wont)
	}
}

func TestTopStories(t *testing.T) {
	router := http.NewServeMux()
	ids := []int{10}
	stories := []Story{
		{
			Id:    10,
			Title: "test 10",
			Score: 15,
		},
	}
	router.Handle("/v0/topstories.json", handleTopStories(ids))
	router.Handle("/v0/item", handleGetStory(stories))
	mockServer := httptest.NewServer(router)

	wont := stories
	//Act
	ss := NewStoryService(mockServer.URL)
	got := ss.GetStories(1)

	//Asssert
	if !reflect.DeepEqual(got, wont) {
		t.Fatalf("Get %v,want %v", got, wont)
	}
}
