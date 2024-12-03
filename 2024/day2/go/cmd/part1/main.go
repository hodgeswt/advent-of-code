package main

import (
	"fmt"

	"github.com/hodgeswt/advent-of-code/2024/day2/go/internal/util"
)

func main() {
	levels, err := util.LoadData(true)

	if err != nil {
		panic(err)
	}

	count := 0
	for _, level := range levels {
		safe := util.IsValid(level)

		if safe {
			count++
		}
	}

	fmt.Printf("%d levels are safe\n", count)
}
