package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		inputPath string
		expected  int64
	}{
		{"input_test", 11},
		{"input", 1319616},
	}

	for _, tc := range testCases {
		t.Run(tc.inputPath, func(t *testing.T) {
			actual := Part1(tc.inputPath)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		inputPath string
		expected  int64
	}{
		{"input_test", 31},
		{"input", 27267728},
	}

	for _, tc := range testCases {
		t.Run(tc.inputPath, func(t *testing.T) {
			actual := Part2(tc.inputPath)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
