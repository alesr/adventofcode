package day2

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	t.Run("successfully processes input file", func(t *testing.T) {
		t.Parallel()

		f, err := os.CreateTemp("", "test_input_*.txt")
		require.NoError(t, err)
		defer os.Remove(f.Name())

		testData := "1 2 3\n3 2 1\n1 3 2\n"
		_, err = f.WriteString(testData)
		require.NoError(t, err)
		require.NoError(t, f.Close())

		var processedSequences [][]int
		mockProcessor := func(nums []int) bool {
			processedSequences = append(processedSequences, nums)
			return true
		}

		result, err := solve(f.Name(), mockProcessor)
		require.NoError(t, err)
		assert.Equal(t, int64(3), result)
		assert.Len(t, processedSequences, 3)
	})

	t.Run("invalid number format returns error", func(t *testing.T) {
		t.Parallel()

		f, err := os.CreateTemp("", "test_input_*.txt")
		require.NoError(t, err)
		defer os.Remove(f.Name())

		testData := "1 2 x\n"
		_, err = f.WriteString(testData)
		require.NoError(t, err)
		require.NoError(t, f.Close())

		mockProcessor := func(nums []int) bool {
			return true
		}

		_, err = solve(f.Name(), mockProcessor)
		assert.Error(t, err)
	})

	t.Run("file not found returns error", func(t *testing.T) {
		t.Parallel()

		mockProcessor := func(nums []int) bool {
			return true
		}

		_, err := solve("non_existent_file", mockProcessor)
		assert.Error(t, err)
	})
}

func TestCheckSequence(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "empty sequence is safe",
			nums:     []int{},
			expected: true,
		},
		{
			name:     "single number is safe",
			nums:     []int{1},
			expected: true,
		},
		{
			name:     "increasing sequence within limit is safe",
			nums:     []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "decreasing sequence within limit is safe",
			nums:     []int{3, 2, 1},
			expected: true,
		},
		{
			name:     "sequence with difference > 3 is unsafe",
			nums:     []int{1, 5, 9},
			expected: false,
		},
		{
			name:     "sequence with direction change is unsafe",
			nums:     []int{1, 3, 2, 4},
			expected: false,
		},
		{
			name:     "sequence with equal numbers is unsafe",
			nums:     []int{1, 1, 2},
			expected: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, checkSequence(tc.nums))
		})
	}
}

func TestAbs(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{-1, 1},
		{42, 42},
		{-42, 42},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(string(rune(tc.input)), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, abs(tc.input))
		})
	}
}
