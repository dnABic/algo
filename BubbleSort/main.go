package main

import (
	"fmt"
	"math/rand"
	"time"
)

func intInSlice(number int, list []int) bool {
	for _, element := range list {
		if number == element {
			return true
		}
	}
	return false
}

func randomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func sweep(numbers []int, passes int) bool {
	var tempNum int
	var didSwap bool = false

	for cnt := 0; cnt < (len(numbers) - 1 - passes); cnt++ {
		if numbers[cnt] > numbers[cnt+1] {
			tempNum = numbers[cnt+1]
			numbers[cnt+1] = numbers[cnt]
			numbers[cnt] = tempNum
			didSwap = true
		}
	}

	return didSwap
}

func bubbleSort(numbers []int) {
	for passes := 0; passes < len(numbers); passes++ {
		if !sweep(numbers, passes) {
			break
		}
	}
}

func main() {
	var randNumbers []int

	rand.Seed(time.Now().Unix())

	for len(randNumbers) < 20 {
		randInt := randomInt(1, 100)
		if !intInSlice(randInt, randNumbers) {
			randNumbers = append(randNumbers, randInt)
		}
	}
	fmt.Println("Unsorted:", randNumbers)

	bubbleSort(randNumbers)
	fmt.Println("Sorted:", randNumbers)
}
