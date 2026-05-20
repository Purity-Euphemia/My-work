package main

import "fmt"

func countGreater5 (num []int) int{
	value := 0
	for count := 0; count < len(num); count++{
		if num[count] > 5 {
			value++
		}
	}
	return value
}

func main() {
	number := []int{2, 9, 4, 7, 1, 10}
	fmt.Println(countGreater5(number))
}