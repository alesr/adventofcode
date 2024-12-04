package fileloader

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ExampleLoadLines() {
	file, err := os.CreateTemp("", "testfile")
	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())

	content := "line1\nline2\nline3\n"
	if _, err = file.WriteString(content); err != nil {
		panic(err)
	}

	file.Close()

	linesCh, errCh := LoadLines(file.Name())

	var lines [][]byte
	for line := range linesCh {
		lines = append(lines, line)
	}

	err = <-errCh
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		fmt.Print(string(line))
	}

	// Output:
	// line1
	// line2
	// line3
}

func TestConfiguration(t *testing.T) {
	t.Parallel()

	t.Run("WithChunkSize sets custom chunk size", func(t *testing.T) {
		t.Parallel()

		cfg := Config{}
		WithChunkSize(1024)(&cfg)
		assert.Equal(t, 1024, cfg.chunkSize)
	})

	t.Run("WithChunkSize does not set negative chunk size", func(t *testing.T) {
		t.Parallel()

		cfg := Config{}
		WithChunkSize(-1)(&cfg)
		assert.Equal(t, 0, cfg.chunkSize)
	})

	t.Run("WithLineDelimiter sets custom delimiter", func(t *testing.T) {
		t.Parallel()

		cfg := Config{}
		WithLineDelimiter(',')(&cfg)
		assert.Equal(t, byte(','), cfg.lineDelimiter)
	})
}

func TestLoadLines(t *testing.T) {
	t.Parallel()

	t.Run("LoadLines with default buffer size", func(t *testing.T) {
		t.Parallel()

		file, err := os.CreateTemp("", "testfile")
		require.NoError(t, err)

		defer os.Remove(file.Name())

		content := "line1\nline2\nline3\n"
		_, err = file.WriteString(content)
		require.NoError(t, err)

		file.Close()

		linesCh, errCh := LoadLines(file.Name())

		var lines [][]byte
		for line := range linesCh {
			lines = append(lines, line)
		}

		err = <-errCh
		assert.NoError(t, err)

		expectedLines := [][]byte{
			[]byte("line1\n"),
			[]byte("line2\n"),
			[]byte("line3\n"),
		}
		assert.Equal(t, expectedLines, lines)
	})

	t.Run("LoadLines with custom delimiter", func(t *testing.T) {
		t.Parallel()

		file, err := os.CreateTemp("", "testfile")
		require.NoError(t, err)
		defer os.Remove(file.Name())

		content := "value1,value2,value3,value4"
		_, err = file.WriteString(content)
		require.NoError(t, err)
		file.Close()

		linesCh, errCh := LoadLines(file.Name(), WithLineDelimiter(','))

		var lines [][]byte
		for line := range linesCh {
			lines = append(lines, line)
		}

		err = <-errCh
		assert.NoError(t, err)

		expectedLines := [][]byte{
			[]byte("value1,"),
			[]byte("value2,"),
			[]byte("value3,"),
			[]byte("value4"),
		}
		assert.Equal(t, expectedLines, lines)
	})

	t.Run("LoadLines with non-existent file returns error", func(t *testing.T) {
		t.Parallel()

		linesCh, errCh := LoadLines("non_existent_file.txt")

		var lines [][]byte
		for line := range linesCh {
			lines = append(lines, line)
		}

		err := <-errCh
		assert.Error(t, err)
		assert.Empty(t, lines)
	})

	t.Run("LoadLines with invalid file path returns error", func(t *testing.T) {
		t.Parallel()

		linesCh, errCh := LoadLines("/invalid/path/\000/file.txt")
		assert.Error(t, <-errCh)
		assert.Empty(t, <-linesCh)
	})
}
