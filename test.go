package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	file = flag.String("file", "./word.csv", "Word file")
	quer = flag.String("q", "k", "Quary word")
)

func ReadData(fileName string) string {
	data, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	file, _ := io.ReadAll(data)

	spldta := strings.Split(string(file), ",")

	fmt.Println(spldta)
	return fileName

}


func main() {


	

}
