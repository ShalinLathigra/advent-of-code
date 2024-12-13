package main

// when called, pass in test text, expected value

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var isPartTwo bool

const (
	partOneExpected int = 123
	partTwoExpected int = 246
)

func main() {
	// parse flags
	flag.BoolVar(&isPartTwo, "p", false, "perform part two calculations")
	flag.Parse()
	
	// read in text as a string
	var scanner *bufio.Scanner
	if info, err := os.Stdin.Stat(); err != nil {
		panic(err)
	} else {
		if info.Mode() & os.ModeCharDevice != 0 {
			panic("no input provided in stdin")
		}
		scanner = bufio.NewScanner(os.Stdin)
	}

	// split into rules and 
	
	parseInput(scanner)
	if isPartTwo {
		runPartTwo()
	} else {
		runPartOne()
	}
}

func parseInput(scanner *bufio.Scanner) {
	for scanner.Scan() {
		// text := scanner.Text()
	}
	return
}

func runPartOne() {
	fmt.Println("Hit Part One")
	fmt.Println("Done Part One", "if test", partOneExpected)
}
func runPartTwo() {
	fmt.Println("Hit Part Two")
	fmt.Println("Hit Part Two", "if test", partTwoExpected)
}

func mapSlice[T any, U any](arr []T, mapFunc func(T) U) (mappedArr []U) {
	mappedArr = make([]U, 0, len(arr))
	for _, a := range arr {
		mappedArr = append(mappedArr, mapFunc(a))
	}
	return
}
