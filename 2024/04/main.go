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
	partOneExpected int = 18
	partTwoExpected int = 9
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

	var grid [][]byte = make([][]byte, 0, 8)
	for scanner.Scan() {
		// read in all of the input as a 2D byte array
		// starting from the top left, if 
		text := scanner.Text()
		grid = append(grid, []byte(text))
	}

	if isPartTwo {
		runPartTwo(grid)
	} else {
		runPartOne(grid)
	}
}

// Setting up filters
// For each letter that we receive, there are four pairs that need to be checked

/*
      ABC
      DMD
      CBA

For each letter M, we have different requirements for the pairs

Starting with each M we find
Check all eight orders of neighbors

Neighbors defined as an offset direction and multiplier

Pair Offset Index

D  1, 0  0
C  1, 1  1
B  0, 1  2
A -1, 1  3
D -1, 0  4
C -1,-1  5
B  0,-1  6
A  1,-1  7
*/


type vec2 struct {
	x int
	y int
}


var topLeft vec2 = vec2{-1,-1}
var topRight vec2 = vec2{1,-1}
var botLeft vec2 = vec2{-1,1}
var botRight vec2 = vec2{1,1}

var xmasNeighbors = [8]vec2 {
	vec2{ 1, 0},
	botRight,
	vec2{ 0, 1},
	botLeft,
	vec2{-1, 0},
	topLeft,
	vec2{ 0,-1},
	topRight,
}

const (
	RuneX byte = 'X'
	RuneM      = 'M'
	RuneA      = 'A'
	RuneS      = 'S'
)

func runPartOne(grid [][]byte) {
	fmt.Println("Hit Part One")
	count := 0
	for y := range grid {
		for x := range grid[y] {
			coords := vec2{x, y}
			if !sample(grid, coords, RuneM) {
				continue
			}
			for i := range 8 {
				xCoords := add(coords, xmasNeighbors[i])
				if !sample(grid, xCoords, RuneX) {
					continue
				}
				aCoords := add(coords, xmasNeighbors[(i+4)%8])
				if !sample(grid, aCoords, RuneA) {
					continue
				}
				sCoords := add(coords, mult(xmasNeighbors[(i+4)%8], 2))
				if !sample(grid, sCoords, RuneS) {
					continue
				}
				count += 1
			}
		}
	}
	fmt.Println("Done Part One", count, "if test", partOneExpected)
}

/*
Slight variation on the formula. Need to find X Mas symbols

Therefore we now are operating on pairs of offsets

BotPair   =  1,1,  -1,1
RightPair =  1,-1,  1,1
TopPair   =  1,-1, -1,-1
LeftPair  = -1,-1, -1,1
*/

type coordPair struct {
	a vec2
	b vec2
}

var botPair = coordPair{
	a: botLeft,
	b: botRight,
}

var topPair = coordPair{
	a: topLeft,
	b: topRight,
}

var leftPair = coordPair{
	a: topLeft,
	b: botLeft,
}

var rightPair = coordPair{
	a: topRight,
	b: botRight,
}

type masPattern struct {
	m coordPair
	s coordPair
}

var masPatterns [4]masPattern = [4]masPattern {
	masPattern{
		m:botPair,
		s:topPair,
	}, masPattern{
		m:topPair,
		s:botPair,
	}, masPattern{
		m:rightPair,
		s:leftPair,
	}, masPattern{
		m:leftPair,
		s:rightPair,
	},
}

func runPartTwo(grid [][]byte) {
	fmt.Println("Hit Part Two")
	count := 0
	for y := range grid {
		for x := range grid[y] {
			coords := vec2{x, y}
			if !sample(grid, coords, RuneA) {
				continue
			}
			for i, mas := range masPatterns{
				if !sample(grid, add(coords, mas.m.a), RuneM) {
					continue
				}
				if !sample(grid, add(coords, mas.m.b), RuneM) {
					continue
				}
				if !sample(grid, add(coords, mas.s.a), RuneS) {
					continue
				}
				if !sample(grid, add(coords, mas.s.b), RuneS) {
					continue
				}
				fmt.Println(x, y, i, mas)
				count += 1
			}
		}
	}
	fmt.Println("Hit Part Two", count, "if test", partTwoExpected)
}

func sample(grid [][]byte, coords vec2, expected byte) bool {
	if coords.y < 0 || coords.y >= len(grid) {
		return false
	}
	if coords.x < 0 || coords.x >= len(grid[coords.y]) {
		return false
	}
	return grid[coords.y][coords.x] == expected
}

func mult(a vec2, b int) vec2 {
	return vec2 {
		x: a.x * b,
		y: a.y * b,
	}
}

func add(a vec2, b vec2) vec2 {
	return vec2 {
		x: a.x + b.x,
		y: a.y + b.y,
	}
}

