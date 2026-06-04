package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Contains example")
	fmt.Println(strings.Contains("Hello, Go!", "Go"))

	fmt.Println("\nCount examples")
	fmt.Println(strings.Count("Hello,Go!", "o"))
}
