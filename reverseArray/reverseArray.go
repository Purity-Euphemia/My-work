package main

import "fmt"

func reverseArray(arr []int) []int {
	reverse := []int{}
	for count := len(arr) - 1; count >= 0; count--{
		reverse = append(reverse, arr[count])
	}
	return reverse
}

func main() {
	numbers := [] int{10, 20, 30, 40}
	fmt.Println(reverseArray(numbers))
}
