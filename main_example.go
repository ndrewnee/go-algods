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
	tree := make([]int, n)
	for i := 0; i < n; i++ {
		var number int
		if _, err := fmt.Scan(&number); err != nil {
			log.Fatal(err)
		}
		tree[i] = number
	}

	height := exercises.Height(tree)
	fmt.Println(height)
}
