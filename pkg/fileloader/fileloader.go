package fileloader

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	// DefaultLineDelimiter is the default line delimiter for end of lines.
	DefaultLineDelimiter byte = '\n'
	// DefaultChunkSize is the default chunk size for reading files (1MB).
	DefaultChunkSize = 1 * 1024 * 1024
)

// Option represents a configuration option for the file loader.
type Option func(*Config)

type Config struct {
	chunkSize     int                       // size of chunks to read from the file
	lineParser    func(string) (any, error) // optional parser for lines
	lineDelimiter byte                      // delimiter for line separation
}

// WithChunkSize sets a custom chunk size for reading.
func WithChunkSize(size int) Option {
	return func(cfg *Config) {
		if size > 0 {
			cfg.chunkSize = size
		}
	}
}

// WithLineParser sets a custom line parser
func WithLineParser[T any](parser func(string) (T, error)) Option {
	return func(cfg *Config) {
		cfg.lineParser = func(s string) (any, error) {
			return parser(s)
		}
	}
}

// WithLineDelimiter sets a custom line delimiter
func WithLineDelimiter(delimiter byte) Option {
	return func(cfg *Config) {
		cfg.lineDelimiter = delimiter
	}
}

// LoadLines reads a file in chunks and sends raw bytes through a channel
func LoadLines(filePath string, options ...Option) (<-chan []byte, <-chan error) {
	cfg := Config{
		chunkSize:     DefaultChunkSize,
		lineDelimiter: DefaultLineDelimiter,
	}

	for _, opt := range options {
		opt(&cfg)
	}

	resultCh := make(chan []byte)
	errCh := make(chan error, 1) // buffered channel to prevent blocking

	file, err := os.Open(filePath)
	if err != nil {
		close(resultCh)
		errCh <- &FileError{LoaderError{Message: "failed to open file", Err: err}}
		close(errCh)
		return resultCh, errCh
	}

	go func() {
		defer close(resultCh)
		defer close(errCh)
		defer file.Close()

		reader := bufio.NewReaderSize(file, cfg.chunkSize)

		for lineNum := 1; ; lineNum++ {
			lineData, err := reader.ReadString(cfg.lineDelimiter)
			if err != nil && err != io.EOF {
				errCh <- &LineError{LoaderError{Message: "failed to read line", Err: err}}
				continue
			}

			if len(lineData) > 0 {
				if cfg.lineParser != nil {
					if parsed, err := cfg.lineParser(lineData); err == nil {
						switch v := parsed.(type) {
						case string:
							resultCh <- []byte(v)
						case []byte:
							resultCh <- v
						default:
							errCh <- fmt.Errorf("could not parse line: %w", err)
						}
						continue
					}
					errCh <- fmt.Errorf("could not parse line: %w", err)
				} else {
					chunk := make([]byte, len(lineData))
					copy(chunk, lineData)
					resultCh <- chunk
				}
			}

			if err == io.EOF {
				break
			}
		}
	}()
	return resultCh, errCh
}
