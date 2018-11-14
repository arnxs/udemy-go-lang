package main

import "fmt"

type bot interface {
	getGreeting() string
	sayFarewell() string
}

type spanishBot struct{}
type englishBot struct{}
type germanBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	gb := germanBot{}

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(gb)

	printFarewell(eb)
	printFarewell(sb)
	printFarewell(gb)
	// printGreetingSpanish(sb)
}

func (sb spanishBot) getGreeting() string {
	return "Hola"
}

func (eb englishBot) getGreeting() string {
	return "Hello"
}

func (gb germanBot) getGreeting() string {
	return "German Hello"
}

func (sb spanishBot) sayFarewell() string {
	return "sp bye"
}

func (eb englishBot) sayFarewell() string {
	return "eb bye"
}

func (gb germanBot) sayFarewell() string {
	return "gb bye"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func printFarewell(b bot) {
	fmt.Println(b.sayFarewell())
}

// func printGreetingSpanish(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }
