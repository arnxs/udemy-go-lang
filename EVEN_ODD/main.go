package main

import "fmt"

func main() {
	data := [10]int{}
	for i := range data {
		v := i + 1
		if v%2 == 0 {
			fmt.Println(v, "is even")
		} else {
			fmt.Println(v, "is odd")
		}
	}
}
