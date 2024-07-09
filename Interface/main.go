package main

import "fmt"

type bot interface {
	getgreeting() string
}

type englishbot struct{}
type spanishbot struct{}

func main() {

	eb := englishbot{}
	sb := spanishbot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(eb englishbot) {
	fmt.Println(eb.getgreeting())
}

func printGreeting(sb spanishbot) {
	fmt.Println(sb.getgreeting())
}

func (eb englishbot) getgreeting() string {
	return "Hi There!"
}

func (eb spanishbot) getgreeting() string {
	return "Hola!"
}
