package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type UrlData struct {
	URL     string
	Message string
	Error   error
}

func pingURL(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Printf("Got response for %s: %d\n", url, resp.StatusCode)
	return nil
}

func fetchURLs(urls []string, conc int) chan UrlData {
	processQueue := make(chan string, conc)
	chanStr := make(chan UrlData)
	var wg sync.WaitGroup

	//implement concurency
	//nest one more goroutine wrapping this for urls
	go func() {
		wg.Add(len(urls))
		for _, url := range urls {
			processQueue <- url

			go func(u string) {
				err := pingURL(u)
				if err != nil {
					chanStr <- UrlData{Error: err}
				} else {
					<-processQueue
					chanStr <- UrlData{URL: u, Message: "ok", Error: err}
				}

			}(url)
		}
		wg.Wait()
		close(chanStr)
	}()

	return chanStr
}

func main() {
	var concurrency int
	urls := []string{}

	flag.IntVar(&concurrency, "c", 2, "It is the number of concurency")
	flag.Parse()

	urls = flag.Args()
	if len(urls) < 1 {
		fmt.Println("flagprinter [-c]  <list of urls>")
		flag.Usage()
		return
	}

//	fmt.Println(concurrency)
	dataValue:=fetchURLs(urls,concurrency)
	for range dataValue{
		fmt.Println(dataValue)
	}
}
