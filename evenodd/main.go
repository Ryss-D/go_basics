package main

import "fmt"

func main() {
	numbers := createNumbers()

	fmt.Println(numbers)

	for i, number := range numbers {
		fmt.Println(i, " ", determineType(number))
	}

}
