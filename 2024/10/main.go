package main

// when called, pass in test text, expected value

import (
	"bufio"
	"flag"
	"fmt"
	"os"
    "slices"
)

var isPartTwo bool

const (
	partOneExpected int = 36
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
    trailHeads, heightMap:= parseInput(scanner)
	if isPartTwo {
		runPartTwo(trailHeads, heightMap)
	} else {
		runPartOne(trailHeads, heightMap)
	}
}

func parseInput(scanner *bufio.Scanner) (trailHeads []vec2, heightMap[][]int) {
    trailHeads = make ([]vec2, 0, 8)
    heightMap= make ([][]int, 0, 8)
    y := 0
	for scanner.Scan() {
		text := scanner.Text()
        line := make([]int, 0, 8)
        for x, rune := range text {
            if rune == '0' {
                trailHeads = append(trailHeads, vec2{x: x, y: y})
            }
            line = append(line, int(rune - '0'))
        }
        y += 1
        heightMap= append(heightMap, line)
	}
	return
}

func runPartOne(trailHeads []vec2, heightMap[][]int) {
	fmt.Println("Hit Part One")
    count := 0
    for _, trail := range trailHeads {
        points := make([]vec2, 0, 8)
        points = append(points, trail)
        for i := range 9 {
            searchTerm := i + 1
            totalNeighbors := make([]vec2, 0, 4 * len(points))
            for _, point := range points {
                for _, dir := range neighbors {
                    if n := add(point, dir); validatePoint(n, heightMap, searchTerm) {
                        if !slices.Contains(totalNeighbors, n) {
                            totalNeighbors = append(totalNeighbors, n)
                        }
                    }
                }
            }
            points = totalNeighbors
            if len(points) == 0 {
                break
            }
        }
        count += len(points)
    }
	fmt.Println("Done Part One", count, "if test", partOneExpected)
}

func validatePointInGrid(p vec2, grid [][]int) bool {
    return !(p.x < 0 || p.y < 0 || p.x >= len(grid[0]) || p.y >= len(grid))
}

func validatePoint(p vec2, grid [][]int, searchTerm int) bool {
    if !validatePointInGrid(p, grid) {
        return false
    }
    return grid[p.y][p.x] == searchTerm
}

func runPartTwo(trailHeads []vec2, heightMap[][]int) {
	fmt.Println("Hit Part Two")
    count := 0
    for _, trail := range trailHeads {
        trails := runDfs(trail, 0, heightMap)
        count += trails
    }
	fmt.Println("Hit Part Two", count, "if test", partTwoExpected)
}

func runDfs(start vec2, currentTerm int, heightMap[][]int) (numRoutes int) {
    if !validatePointInGrid(start, heightMap) {
        return 0
    } else if heightMap[start.y][start.x] != currentTerm {
        return 0
    } else if currentTerm == 9 {
        return 1
    }
    numRoutes += runDfs(add(start, neighbors[0]), currentTerm + 1, heightMap)
    numRoutes += runDfs(add(start, neighbors[1]), currentTerm + 1, heightMap)
    numRoutes += runDfs(add(start, neighbors[2]), currentTerm + 1, heightMap)
    numRoutes += runDfs(add(start, neighbors[3]), currentTerm + 1, heightMap)
    return
}

type vec2 struct {
    x int
    y int
}

func add (a vec2, b vec2) (c vec2) {
    return vec2{x: a.x + b.x, y: a.y + b.y}
}

func sub (a vec2, b vec2) (c vec2) {
    return vec2{x: a.x - b.x, y: a.y - b.y}
}

func dot (a vec2, b vec2) (c int) {
    return a.x * b.x + a.y * b.y
}

func scale (a vec2, b int) (c vec2) {
    return vec2{x: a.x * b, y: a.y * b}
}

var neighbors []vec2 = []vec2{
    vec2{1,0},
    vec2{0,-1},
    vec2{-1,0},
    vec2{0,1},
}

func mapSlice[T any, U any](arr []T, mapFunc func(T) U) (mappedArr []U) {
	mappedArr = make([]U, 0, len(arr))
	for _, a := range arr {
		mappedArr = append(mappedArr, mapFunc(a))
	}
	return
}
