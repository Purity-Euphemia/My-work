package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40}

	for count := 0; count < len(nums); count++ {
		fmt.Println(nums[count])
	}
}
