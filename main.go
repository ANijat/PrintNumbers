package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for len(nums) > 0 {

		numsize := len(nums)
		r := rand.Intn(numsize)
		//random number print.
		fmt.Println(nums[r])
		//Delete previously printed element from list.
		nums = append(nums[:r], nums[r+1:]...)

	}

}
