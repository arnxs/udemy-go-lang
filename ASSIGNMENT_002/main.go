package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("reading this file...", os.Args[1])
		filename := os.Args[1]
		// bs, err := ioutil.ReadFile(filename)
		r, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			// printByteSlice(bs)
			// fr := fileReader{contents: bs}
			io.Copy(os.Stdout, r)
		}

	} else {
		fmt.Println("Not enough arguments.  Please pass filename to read")
		os.Exit(1)
	}
}
