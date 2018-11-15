package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		filename := os.Args[1]
		r, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			io.Copy(os.Stdout, r)
		}
	} else {
		fmt.Println("Not enough arguments.  Please pass filename to read")
		os.Exit(1)
	}
}
