package main

import "fmt"

func removeDuplicate(nums []int) int {
	value := 0
	for count := 0; count < len(nums); count++{
		for counter := count + 1; counter < len(nums); counter++ {
			if nums[count] == nums[counter]{
				value++
			}
		}
	}
	return value
}
func main() {
	nums := []int{1, 2, 2, 3, 3, 3}
	fmt.Println(removeDuplicate(nums))
	
}
