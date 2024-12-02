package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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
            y, err  := strconv.ParseInt(x, 10, 64)

            if err != nil {
                return nil, err
            }

            level = append(level, int(y))
        }

        levels = append(levels, level)
    }

    return levels, nil
}
