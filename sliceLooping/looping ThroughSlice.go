package main

import "fmt"

func main() {
	foods := []string{"Rice", "Beans", "Oil"}

	for count, value := range foods {
		fmt.Println(count, value)
	}
}