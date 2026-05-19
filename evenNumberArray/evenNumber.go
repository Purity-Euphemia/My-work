package main

import "fmt"

func evenNumber(arr [] int) int {
	counter := 0
	for count := 0; count < len(arr); count++{
		if arr[count] % 2 == 0 {
			counter++
		}
	}
	return counter
}
 func main() {
	number := [] int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(evenNumber(number))
 }