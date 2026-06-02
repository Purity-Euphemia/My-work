package main

import "fmt"

var greet = func(name string) string {
	return "Hello," + name
}

func processGreeting(f func(string)string, name string) {
	fmt.Println(f(name))
}
func main() {
	message := greet("Spock")
	fmt.Println(message)

	processGreeting(greet, "Ripley")
}
