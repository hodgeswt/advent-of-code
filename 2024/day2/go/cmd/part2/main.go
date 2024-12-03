package main

import (
	"fmt"

	"github.com/hodgeswt/advent-of-code/2024/day2/go/internal/util"
)

func main() {
	levels, err := util.LoadData(false)

	if err != nil {
		panic(err)
	}

	count := 0
	for _, level := range levels {
		if util.MakeValid(level) {
			count++
		}
	}

	fmt.Printf("Safe levels: %d\n", count)
}
