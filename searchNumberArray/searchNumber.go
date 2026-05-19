package main

import "fmt"

func searchNunmber(arr []int) string {
	for count := 0; count < len(arr); count++{
		if arr[count] == 7 {
			return "Found"
		}
	}
	return "Not Found"
}

func main() {
	numbers := []int{3, 4, 5, 7, 8}
	fmt.Println(searchNunmber(numbers))
}