package main

import "fmt"

func main() {
	data := map[string]string{
		"foo":   "bar",
		"hello": "world",
	}

	fmt.Println(data)

	otherData := map[string]int{
		"foo":   123,
		"hello": 555,
	}

	fmt.Println(otherData)
}
