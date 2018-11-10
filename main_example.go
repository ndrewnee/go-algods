package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/ndrewnee/go-algods/exercises"
)

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		log.Fatal(err)
	}

	queries := make([]string, n)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		queries[i] = scanner.Text()
	}

	exercises.MaxStackHandler(queries)
}
