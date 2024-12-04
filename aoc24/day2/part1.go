package day2

import (
	"fmt"
	"os"
)

func Part1(inputPath string) int64 {
	res, err := solve(inputPath, checkSequence)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to solve part 1: %v\n", err)
		return 0
	}
	return res
}
