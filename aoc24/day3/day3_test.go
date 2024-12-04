package day3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		inputPath string
		expected  int64
	}{
		{"input_test", 161},
		{"input", 160672468},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Part1 %s", tc.inputPath), func(t *testing.T) {
			actual := Part1(tc.inputPath)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
