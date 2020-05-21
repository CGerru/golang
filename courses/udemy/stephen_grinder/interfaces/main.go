package main

import "fmt"

/*
Here  is an extraordinary example of interface implementation.
In my opinion theres no much difference with how is made in Java
Well, it might be a slight difference, because I do not see how this
interfaces can be implemented and declared in several files, just as
we do in Java with the interface model (interface declaration and
interface implementation) */
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

// getGreeting for only englishBot type. If the receiver value is not going to be used, it can be omitted
func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "HI there!"
}

// getGreeting for only spanishBot type
func (spanishBot) getGreeting() string {
	// VERY custom logic for generating a spanish greeting
	return "¡Hola Valeria!"
}

// Implementation of the interface declared above
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// Commented code after refactor, we keep this pedagogically
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// Overloading is not permitted in Go (sigh), there cannot be two functions with identical names (without different receivers)
// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }
