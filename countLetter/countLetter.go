package main

import "fmt"

func countAlphe(word string) int {
	counter := 0

	for count := 0; count < len(word); count++{
		if (word[count] >= 'a' && word[count] <= 'z' || word[count] >= 'A' && word[count] <= 'Z') {
			counter++
		}
	}
	return counter
}

func main() {
	value := "za4r75"
	fmt.Println(countAlphe(value))
}


