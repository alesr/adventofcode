package day3

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/alesr/adventofcode/pkg/fileloader"
)

var segmentPool = sync.Pool{
	New: func() any {
		b := make([]byte, 0, 32)
		return &b
	},
}

type operation func(x, y int) int

func Part1(inputPath string) int64 {
	res, err := solve(inputPath, processLine)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to solve part 1: %v\n", err)
		return 0
	}
	return res
}

func processLine(line []byte) int64 {
	results := runInstructions(line, func(x, y int) int {
		return x * y
	})

	var sum int64
	for _, result := range results {
		sum += int64(result)
	}
	return sum
}

func solve(inputPath string, processFn func([]byte) int64) (int64, error) {
	linesCh, errCh := fileloader.LoadLines(inputPath)

	var (
		numWorkers = runtime.NumCPU() * 2
		wg         sync.WaitGroup
		totalSum   atomic.Int64
	)

	workCh := make(chan []byte, numWorkers*2)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for line := range workCh {
				result := processFn(line)
				totalSum.Add(result)
			}
		}()
	}

	go func() {
		for line := range linesCh {
			workCh <- line
		}
		close(workCh)
	}()

	wg.Wait()

	if err := <-errCh; err != nil {
		return 0, fmt.Errorf("error reading file: %w", err)
	}
	return totalSum.Load(), nil
}

func runInstructions(instructions []byte, opFn operation) []int {
	results := make([]int, 0, len(instructions)/8)
	segments := splitInstructions(instructions)

	for _, segment := range segments {
		x, y, ok := parseMulInstruction(segment)
		if ok {
			results = append(results, opFn(x, y))
		}
	}
	return results
}

func splitInstructions(instructions []byte) [][]byte {
	segments := make([][]byte, 0, len(instructions)/8)
	currentSegment := *segmentPool.Get().(*[]byte)
	var inInstruction bool

	for i := 0; i < len(instructions); i++ {
		if i <= len(instructions)-4 &&
			instructions[i] == 'm' &&
			instructions[i+1] == 'u' &&
			instructions[i+2] == 'l' &&
			instructions[i+3] == '(' {
			if len(currentSegment) > 0 {
				newSegment := make([]byte, len(currentSegment))
				copy(newSegment, currentSegment)
				segments = append(segments, newSegment)
				currentSegment = currentSegment[:0]
			}
			inInstruction = true
		}

		if inInstruction {
			currentSegment = append(currentSegment, instructions[i])
			if instructions[i] == ')' {
				newSegment := make([]byte, len(currentSegment))
				copy(newSegment, currentSegment)
				segments = append(segments, newSegment)
				currentSegment = currentSegment[:0]
				inInstruction = false
			}
		}
	}
	segmentPool.Put(&currentSegment)
	return segments
}

func parseMulInstruction(segment []byte) (x, y int, ok bool) {
	if len(segment) < 5 || !bytes.HasPrefix(segment, []byte("mul(")) {
		return 0, 0, false
	}

	content := segment[4 : len(segment)-1]
	if segment[len(segment)-1] != ')' {
		return 0, 0, false
	}

	parts := bytes.Split(content, []byte(","))
	if len(parts) != 2 {
		return 0, 0, false
	}

	xParsed, errX := strconv.Atoi(string(bytes.TrimSpace(parts[0])))
	yParsed, errY := strconv.Atoi(string(bytes.TrimSpace(parts[1])))

	if errX != nil || errY != nil {
		return 0, 0, false
	}

	if xParsed < 1 || xParsed > 999 || yParsed < 1 || yParsed > 999 {
		return xParsed, yParsed, true
	}
	return xParsed, yParsed, true
}
