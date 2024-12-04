package day1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alesr/adventofcode/pkg/fileloader"
)

type (
	processFunc func([]float64, []float64) int64
	result      struct {
		valA, valB float64
		err        error
	}
)

func solve(inputPath string, processor processFunc) (int64, error) {
	linesCh, errCh := fileloader.LoadLines(inputPath)

	var listA, listB []float64

	for line := range linesCh {
		valA, valB, err := processLine(line)
		if err != nil {
			return 0, fmt.Errorf("error processing line: %w", err)

		}
		listA = append(listA, valA)
		listB = append(listB, valB)
	}

	if err := <-errCh; err != nil {
		return 0, fmt.Errorf("error reading file: %w", err)
	}
	return processor(listA, listB), nil
}

func processLine(line []byte) (float64, float64, error) {
	parts := strings.Fields(string(line))
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid input format: expected 2 numbers, got '%d' parts", len(parts))
	}

	valA, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to convert first number: %w", err)
	}

	valB, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to convert second number: %w", err)
	}
	return valA, valB, nil
}
