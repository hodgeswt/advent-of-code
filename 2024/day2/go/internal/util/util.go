package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func remove(slice []int, s int) []int {
    sliced := []int{}
    for i, v := range slice {
        if i != s {
            sliced = append(sliced, v)
        }
    }

    return sliced
}

func MakeValid(level []int) bool {
	if IsValid(level) {
		return true
	}

	for i := 0; i < len(level); i++ {
		skipped := remove(level, i)
		if IsValid(skipped) {
			return true
		}
	}

	return false
}

func IsValid(level []int) bool {
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

func LoadData(test bool) ([][]int, error) {
	var path string

	if test {
		path = "../day2.test"
	} else {
		path = "../day2.input"
	}

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	s := bufio.NewScanner(file)

	levels := [][]int{}
	for s.Scan() {
		d := strings.Split(s.Text(), " ")

		level := []int{}
		for _, x := range d {
			y, err := strconv.ParseInt(x, 10, 64)

			if err != nil {
				return nil, err
			}

			level = append(level, int(y))
		}

		levels = append(levels, level)
	}

	return levels, nil
}
