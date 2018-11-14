package main

import "fmt"

func main() {
	// var data map[string]string
	data := make(map[string]string)
	data = map[string]string{
		"foo":   "bar",
		"hello": "world",
	}
	data["sushi"] = "wasabi"

	fmt.Println(data)
	fmt.Println(data["sushi"])

	delete(data, "foo")

	fmt.Println(data)

	otherData := map[string]int{
		"foo":   123,
		"hello": 555,
	}

	fmt.Println(otherData)
}
