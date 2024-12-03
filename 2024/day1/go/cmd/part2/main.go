package main

import (
	"fmt"
	"os"

	"github.com/hodgeswt/advent-of-code/2024/day1/internal/util"
)

func countMatches(m []*util.Entry, target int64) int64 {
	var c int64 = 0

	for _, v := range m {
		if v.Value == target {
			c = c + 1
		}
	}

	return c
}

func main() {
	x, y, err := util.LoadFile(false)

	if err != nil {
		fmt.Printf("Error loading file: %v\n", err)
		os.Exit(1)
	}

	var s int64 = 0

	for i := 0; i < len(x); i++ {
		c := countMatches(y, x[i].Value)

		s = s + (x[i].Value * c)
	}

	fmt.Printf("Found answer: %d\n", s)
}
