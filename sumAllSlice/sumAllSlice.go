package main

import "fmt"

func main() {
	total := 0
	nums := []int{10, 20, 30, 40}
	for count := 0; count < len(nums); count++ {
		total = total + nums[count]
	}
	fmt.Println(total)
}
