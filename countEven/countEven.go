package main

import "fmt"


func countEven(arr []int) int{
	value := 0
	for count := 0; count < len(arr); count++{
		if arr[count] % 2 == 0 {
			value++
		}
	}
	return value
}


func main() {
	number := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(countEven(number))
}