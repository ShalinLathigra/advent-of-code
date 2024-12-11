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
	partOneExpected int = 41
	partTwoExpected int = 246
)

// points organized into two 2d arrays
// First one, stores organized by X axis
// Second one, organized by Y axis

/*
..............
..#..#..#..#..
..............
...#..#..#..#.
..............
.#..#..#..#...
..............

this makes two grids, Y axis one looks like:

[],[1],[5],[3],[1],[5],[3],[1],[5],[3],[1],[5],[3],[]

X axis one is like

[],[1,4,7,10],[],[3,6,9,12],[],[2,5,8,11],[]


then, when we are flicking through we go:
"I am now at x=12, y=4, and I'm going up."
look at all elements of array x[12], if any of them are > my thing, then we ball

*/
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
	
    start, rows, cols := parseInput(scanner)
    fmt.Println("start", start)
    fmt.Println("rows", rows)
    fmt.Println("cols", cols)
	if isPartTwo {
		runPartTwo()
	} else {
		runPartOne(start, rows, cols)
	}
}

type vec2 struct {
    x int
    y int
}

func parseInput(scanner *bufio.Scanner) (start vec2, rows [][]int, cols [][]int) {
    y := 0
    rows = make([][]int, 0, 12)
    cols = make([][]int, 0, 12)
    // get dimensions of the thing? For each Y, add a new X
	for scanner.Scan() {
        if y > len(rows)-1 {
            rows = append(rows, make([]int, 0, 12)) 
        }
	    for x, rune := range scanner.Text() {
            if x > len(cols)-1 {
                cols = append(cols, make([]int, 0, 12))
            }
            switch rune {
            case '^':
                start = vec2{x: x, y: y}
            case '#':
                rows[y] = append(rows[y], x)
                cols[x] = append(cols[x], y)
            default:
            }
        }
        y += 1
	}
	return
}

func turnRight(dir vec2) (rotated vec2) {
    rotated.x = -dir.y
    rotated.y = dir.x
    return
}

func runPartOne(start vec2, rows [][]int, cols [][]int) {
	fmt.Println("Hit Part One")
    dir := vec2{0, -1}
    pos := start
    count := 1
    TraversalLoop:
    for {
        // for y from 0 if magnitude > 0, from len(y) if magnitude < 0
        if dir.y < 0 {
            for y := range cols[pos.x] {
                i := len(cols[pos.x]) - 1 - y // work backwards through the thing
                // stop when/if we encounter a number that is below the current y value
                if cols[pos.x][i] < pos.y {
                    dist := pos.y - cols[pos.x][i] - 1
                    fmt.Println(pos.x, pos.y, " hit ", pos.x, i, "traveled", dist)
                    count += dist
                    pos.y = cols[pos.x][i] + 1
                    dir = turnRight(dir)
                    continue TraversalLoop
                }
            }
        } else if dir.y > 0 {
            for y := range cols[pos.x] {
                // stop when/if we encounter a number that is below the current y value
                if cols[pos.x][y] > pos.y {
                    dist := cols[pos.x][y] - pos.y - 1
                    fmt.Println(pos.x, pos.y, " hit ", pos.x, y, "traveled", dist)
                    count += dist
                    pos.y = cols[pos.x][y] - 1
                    dir = turnRight(dir)
                    continue TraversalLoop
                }
            }
        } else if dir.x < 0 {
            for x := range rows[pos.y] {
                i := len(rows[pos.y]) - 1 - x // work backwards through the thing
                // stop when/if we encounter a number that is below the current y value
                if rows[pos.y][i] < pos.x {
                    dist := pos.x - rows[pos.y][i] - 1
                    fmt.Println(pos.x, pos.y, " hit ", i, pos.y, "traveled", dist)
                    count += dist
                    pos.x = rows[pos.y][i] + 1
                    dir = turnRight(dir)
                    continue TraversalLoop
                }
            }
        } else if dir.x > 0 {
            for x := range rows[pos.y] {
                // stop when/if we encounter a number that is below the current y value
                if rows[pos.y][x] > pos.x {
                    dist := rows[pos.y][x] - pos.x - 1
                    fmt.Println(pos.x, pos.y, " hit ", x, pos.y, "traveled", dist)
                    count += dist
                    pos.x = rows[pos.y][x] - 1
                    dir = turnRight(dir)
                    continue TraversalLoop
                }
            }
        }
        // if we ever reach here, break out of the loop
        break
    }
	fmt.Println("Done Part One", count, "if test", partOneExpected)
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
