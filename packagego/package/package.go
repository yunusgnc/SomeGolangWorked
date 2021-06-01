package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"runtime"
)

func çeviri(key byte, text []byte) []byte {
	rslt := []byte{}

	for _, c := range text {
		rslt = append(rslt, c^key)
	}

	return rslt
}

const key = 71

func sendToreis(message []byte) {

	m1 := çeviri(key, message)

	fmt.Printf(" \n m1 : %s", m1)

}

func Rules() {
	var control = regexp.MustCompile(`^[a-z]+[0-9]`)

	fmt.Println(control.MatchString("Auba15"))
	fmt.Println(control.MatchString("auba1"))

}

func main() {
	f, err := ioutil.ReadFile("./text.csv")

	if err != nil {
		log.Fatal(err, f)
	}

	Rules()

	if r := runtime.GOOS; r == "windows" {
		fmt.Println("Windows için yönetici olarak çalıştırın.")
	} else if r == "linux" {
		fmt.Println("Linux için sudo komutu ile çalıştırın.")
	} else {
		fmt.Println("Geçersiz işletim sistemi!")
	}

}
