package util

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func RunPart1(testMode bool) {
	data, err := LoadData(testMode)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Found XMAS %d times\n", countXmas(data))
}

func RunPart2(testMode bool) {

}

func countXmas(data [][]rune) int {
	count := 0
	for i, vi := range data {
		for j := range vi {
			count += isXmas(i, j, data)
		}
	}

	return count
}

func getNext(c rune) (rune, bool) {
	switch c {
	case 'X':
		return 'M', false
	case 'M':
		return 'A', false
	case 'A':
		return 'S', false
	case 'S':
		return '.', true
	default:
		panic(errors.New("ErrInvalidRune"))
	}
}

func findXmas(c rune, i int, j int, di int, dj int, data [][]rune) int {
	if i < 0 || j < 0 || i > len(data)-1 || j > len(data[i])-1 {
		return 0
	}

	r := data[i][j]

	if c != r {
		return 0
	}

	n, done := getNext(c)

	if done {
		return 1
	}

	return findXmas(n, i+di, j+dj, di, dj, data)
}

func isXmas(i int, j int, data [][]rune) int {
	r := data[i][j]

	if r != 'X' {
		return 0
	}

	n := 'M'

	xmasLeft := findXmas(n, i, j-1, 0, -1, data)
	xmasRight := findXmas(n, i, j+1, 0, 1, data)
	xmasUp := findXmas(n, i-1, j, -1, 0, data)
	xmasDown := findXmas(n, i+1, j, 1, 0, data)

	xmasDiagLeftUp := findXmas(n, i-1, j-1, -1, -1, data)
	xmasDiagRightUp := findXmas(n, i-1, j+1, -1, 1, data)
	xmasDiagLeftDown := findXmas(n, i+1, j-1, +1, -1, data)
	xmasDiagRightDown := findXmas(n, i+1, j+1, +1, 1, data)

	c := xmasLeft + xmasRight + xmasUp + xmasDown + xmasDiagLeftUp + xmasDiagRightUp + xmasDiagLeftDown + xmasDiagRightDown

	return c
}

func LoadData(testMode bool) ([][]rune, error) {
	var path string

	if testMode {
		path = "../day4.test"
	} else {
		path = "../day4.input"
	}

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	s := bufio.NewScanner(file)

	o := [][]rune{}

	for s.Scan() {
		l := []rune(s.Text())
		o = append(o, l)
	}

	return o, nil
}
