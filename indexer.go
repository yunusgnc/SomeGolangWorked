package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Documant struct {
	ID   int
	Freq string
}

var (
	file = flag.String("file", "./word.csv", "Word file")
	quer = flag.String("q", "k", "Quary word")
)

func main() {
	indexer := map[string][]Documant{}

	data, err := os.Open(*file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data, indexer)

}
