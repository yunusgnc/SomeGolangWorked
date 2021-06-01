package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	candles := []int{}
	age := 5
	for i := 0; i < age; i++ {
		candles = append(candles, rand.Intn(11))
	}

	fmt.Println(candles)

	a := Calculater(candles)

	fmt.Println(a)

}

func Calculater(array []int) int {
	var max int = 0
	count := 0
	for _, value := range array {
		if max < value {
			max = value
		}
	}

	for i := 0; i < len(array); i++ {
		if array[i] == max {
			count++
		}
	}

	return count
}
