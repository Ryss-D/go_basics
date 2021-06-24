package main

import (
	"math/rand"
	"time"
)

func createNumbers() []int {

	//lets make random real random
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	numbers := []int{}
	for i := 0; i < 10; i++ {
		newNumber := r.Int()
		numbers = append(numbers, newNumber)
	}

	return numbers
}

func determineType(n int) string {

	//lest say if its even or odd
	if n%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
