package main

import (
	"fmt"
	"math/rand"
)

func cal() float64 {
	sum := 0.0
	count := 0.0

	for sum < 1 {
		r := rand.Float64()
		sum += r
		count++
	}
	return count
}
func sim(x int64) float64 {
	sum := 0.0
	var i int64
	for i = 0; i < x; i++ {
		sum += cal()
	}
	return sum / float64(x)

}
func main() {
	fmt.Println(sim(99999))
}
