package handlers

import (
	"encoding/json"
	"main/story"
	"net/http"
	"time"
)

type Storage interface {
	GetLastStoryTimeStamp() time.Time
	GetStories() []story.Story
}

func HandleTopStories(storage Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		//fetch top 10
		ss := story.NewStoryService("https://hacker-news.firebaseio.com")
		treshold := time.Now().Add(-time.Hour)
		var stList []story.Story
		if storage.GetLastStoryTimeStamp().Before(treshold) {
			stList = ss.GetStories(10)
		}else{
			stList=storage.GetStories()
		}

		json.NewEncoder(w).Encode(stList)

	}
}
