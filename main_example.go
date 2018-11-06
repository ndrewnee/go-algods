package main

import (
	"fmt"
	"log"

	"github.com/ndrewnee/go-algods/exercises"
)

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		log.Fatal(err)
	}

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		var number int
		if _, err := fmt.Scan(&number); err != nil {
			log.Fatal(err)
		}
		numbers[i] = number
	}

	var windowSize int
	if _, err := fmt.Scan(&windowSize); err != nil {
		log.Fatal(err)
	}

	maxWindows := exercises.FindMaxWindows(numbers, windowSize)
	for _, max := range maxWindows {
		fmt.Print(max, " ")
	}
	fmt.Println()
}
