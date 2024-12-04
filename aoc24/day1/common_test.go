package day1

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

		testData := "1 2\n3 4\n5 6\n"
		_, err = f.WriteString(testData)
		require.NoError(t, err)
		require.NoError(t, f.Close())

		mockProc := func(a, b []float64) int64 {
			assert.Equal(t, []float64{1, 3, 5}, a)
			assert.Equal(t, []float64{2, 4, 6}, b)
			return 42
		}

		result, err := solve(f.Name(), mockProc)
		require.NoError(t, err)
		assert.Equal(t, int64(42), result)
	})

	t.Run("invalid input format return an error", func(t *testing.T) {
		t.Parallel()

		f, err := os.CreateTemp("", "test_input_*.txt")
		require.NoError(t, err)
		defer os.Remove(f.Name())

		testData := "invalid data\n"
		_, err = f.WriteString(testData)
		require.NoError(t, err)
		require.NoError(t, f.Close())

		mockProc := func(a, b []float64) int64 {
			return 42
		}

		_, err = solve(f.Name(), mockProc)
		assert.Error(t, err)
	})

	t.Run("file not found return an error", func(t *testing.T) {
		t.Parallel()

		var processCalled bool
		mockProc := func(a, b []float64) int64 {
			processCalled = true
			return 42
		}

		result, err := solve("non_existent_file", mockProc)
		require.Error(t, err)

		require.False(t, processCalled)
		assert.Equal(t, int64(0), result)
	})
}
