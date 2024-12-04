package day1

import (
	"fmt"
	"os"
)

func Part2(inputPath string) int64 {
	res, err := solve(inputPath, processPart2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to solve part 1: %v\n", err)
		return 0
	}
	return res
}

func processPart2(listA, listB []float64) int64 {
	bFreq := make(map[float64]float64, len(listB))
	for _, b := range listB {
		bFreq[b]++
	}

	var res float64
	for _, a := range listA {
		var total float64

		count := bFreq[a] * a

		if count == 0 {
			continue
		}
		total += count
		res += total
	}
	return int64(res)
}
