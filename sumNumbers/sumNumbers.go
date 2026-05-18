package main

import "fmt"

func main() {
	value := 0
	for count := 1; count <= 100; count++{
		value+=count
	}
	fmt.Println(value)
}
