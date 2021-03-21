package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for len(numbers) > 0 {

		numbersize := len(numbers)
		random := rand.Intn(numbersize)
		//random number print.
		fmt.Println(numbers[random])
		//Delete previously printed element from list.
		numbers = append(numbers[:random], numbers[random+1:]...)

	}

}
