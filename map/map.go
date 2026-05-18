package main

import "fmt"

func main() {
	m := map[string]int{
		"Alice": 25,
		"Bob": 78,
		"Carol": 22,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}
}
