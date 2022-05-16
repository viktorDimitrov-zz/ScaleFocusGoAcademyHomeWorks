package story

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type StoryService struct {
	urlBase string
}

type Story struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Score int    `json:"score"`
}

func NewStoryService(url string) *StoryService {
	return &StoryService{urlBase: url}
}

func (ss *StoryService) getTopStoriesIds(maxCount int) []int {
	req, err := http.NewRequest("GET", ss.urlBase+"/v0/topstories.json", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var ids []int
	json.NewDecoder(resp.Body).Decode(&ids)
	//fmt.Println(ids)
	//defence
	return ids[:maxCount]
}

func (ss *StoryService) GetStories(maxCount int) []Story {
	ids := ss.getTopStoriesIds(maxCount)
	dataChan := make(chan Story, len(ids))
	wg := sync.WaitGroup{}

	for _, id := range ids {
		wg.Add(1)
		go func(id int) {
			dataChan <- ss.GetStoryById(id)
			defer wg.Done()
		}(id)
	}
	wg.Wait()
	close(dataChan)

	result := make([]Story, 0, len(ids))

	for v := range dataChan {
		result = append(result, v)
	}

	return result
}

func (ss *StoryService) GetStoryById(id int) Story {
	// 	"https://hacker-news.firebaseio.com/v0/item/30887445.json?print=pretty"
	url := fmt.Sprintf("%s/v0/item/%d.json?print=pretty", ss.urlBase, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var st Story
	json.NewDecoder(resp.Body).Decode(&st)
	//fmt.Println(st)
	return st
}
