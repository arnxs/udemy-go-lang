package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	// url := "https://medium.com/@avery_1242/my-experience-with-google-foobar-tips-for-tackling-googles-legendary-coding-challenge-dbc20a054e4e"
	url := "http://google.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// size, err := resp.Body.Read(bs)
	// fmt.Println("response:", size, string(bs))
	// fmt.Println(resp.ContentLength)

	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println("using logWriter")
	return 1, nil
}
