package day1

import (
	"fmt"
	"math"
	"os"
	"slices"
)

func Part1(inputPath string) int64 {
	res, err := solve(inputPath, processPart1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to solve part 1: %v\n", err)
		return 0
	}
	return res
}

func processPart1(listA, listB []float64) int64 {
	slices.Sort(listA)
	slices.Sort(listB)

	var offset float64
	for i := 0; i < len(listA); i++ {
		offset += math.Abs(listA[i] - listB[i])
	}
	return int64(offset)
}
