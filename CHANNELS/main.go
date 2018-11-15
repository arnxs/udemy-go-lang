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

	c := make(chan string)

	for _, link := range links {
		go checkLink(c, link)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c) // this is blocking
	}

}

func checkLink(c chan string, link string) {
	_, err := http.Get(link) // this is blocking call!!! ==> make a goroutine
	if err != nil {
		fmt.Println(link, "might be down!")
	} else {
		fmt.Println(link, "is up!")
	}
	c <- "done"
}
