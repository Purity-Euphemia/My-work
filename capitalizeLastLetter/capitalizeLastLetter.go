package main

import "fmt"



func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	last := word[len(word)-1] - 32
	return word[:len(word)-1] + string(last)
}
	
func main() {
	value := "hello"
	fmt.Println(capitalize(value))
}
