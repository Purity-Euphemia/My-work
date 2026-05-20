package main

import "fmt"

func countdivisble3And5 (arr []int) int {
	value := 0
	for count := 0; count < len(arr); count++{
		if arr[count] % 3 == 0 && arr[count] % 5 == 0 {
			value++
		}
	}
	return value
}

func main() {
	number := []int{3, 5, 15, 10, 30, 7, 9}
	fmt.Println(countdivisble3And5(number))
}
