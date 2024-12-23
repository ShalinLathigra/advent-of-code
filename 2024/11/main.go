package main

// https://javorszky.co.uk/2024/12/12/advent-of-code-2024-day-11/
// Skimming over first bit of this wiki gave hint about memoization and 
// caching count specifically rather than the whole lists
// when called, pass in test text, expected value

import (
	"bufio"
	"flag"
	"fmt"
    "math"
	"os"
    "slices"
    "strconv"
    "strings"
    "time"
)

var isPartTwo bool

const (
	partOneExpected int = 55312
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
	
    nums := parseInput(scanner)
	if isPartTwo {
		runPartTwo(nums)
	} else {
		runPartOne(nums)
	}
}

func parseInput(scanner *bufio.Scanner) []int{
	scanner.Scan()
	text := scanner.Text()
	return mapSlice(strings.Split(text, " "), func(in string) (out int) {
        if out, err := strconv.Atoi(in); err != nil {
            panic(err)
        } else {
            return out
        } 
        return -1
    })
}

func runPartOne(nums []int) {
	fmt.Println("Hit Part One", nums)
    // starting with nums, process left to right through the set applying rules
    // one by one
    
    startSet := nums
    for range 25 {
        swapSet := make([]int, 0, len(startSet))
        for _, num := range startSet {
            if num == 0 {
                swapSet = append(swapSet, 1)
            } else if numDigits := getOrder(num); numDigits % 2 == 0 {
                halfDigs := int(math.Pow(10, float64(numDigits / 2)))
                swapSet = append(swapSet, num / halfDigs)
                swapSet = append(swapSet, num - (num / halfDigs) * halfDigs)
            } else {
                swapSet = append(swapSet, num * 2024)
            }
            // if 0, add a 1
            // if even number of digits, split in two chunks
            // else multiply by 2024
        }
        startSet = swapSet
    }
	fmt.Println("Done Part One", len(startSet), "if test", partOneExpected)
}

func getOrder(in int) (ord int) {
    return int(math.Log10(float64(in))) + 1
}

// what we need to do is establish a relationship between
// # iterations, value, resulting number of terms
type pair [2]int
var cachedTerms = make(map[pair]int)
func runPartTwo(nums []int) {
	fmt.Println("Hit Part Two")
    start := time.Now()
    sum := 0
    for i, term := range nums {
        sum += expand(term, 75)
        duration := time.Since(start).Milliseconds()
        fmt.Println(i, sum, duration)
    }
	fmt.Println("Hit Part Two", sum, "if test", partTwoExpected)
}

func expand(num int, iters int) (out int) {
    p := pair{num, iters}
    if count, ok := cachedTerms[p]; ok {
        return count
    }
    if iters == 0 {
        return 1
    }
    if num == 0 {
        out = expand(1, iters - 1)
    } else if numDigits := getOrder(num); numDigits % 2 == 0 {
        halfDigs := int(math.Pow(10, float64(numDigits / 2)))
        out = expand(num / halfDigs, iters - 1) +
            expand(num - (num / halfDigs) * halfDigs, iters - 1)
    } else {
        out = expand(num * 2024, iters - 1)
    }
    cachedTerms[p] = out
    return
}

func newEmptyTermData() [][]int {
    return slices.Repeat([][]int{make([]int, 0, 5)}, 75)
}

func mapSlice[T any, U any](arr []T, mapFunc func(T) U) (mappedArr []U) {
	mappedArr = make([]U, 0, len(arr))
	for _, a := range arr {
		mappedArr = append(mappedArr, mapFunc(a))
	}
	return
}
