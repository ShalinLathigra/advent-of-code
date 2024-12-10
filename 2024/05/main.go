package main

// when called, pass in test text, expected value

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var isPartTwo bool

const (
	partOneExpected int = 143
	partTwoExpected int = 123
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
	
	rules, updates := parseInput(scanner)
	if isPartTwo {
		runPartTwo(rules, updates)
	} else {
		runPartOne(rules, updates)
	}
}

// parses out instructions of the form
// A|B
// <empty line>
// 1,2,3
// returns:
/*
	map pairing every rule start with every number that must appear after it
	set containing all of the updates as arrays of ints
*/
func parseInput(scanner *bufio.Scanner) (rules map[int][]int, updates [][]int) {
	rules = make(map[int][]int)
	updates = make([][]int, 0, 8)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "|") {
			chunks := mapSlice(strings.Split(text, "|"), convert)
			if _, ok := rules[chunks[0]]; !ok {
				rules[chunks[0]] = make([]int, 0, 4)
			}
			rules[chunks[0]] = append(rules[chunks[0]], chunks[1])
		} else if strings.Contains(text, ",") {
			chunks := mapSlice(strings.Split(text, ","), convert)
			updates = append(updates, chunks)
		}
	}
	return
}

func convert(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}


func runPartOne(rules map[int][]int, updates [][]int) {
	fmt.Println("Hit Part One")
	count := 0
	UpdateLoop:
	for i, update := range updates {
		fmt.Println(i, update)
		for j, num := range update {
			if _, ok := rules[num]; !ok {
				continue // to next number
			} else {
			 	// scrub through nums[0-j]
				for _, trg := range update[0:j] {
					if slices.Contains(rules[num], trg) {
						continue UpdateLoop
					}
				}
			}
		}
		// if we reach this point, take element at middle of the update and add it
		count += update[len(update) / 2]
	}
	fmt.Println("Done Part One", count, "if test", partOneExpected)
}

// what if we think about it as ranges instead?

func runPartTwo(rules map[int][]int, updates [][]int) {
	fmt.Println("Hit Part Two")
	count := 0
	for _, update := range updates {
		count += processUpdate(rules, update)
	}
	fmt.Println("Hit Part Two", count, "if test", partTwoExpected)
}

func processUpdate(rules map[int][]int, update[]int) int {
    didUpdate := false
    for i := 0; i < len(update); i ++ {
        if _, ok := rules[update[i]]; !ok {
            continue // to next number
        } else {
            // scrub through nums[0-j]
            for j := 0; j < i; j ++ {
                if slices.Contains(rules[update[i]], update[j]) {
                    didUpdate = true
                    update = slices.Insert(update, i+1, update[j])
                    update = slices.Delete(update, j, j+1)
                    // j -= 1, i -= 1
                    j = j - 1
                    i = i - 1
                }
            }
        }
    }

    if didUpdate {
        return update[len(update)/2]
    }
	return 0
}

func mapSlice[T any, U any](arr []T, mapFunc func(T) U) (mappedArr []U) {
	mappedArr = make([]U, 0, len(arr))
	for _, a := range arr {
		mappedArr = append(mappedArr, mapFunc(a))
	}
	return
}
