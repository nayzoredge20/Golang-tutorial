package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct {
}

type spanishBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func (b englishBot) getGreeting() string {
	// very custom logic for generating english greeting
	return "Hi There"
}

func (sb spanishBot) getGreeting() string {
	// very custom logic for generating english greeting
	return "Hola"
}

// we can note here than in these two functions we are not using eb,sb respectively so we can also omit them like this
// func(englishBot) getGreeting() string{} // this type of declaration is also fine.

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(sb spanishBot){
// 	fmt.Println(sb.getGreeting())
// }
