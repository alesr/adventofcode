package day2

import (
	"fmt"
	"os"
)

func Part2(inputPath string) int64 {
	res, err := solve(inputPath, checkSequenceWithDampener)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to solve part 2: %v\n", err)
		return 0
	}
	return res
}

func checkSequenceWithDampener(nums []int) bool {
	if checkSequence(nums) {
		return true
	}

	// Try removing each number one at a time
	for i := 0; i < len(nums); i++ {
		// Create new slice without current number
		dampened := make([]int, 0, len(nums)-1)
		dampened = append(dampened, nums[:i]...)
		dampened = append(dampened, nums[i+1:]...)

		if checkSequence(dampened) {
			return true
		}
	}
	return false
}
