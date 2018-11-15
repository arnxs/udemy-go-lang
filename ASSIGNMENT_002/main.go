package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type fileReader struct {
	contents []byte
}

func main() {
	if len(os.Args) > 1 {
		fmt.Println("reading this file...", os.Args[1])
		filename := os.Args[1]
		bs, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			// printByteSlice(bs)
			// fr := fileReader{contents: bs}
			fr := bytes.NewReader(bs)
			io.Copy(os.Stdout, fr)
		}

	} else {
		fmt.Println("Not enough arguments.  Please pass filename to read")
		os.Exit(1)
	}
}

// func (fr fileReader) Read(p []byte) (n int, err error) {
// 	// fmt.Println(string(fr.contents))
// 	// copy(p, fr.contents)

// 	return len(fr.contents), nil
// 	// return 1, nil
// }
