package main

import (
	"strings"
)

type card struct {
	suit  string
	value string
}

func (c card) toString() string {
	return c.value + " of " + c.suit
}

func parseStringToCard(str string) card {
	a := strings.Split(str, " of ")
	return card{suit: a[1], value: a[0]}
}
