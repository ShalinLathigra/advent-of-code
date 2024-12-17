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
	partOneExpected int = 1928
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
	
    input, length := parseInput(scanner)
	if isPartTwo {
		runPartTwo()
	} else {
		runPartOne(input, length)
	}
}

type file struct {
    id int
    offset int
    length int
    blankLength int
}

func parseInput(scanner *bufio.Scanner) (diskMap []file, compactLength int) {
	scanner.Scan()
    text := scanner.Text()
    diskMap = make([]file, 0, len(text)/2+1)
    compactLength = 0
    offset := 0
    for id, rune := range text {
        isEmpty := id%2==1
        length := int(rune - 48)
        blankLength := 0
        if isEmpty {
            continue
        }
        if id+1 < len(text) {
            blankLength = int(text[id+1]-48)
        }
        file := file{
            id: id/2,
            offset: offset,
            length: length,
            blankLength: blankLength,
        }
        offset += length
        diskMap = append(diskMap, file)
        compactLength += length
    }

    return
}

func runPartOne(input []file, length int) {
	fmt.Println("Hit Part One")
    // what do I do? First, generate the set of actual files that exist
    // creating an array of integers
    // no clear idea how long it should be really, we could probably make a 
    denseMap := make([]int, 0, length)
    forwardIndex, backIndex := 0, len(input) - 1
    for forwardIndex <= backIndex {
        current := input[forwardIndex]
        for range current.length {
            denseMap = append(denseMap, current.id)
        }
        forwardIndex += 1
        for range current.blankLength {
            // find first element starting from blankLength that has a length > 0
            var last *file
            for last == nil {
                if input[backIndex].length> 0 {
                    last = &input[backIndex]
                } else {
                    backIndex -= 1
                }
            }
            denseMap = append(denseMap, last.id)
            last.length -= 1
            if last.length <= 0 {
                backIndex -= 1
            }
        }
    }
    // extra element, not sure why, but skipping it for now
    // I think it's because of how I'm adding the forward and back elements
    sum := 0
    for i, block := range denseMap[:len(denseMap)] {
        sum += i * block
    }
	fmt.Println("Done Part One", sum, "if test", partOneExpected)
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
