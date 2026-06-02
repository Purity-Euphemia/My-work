package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()
	fmt.Println("starting the program")

	panic("Something went wrong")
	
	fmt.Println("This will not be printed")

}
