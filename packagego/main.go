package main

import "fmt"

const key = 71

func Encrpty(key byte, planeText []byte) []byte {
	result := []byte{}
	for _, c := range planeText {
		result = append(result, c^key)
	}
	return result
}
func Decrpty(key byte, encrptedText []byte) []byte {
	result := []byte{}
	for _, c := range encrptedText {
		result = append(result, c^key)
	}
	return result
}
func sendToAuba(mesage []byte) {
	fmt.Printf("Message : %s:", mesage)
	m2 := Decrpty(key, mesage)
	fmt.Printf("\n Sır ilen gelen mesage : %s", m2)
}

func main() {
	message := "Kazananlar hiç hata yapmayanlar değil, asla vazgeçmeyenlerdir.!"

	m1 := Encrpty(byte(key), []byte(message))
	sendToAuba(m1)
}
