package main

import (
	"log"
	"os"
)

func main() {
	originalPath := "test108.txt"
	newPath := "test.txt"
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
