package main

import (
	"fmt"
	"net/http"
)

// var errRequestFailed = errors.New("Request Failed.")

type checkresult struct {
	url    string // ok
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan checkresult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}

}

// 함수에서, 이 채널은 Send only 이다! 를 channel type에서 명시 (<- 를 붙여줌)
func hitURL(url string, c chan<- checkresult) {
	// fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- checkresult{url: url, status: status}

}
