package main

import "fmt"

func main() {
	foods := []string{"Rice", "Beans", "Garri"}

	foods[1] = "yam"

	fmt.Println(foods)
}
