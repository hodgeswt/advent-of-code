package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/hodgeswt/advent-of-code/2024/day1/internal/util"
)

func findSmallest(m []*util.Entry) (int64, error) {
	var minimum int64 = 0
	first := true
	var found *util.Entry

	for _, v := range m {
		if v.Used {
			continue
		}

		if first {
			minimum = v.Value
			first = false
			found = v
			continue
		}

		if v.Value < minimum {
			minimum = v.Value
			found = v
		}
	}

	if found == nil {
		return -1, errors.New("ErrNotFound")
	}

	found.Used = true
	return found.Value, nil
}

func main() {
	x, y, err := util.LoadFile(false)

	if err != nil {
		fmt.Printf("Error loading file: %v\n", err)
		os.Exit(1)
	}

	var s int64 = 0

	for {
		xMin, err := findSmallest(x)

		if err != nil {
			break
		}

		yMin, err := findSmallest(y)

		if err != nil {
			break
		}

		d := yMin - xMin

		if d < 0 {
			d = d * -1
		}

		s = s + d
	}

	fmt.Printf("Found answer: %d\n", s)
}
