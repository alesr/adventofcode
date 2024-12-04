package main

import (
	"fmt"

	"github.com/alesr/adventofcode/aoc24/day1"
	"github.com/alesr/adventofcode/aoc24/day2"
	"github.com/alesr/adventofcode/aoc24/day3"
)

func main() {
	var totalStars int

	d1p1 := day1.Part1("aoc24/day1/input")
	fmt.Printf("DAY1P1=%d\n", d1p1)

	d1p2 := day1.Part2("aoc24/day1/input")
	fmt.Printf("DAY1P2=%d\n", d1p2)

	totalStars++ // Day 1 complete

	d2p1 := day2.Part1("aoc24/day2/input")
	fmt.Printf("DAY2P1=%d\n", d2p1)

	d2p2 := day2.Part2("aoc24/day2/input")
	fmt.Printf("DAY2P2=%d\n", d2p2)

	totalStars++ // Day 2 complete

	fmt.Printf("STARS=%d\n", totalStars)

	d3p1 := day3.Part1("aoc24/day3/input")
	fmt.Printf("DAY3P1=%d\n", d3p1)
}
