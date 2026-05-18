package main

import "fmt"

func main() {
	cars := []string{"Toyota", "BMW", "Honda"}

	for count := 0; count < len(cars); count++ {
		fmt.Println(cars[count])
	}
}
