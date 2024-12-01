package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Entry struct {
	Value int64
	Used  bool
}

func LoadFile(test bool) ([]*Entry, []*Entry, error) {
	path := "./input.txt"

	if test {
		path = "./input-test.txt"
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	s := bufio.NewScanner(file)

	var x []*Entry
	var y []*Entry

	if test {
		x = make([]*Entry, 6)
		y = make([]*Entry, 6)
	} else {
		x = make([]*Entry, 1000)
		y = make([]*Entry, 1000)
	}

	i := 0
	for s.Scan() {
		d := strings.Split(s.Text(), " ")

		a, err := strconv.ParseInt(d[0], 10, 64)

		if err != nil {
			return nil, nil, err
		}

		x[i] = &Entry{
			Value: a,
			Used:  false,
		}

		b, err := strconv.ParseInt(d[3], 10, 64)

		if err != nil {
			return nil, nil, err
		}

		y[i] = &Entry{
			Value: b,
			Used:  false,
		}

		i = i + 1
	}

	return x, y, nil
}
