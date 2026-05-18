package main

import "fmt"

func main() {
	foods := []string{"rice", "stew", "potato"}
		foods = append(foods, "Garri")
		fmt.Println(foods)
}
