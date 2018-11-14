package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// url := "https://medium.com/@avery_1242/my-experience-with-google-foobar-tips-for-tackling-googles-legendary-coding-challenge-dbc20a054e4e"
	url := "http://google.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// fmt.Println(resp.Status)
	// fmt.Println(resp.StatusCode)
	// bs := []byte{}
	bs := make([]byte, 99999)
	size, err := resp.Body.Read(bs)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	os.Exit(1)
	// }
	fmt.Println("response:", size, string(bs))
	fmt.Println(resp.ContentLength)
}
