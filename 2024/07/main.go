package main

// when called, pass in test text, expected value

import (
	"bufio"
	"flag"
	"fmt"
    "strconv"
    "strings"
	"os"
)

var isPartTwo bool

const (
	partOneExpected int = 3749
	partTwoExpected int = 11387
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
	
    formulas := parseInput(scanner)
	if isPartTwo {
		runPartTwo(formulas)
	} else {
		runPartOne(formulas)
	}
}

// element 0 is the value
// elements 1+ are the terms
func parseInput(scanner *bufio.Scanner) (formulas [][]int){
    formulas = make([][]int, 0, 9)
	for scanner.Scan() {
		text := scanner.Text()
        chunks := strings.Split(text, ": ")
        if len(chunks) <= 0 {
            panic("unable to split")
        }
        formula := mapSlice(strings.Split(chunks[1], " "), convert)
        formula = append(formula, convert(chunks[0]))
        formulas = append(formulas, formula) 
	}
	return
}

func convert(input string) (output int) {
    output, err := strconv.Atoi(input)
    if err != nil {
        panic(err)
    }
    return
}

func runPartOne(formulas [][]int) {
	fmt.Println("Hit Part One")
    sum := 0
    for i, formula := range(formulas) {
        value := formula[len(formula)-1]
        terms := formula[:len(formula)-1]
        fmt.Printf("[%d] %d: %v\n", i, value, terms)
        if checkTerm(terms[0], value, terms[1:]) {
            sum += value
        }
    }
    // operations = [len(terms)-1]bool
    // true == add, false == multiply
    // There's got to be a more sophisticated way to iterate through though.
    // Just be naive for now.
    // Could expand the whole space?
	fmt.Println("Done Part One", sum, "if test", partOneExpected)
}

func checkTerm(value int, target int, terms []int) bool {
    if len(terms) == 0 {
        return value == target
    }
    if checkTerm(value + terms[0], target, terms[1:]) {
        // fmt.Println("+", value, terms[0], value + terms[0])
        return true
    }
    if checkTerm(value * terms[0], target, terms[1:]) {
        // fmt.Println("*", value, terms[0], value * terms[0])
        return true
    }
    
    // in this case, we're going to take left * significance of right
    // how would I get the largest multiple of 10?
    // quickest and easiest way to do it is below
    next := value * getLarger(terms[0]) + terms[0]
    if checkTerm(next, target, terms[1:]) {
        // fmt.Println("||", value, terms[0], next)
        return true
    }
    return false
}


// takes in a number, returns the largest power of 10 that can contain this number
// 8 -> 10
// 12 -> 100
// 99999 -> 100000
func getLarger(input int) int {
    i := 10

    for {
        if i > input {
            return i
        }
        i *= 10
    }
}

func runPartTwo(formulas [][]int) {
	fmt.Println("Hit Part Two")
    sum := 0
    for i, formula := range(formulas) {
        value := formula[len(formula)-1]
        terms := formula[:len(formula)-1]
        fmt.Printf("[%d] %d: %v\n", i, value, terms)
        if checkTerm(terms[0], value, terms[1:]) {
            sum += value
        }
    }
    // operations = [len(terms)-1]bool
    // true == add, false == multiply
    // There's got to be a more sophisticated way to iterate through though.
    // Just be naive for now.
	fmt.Println("Hit Part Two", sum, "if test", partTwoExpected)
}

func mapSlice[T any, U any](arr []T, mapFunc func(T) U) (mappedArr []U) {
	mappedArr = make([]U, 0, len(arr))
	for _, a := range arr {
		mappedArr = append(mappedArr, mapFunc(a))
	}
	return
}
