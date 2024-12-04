package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alesr/adventofcode/pkg/fileloader"
)

type processFunc func([]int) bool

func solve(inputPath string, processor processFunc) (int64, error) {
	lineCh, errCh := fileloader.LoadLines(inputPath)

	select {
	case loadErr := <-errCh:
		if loadErr != nil {
			return 0, fmt.Errorf("failed to load file: %w", loadErr)
		}
	default:
	}

	var safeReports int64

	for line := range lineCh {
		lvls := make([]int, 0)

		for _, field := range strings.Fields(string(line)) {
			lvl, err := strconv.Atoi(field)
			if err != nil {
				return 0, fmt.Errorf("failed to parse number %q: %w", field, err)
			}
			lvls = append(lvls, lvl)
		}

		if processor(lvls) {
			safeReports++
		}

		select {
		case readErr := <-errCh:
			if readErr != nil {
				return 0, fmt.Errorf("error reading file: %w", readErr)
			}
		default:
		}
	}

	if finalErr := <-errCh; finalErr != nil {
		return 0, fmt.Errorf("error after reading file: %w", finalErr)
	}
	return safeReports, nil
}

func checkSequence(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	var increasing, decreasing bool

	firstDiff := nums[1] - nums[0]
	if firstDiff == 0 || abs(firstDiff) > 3 {
		return false
	}

	increasing = firstDiff > 0
	decreasing = firstDiff < 0

	for i := 2; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]

		if diff == 0 || abs(diff) > 3 {
			return false
		}

		if diff > 0 {
			if decreasing {
				return false
			}
			increasing = true
		} else {
			if increasing {
				return false
			}
			decreasing = true
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
