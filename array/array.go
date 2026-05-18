package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	for i := 1; i <= len(arr); i++ {
		arr = append(arr, arr[i])
	}
	fmt.Println(arr)
}