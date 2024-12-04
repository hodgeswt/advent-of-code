package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hodgeswt/advent-of-code/2024/day3/go/internal/util"
)

func main() {

	testMode := flag.Bool("t", false, "indicate if should run test mode")
	part := flag.Int("p", 1, "indicate which part to run")

	flag.Parse()

	var p func(bool)

	if *part == 1 {
		p = util.RunPart1
	} else if *part == 2 {
		p = util.RunPart2
	} else {
		fmt.Printf("Unexpected -p value %d, expected one of [1, 2]\n", part)
		os.Exit(1)
	}

	p(*testMode)
}
