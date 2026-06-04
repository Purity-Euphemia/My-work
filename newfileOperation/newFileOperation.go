package main

import (
	"fmt"
	"os"
)
func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	newFile, err := os.Create("newfile.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer newFile.Close()
	fmt.Println("File opened and created!")
}
