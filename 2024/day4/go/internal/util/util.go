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

	count := 0
	for i, vi := range data {
		for j := range vi {
			count += isXmas(i, j, data)
		}
	}

	fmt.Printf("Found XMAS %d times\n", count)
}

func RunPart2(testMode bool) {
	data, err := LoadData(testMode)

	if err != nil {
		panic(err)
	}

	count := 0
	for i, row := range data {
		for j := range row {
			w := getWindow(i, j, data)
			count += validateWindow(w)
		}
	}

	fmt.Printf("Found X-MAS %d times\n", count)
}

type Window struct {
	Tl rune
	Tr rune
	C  rune
	Bl rune
	Br rune
}

func runeToString(a rune, b rune, c rune) string {
	return fmt.Sprintf("%s%s%s", string(a), string(b), string(c))
}

func validateWindow(w *Window) int {
	if w == nil {
		return 0
	}

	tlDown := runeToString(w.Tl, w.C, w.Br)
	blUp := runeToString(w.Bl, w.C, w.Tr)
	trDown := runeToString(w.Tr, w.C, w.Bl)
	brUp := runeToString(w.Br, w.C, w.Tl)

	count := 0

	if tlDown == "MAS" {
		count++
	}

	if blUp == "MAS" {
		count++
	}

	if trDown == "MAS" {
		count++
	}

	if brUp == "MAS" {
		count++
	}

	if count == 2 {
		return 1
	}

	return 0
}

func getWindow(i int, j int, data [][]rune) *Window {
	if i-1 < 0 || i+1 > len(data)-1 || j-1 < 0 || j+1 > len(data[i])-1 {
		return nil
	}

	return &Window{
		Tl: data[i-1][j-1],
		Tr: data[i-1][j+1],
		C:  data[i][j],
		Bl: data[i+1][j-1],
		Br: data[i+1][j+1],
	}
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
