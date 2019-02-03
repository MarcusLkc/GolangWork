package main

import "fmt"

type bot interface {
	getGreeting() string
	// getBotVersion() float64
	// respondToUser() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// we can omit eb value if we are not making use of the value from receiver
func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for generating a spanish greeting
	return "Hola"
}
