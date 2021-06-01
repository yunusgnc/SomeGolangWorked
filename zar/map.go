package main

import (
	"fmt"
	"math/rand"
	"time"
)

func possibility(x int64) int64 {

	rand.Seed(time.Now().Unix())
	r := rand.Intn(5) + 1
	return int64(r)

}

func main() {
	for i := 0; i < 1000000; i++ {
		if possibility(2) == 6 {
			fmt.Println("yes")
		}
	}
}
