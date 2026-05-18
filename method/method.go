package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}
func main() {
	p := Person{Name: "Alice", Age: 25}
	fmt.Println(p.Greet())
}

