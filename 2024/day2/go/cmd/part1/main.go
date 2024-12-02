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
        safe := isValid(level)

        if safe {
            count++
        }
    }

    fmt.Printf("%d levels are safe\n", count)
}

func isValid(level []int) bool {

	last := 0
	first := true

	isSet := false
	ascending := false

	for i := 0; i < len(level); i++ {
		if first {
			last = level[i]
			first = false
			continue
		}

		x := level[i]
		if !isSet {
			if last < x {
				ascending = true
				isSet = true
			} else if last > x {
				ascending = false
				isSet = true
			}
		}

		diff := x - last
		if ascending {
			if diff == 1 || diff == 2 || diff == 3 {
				last = x
			} else {
				return false
			}
		} else {
		    if diff == -1 || diff == -2 || diff == -3 {
                last = x
            } else {
                return false
            }
        }
	}

    return true
}
