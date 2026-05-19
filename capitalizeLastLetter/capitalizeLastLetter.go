package main

import "fmt"



func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	first := word[0] - 32
	return string(first) + word[1:]
}

func main() {
	value := "hello"
	fmt.Println(capitalize(value))
}
