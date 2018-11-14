package main

import "fmt"

type spanishBot struct{}
type englishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreetingSpanish(sb)
}

func (sb spanishBot) getGreeting() string {
	return "Hola"
}

func (eb englishBot) getGreeting() string {
	return "Hello"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreetingSpanish(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}
