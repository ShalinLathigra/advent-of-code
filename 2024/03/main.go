package main

// when called, pass in test text, expected value

import (
	"bufio"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"os"
)

var isPartTwo bool

const (
	partOneExpected int = 161
	partTwoExpected int = 48
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

	if isPartTwo {
		runPartTwo(scanner)
	} else {
		runPartOne(scanner)
	}
}

func runPartOne(scanner *bufio.Scanner) {
	fmt.Println("Hit Part One")
	re := regexp.MustCompile(`mul\((\d*),(\d*)\)`)
	productSum := 0
	for scanner.Scan() {
		for _, chunk := range re.FindAllSubmatch([]byte(scanner.Text()), -1) {
			productSum += extractAndAdd(chunk[1], chunk[2])
			// if a,err := strconv.Atoi(string(chunk[1])); err != nil {
			// 	panic(err)
			// } else if b, err := strconv.Atoi(string(chunk[2])); err != nil {
			// 	panic(err)
			// } else {
			// 	productSum += a * b
			// }
		}
		// fmt.Printf("%q\n",re.FindAllSubmatch([]byte(scanner.Text()), -1))
	}
	fmt.Println("Done Part One", productSum, "if test", partOneExpected)
}

func runPartTwo(scanner *bufio.Scanner) {
	fmt.Println("Hit Part Two")
	re := regexp.MustCompile(`mul\((\d*),(\d*)\)|do\(\)|don't\(\)`)
	productSum := 0
	doMult := true
	for scanner.Scan() {
		for _, chunk := range re.FindAllSubmatch([]byte(scanner.Text()), -1) {
			switch string(chunk[0]) {
			case "do()":
				doMult = true
				fmt.Println("Found Do")
			case "don't()":
				doMult = false
				fmt.Println("Found Don't")
			default:
				if doMult {
					fmt.Printf("Found mul(%s,%s)\n", chunk[1], chunk[2])
					productSum += extractAndAdd(chunk[1], chunk[2])
				} else {
					fmt.Printf("skip mul(%s,%s)\n", chunk[1], chunk[2])
				}
			}
		}
	}
	fmt.Println("Hit Part Two", productSum, "if test", partTwoExpected)
}

func extractAndAdd(aByte []byte, bByte []byte) int {
	if a, err := strconv.Atoi(string(aByte)); err != nil {
		panic(err)
	} else if b, err := strconv.Atoi(string(bByte)); err != nil {
		panic(err)
	} else {
		return a * b
	}
	return 0
}
