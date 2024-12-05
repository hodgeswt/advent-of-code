package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hodgeswt/advent-of-code/2024/day4/go/internal/util"
)

func main() {
    testMode := flag.Bool("t", false, "indicate if in test mode")
    part := flag.Int("p", 1, "indicate which part to run. valid [1,2]")
    flag.Parse()

    var p func(bool)

    if *part == 1 {
        p = util.RunPart1
    } else if *part == 2 {
        p = util.RunPart2
    } else {
        fmt.Printf("-p must be 1 or 2. Got: %d\n", *part)
        os.Exit(1)
    }

    p(*testMode)
}
