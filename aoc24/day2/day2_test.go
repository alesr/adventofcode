package day2

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
		{"input_test", 2},
		{"input", 421},
	}

	for _, tc := range testCases {
		t.Run(tc.inputPath, func(t *testing.T) {
			t.Parallel()

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
		{"input_test", 4},
		{"input", 476},
	}

	for _, tc := range testCases {
		t.Run(tc.inputPath, func(t *testing.T) {
			t.Parallel()

			actual := Part2(tc.inputPath)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
