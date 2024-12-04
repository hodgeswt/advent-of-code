package main

import (
	"flag"

	"github.com/hodgeswt/advent-of-code/2024/day3/go/internal/util"
)

func main() {

	testMode := flag.Bool("t", false, "indicate if should run test mode")

    flag.Parse()

    util.RunPart1(*testMode)
}
