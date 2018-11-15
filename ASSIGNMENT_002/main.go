package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println(os.Args[1])
	if len(os.Args) > 1 {
		filename := os.Args[1]
		bs, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			printByteSlice(bs)
			// io.Copy(os.Stdout, printByteSlice(bs))
		}

	} else {
		fmt.Println("Not enough arguments.  Please pass filename to read")
		os.Exit(1)
	}
}

func printByteSlice(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
