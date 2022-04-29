package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type TopStory struct {
	Title string `json:"title"`
	Score int    `json:"score"`
}

type TopStoriesPayload struct {
	TopStories []TopStory
}

type Scraper struct {
	url    string
	Storys []TopStory
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func TakeTopTenStoriesID() []int {
	req, err := http.NewRequest("GET", "https://hacker-news.firebaseio.com/v0/topstories.json", nil)
	checkErr(err)

	resp, err := http.DefaultClient.Do(req)
	checkErr(err)

	var storiesID []int
	json.NewDecoder(resp.Body).Decode(&storiesID)
	storiesID = storiesID[:10]

	return storiesID
}

func (n *Scraper) GetStory() {
	req, err := http.NewRequest("GET", n.url, nil)
	checkErr(err)

	topTen := TakeTopTenStoriesID()

	for _, id := range topTen {
		req.URL.Path = "/v0/item/" + strconv.Itoa(id) + ".json"

		resp, err := http.DefaultClient.Do(req)
		checkErr(err)

		var topOne TopStory
		json.NewDecoder(resp.Body).Decode(&topOne)
		n.Storys = append(n.Storys, topOne)
	}
}

func topStoriesHandler(w http.ResponseWriter, r *http.Request) {
	scraper := Scraper{url: "https://hacker-news.firebaseio.com"}
	scraper.GetStory()
	topPayload := TopStoriesPayload{TopStories: scraper.Storys}
	json.NewEncoder(w).Encode(topPayload)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/top", topStoriesHandler)
	http.ListenAndServe("9000", mux)

}
