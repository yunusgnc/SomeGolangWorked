package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	filepath := "test.txt"
	if Hvefli(filepath) {
		fmt.Println("BuLunyor")
	} else {
		fmt.Println("Bulunmuyor..1")
	}

}

func Hvefli(filename string) bool {

	f, err := os.Stat(filename)

	if os.IsNotExist(err) {
		log.Fatal("ERROR!")
		return false
	}

	return !f.IsDir()

}
