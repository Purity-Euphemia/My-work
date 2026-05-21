package main

import (
	"fmt"
	"strings"
)

func add(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot add zero")
	}
	return a + b, nil
}

func splitName(fullName string) (string, string) {
	parts := strings.Split(fullName, " ")

	if len(parts) < 2 {
		return fullName, ""
	}

	return parts[0], parts[1]
}

func getUserDetails(id int) (name string, age int, err error) {
	return "John Doe", 30, nil
}