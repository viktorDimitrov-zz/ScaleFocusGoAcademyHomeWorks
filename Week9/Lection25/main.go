package main

import (
	"encoding/json"
	//"fmt"
	"main/story"
	"net/http"
)

func HandleTopStories() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		//fetch top 10
		ss := story.NewStoryService("https://hacker-news.firebaseio.com")
		ids := ss.GetTopStoriesIds()

		stList := ss.GetStoriesByIds(ids)
		json.NewEncoder(w).Encode(stList)
		//return
		//fmt.Println(stList)
		//fetch scores for top ten

		//caling item
	}
}

func main() {
	router := http.NewServeMux()
	router.Handle("/api/top", HandleTopStories()) //return top 10
	//router.Handle("top", nil)                    //view html template
	//log.Println("started")
	http.ListenAndServe(":8000", router)
}
