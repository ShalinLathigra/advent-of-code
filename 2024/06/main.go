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
	
    start, grid := parseInput(scanner)
    // fmt.Println("start", start)
    // printGrid(grid)
	if isPartTwo {
		runPartTwo(start, grid)
	} else {
		runPartOne(start, grid)
	}
}

type tileState int

const (
    Empty tileState = iota
    Wall
    Traversed
)

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

func parseInput(scanner *bufio.Scanner) (start vec2, grid [][]tileState) {
    y := 0
    grid = make([][]tileState, 0, 12)
    // get dimensions of the thing? For each Y, add a new X
	for scanner.Scan() {
        if y > len(grid)-1 {
            grid = append(grid, make([]tileState, 0, 12)) 
        }
	    for x, rune := range scanner.Text() {
            switch rune {
            case '^':
                start = vec2{x: x, y: y}
                grid[y] = append(grid[y], Traversed)
            case '#':
                grid[y] = append(grid[y], Wall)
            default:
                grid[y] = append(grid[y], Empty)
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

func runPartOne(start vec2, grid [][]tileState) {
	fmt.Println("Hit Part One")
    dir := vec2{0, -1}
    pos := start
    count := 1
    for {
        next := add(pos, dir)
        if next.y < 0 || next.y >= len(grid) || next.x < 0 || next.x >= len(grid[0]) {
            break
        }
        switch grid[next.y][next.x] {
        case Empty:
            count += 1
            grid[next.y][next.x] = Traversed
        case Wall:
            dir = turnRight(dir)
            continue
        }
        pos = next
    }
	fmt.Println("Done Part One", count, "if test", partOneExpected)
}

func printGrid(grid [][]tileState) {
    fmt.Println("grid")
    for line := range grid {
        fmt.Println(line, grid[line])
    }
}

type obstacle struct {
    vec2
    dir vec2
}

func runPartTwo(start vec2, grid [][]tileState) {
	fmt.Println("Hit Part Two")
    /*
    for this one I think we're back to the first version of parsing. We really only care
    about the positions of the obstacles.

    1. Parse out the path the guard is taking
    2. Find any points where a an up intersects a right, right intersects down, etc.
    3. 


    Really just need any combination of three obstacles such that three things are true:

    HitFromBottom is one right of HitFromLeft
    HitFromLeft is one up from HitFromTop
    HitFromTop is one left from HitFromRight
    HitFromRight is one down from HitFromBottom

    If these three things are true and there are no other obstacles interfering with the
    paths between these things, then we have a valid segment

    Therefore, when walking through obstacles, divide them up into "hit from top, bottom,
    right, and left"

    Divvy them up based on X, Y coords into lists that can be easily searched.

    during traversal, create a 3 deep list for every obstacle that we hit
    obstacles[i] = [i]
    obstacles[i-1] = [i-1, i]
    obstacles[i-2] = [i-2, i-1, i]

    if there are three obstacles, and they have the correct orientations and offsets
    then we can check the paths and see what happens
    */
    dir := vec2{0, -1}
    pos := start
    count := 0
    obstacles := make([][]obstacle, 0, 12)
    for {
        next := add(pos, dir)
        if next.y < 0 || next.y >= len(grid) || next.x < 0 || next.x >= len(grid[0]) {
            break
        }
        if grid[next.y][next.x] == Wall {
            dir = turnRight(dir)
            obstacles = append(obstacles, make([]obstacle, 0, 3))
            i := len(obstacles)-1
            curr := obstacle{
                vec2: pos, //vec2{x:int(next.x), y: int(next.y)},
                dir: dir,
            }
            obstacles[i] = append(obstacles[i], curr)
            if i >= 1 {
                obstacles[i-1] = append(obstacles[i-1], curr)
            }
            if i >= 2{
                obstacles[i-2] = append(obstacles[i-2], curr)
            }
            continue
        }
        pos = next
    }
    for i, chain := range(obstacles) {
        fmt.Println(i, chain)
        if len(chain) == 3 && checkChain(grid, chain) {
            count += 1
            fmt.Println("Passed", count)
        }
    }
	fmt.Println("Hit Part Two", count, "if test", partTwoExpected)
}

func checkChain(grid [][]tileState, chain []obstacle) bool {
    // go along the line from the last obstacle, adding direction each time
    // if we reach a point where rotatedDir dot vector towards chain[0] > 0
    // then we have a valid combination

    // could accomplish all of this with the original parsing method, but not going to
    // bother re-implementing atm
    // if chain[2].x == chain[1].x, then targetting vec2{chain[2].y, chain[0].x]
    var target vec2
    if chain[2].x == chain[1].x {
        target = vec2{chain[0].x, chain[2].y}
    } else {
        target = vec2{chain[2].x, chain[0].y}
    }

    // must be able to walk from chain[2].vec2 to target along chain[2].dir
    // without encountering an obstacle
    toTarget := sub(target, chain[2].vec2)
    steps := dot(chain[2].dir, toTarget)
    if steps <= 0 {
        panic("impossible traversal, should never happen")
    }
    fmt.Println(chain, target, toTarget, steps)

    for i := range(steps) {
        testPoint := add(chain[2].vec2, scale(chain[2].dir, i))
        fmt.Println(testPoint)
        if grid[testPoint.y][testPoint.x] == Wall {
            fmt.Println("Hit Wall")
            return false
        }
    }
    return true
}

func mapSlice[T any, U any](arr []T, mapFunc func(T) U) (mappedArr []U) {
	mappedArr = make([]U, 0, len(arr))
	for _, a := range arr {
		mappedArr = append(mappedArr, mapFunc(a))
	}
	return
}
