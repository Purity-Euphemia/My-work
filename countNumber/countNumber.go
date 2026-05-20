package main

import "fmt"

func countNumber(arr []int) int{
	value := 0
	for count := 0; count < len(arr); count++{
		value++
	}
	return value
}

func main() {
	number := []int{4, 7, 2, 9, 1, 6}
	fmt.Println(countNumber(number)) 
}