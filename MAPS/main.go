package main

import "fmt"

func main() {
	// var data map[string]string
	data := make(map[string]string)
	data = map[string]string{
		"foo":   "bar",
		"hello": "world",
		"green": "apple",
	}
	data["sushi"] = "wasabi"

	fmt.Println(data)
	fmt.Println(data["sushi"])

	printMap(data)

	delete(data, "foo")

	fmt.Println(data)

	otherData := map[string]int{
		"foo":   123,
		"hello": 555,
	}

	fmt.Println(otherData)

}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}
