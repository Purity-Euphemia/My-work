package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)
func main() {
	rootDir := "exampleDir"
	nestedDir := filepath.Join(rootDir, "nestedDir")
	if err := os.MkdirAll(nestedDir, 0755);err != nil {
		log.Fatal("Error creating directions:", err)
	}
	fmt.Println("Directories created:", rootDir, "and", nestedDir)
	filePath := filepath.Join(nestedDir, "file.txt")
	fileContent := []byte("This is a sample file")
	if err := os.WriteFile(filePath, fileContent, 0644);
	err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	fmt.Println("File created:", filePath)

	absRootDir, err := filepath.Abs(rootDir)
	if err != nil {
		log.Fatalf("File created:", filePath)
	}


}
