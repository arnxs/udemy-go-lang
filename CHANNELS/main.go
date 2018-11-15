package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.com",
		"http://amazon.com",
	}

	for _, link := range links {
		// checkLink(link)
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link) // this is blocking call!!! ==> make a goroutine
	if err != nil {
		fmt.Println(link, "might be down!")
	} else {
		fmt.Println(link, "is up!")
	}
}
