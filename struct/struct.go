package main

import "fmt"

type Person struct {

	Name string
	Age int
}

func main() {

	person := Person{"Alice", 25}
	
	fmt.Println(person.Name)
	fmt.Println(person.Age)
}
