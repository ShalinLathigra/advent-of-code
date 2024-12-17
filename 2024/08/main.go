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
	partOneExpected int = 14
	partTwoExpected int = 34
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
	
    size, antennae := parseInput(scanner)
	if isPartTwo {
		runPartTwo(size, antennae)
	} else {
		runPartOne(size, antennae)
	}
}

func parseInput(scanner *bufio.Scanner) (size vec2, antennae map[rune][]vec2) {
    // need to construct a set of coordinates based on the frequency types
    // also need to extract the dimensions of the grid to account for those that would
    // be off the map
	size.y = 0
    size.x = 0
    antennae = make(map[rune][]vec2)
    for scanner.Scan() {
		text := scanner.Text()
        if size.x == 0 {
            size.x = len(text)
        }
        for x, rune := range text {
            if rune == '.' {
                continue
            }
            if _, ok := antennae[rune]; !ok {
                antennae[rune] = make([]vec2, 0, 8)
            }
            antennae[rune] = append(antennae[rune], newVec2(x, size.y))
        }
        size.y += 1
	}
	return
}

func runPartOne(size vec2, antennae map[rune][]vec2) {
	fmt.Println("Hit Part One", size)
    antinodes := make([]vec2, 0, 32)
    for _, nodes := range antennae {
        for i := range nodes {
            if i + 1 >= len(nodes) {
                break
            }
            for _, node := range generateNaiveAntinodes(nodes[i], nodes[i+1:]) {
                antinodes = appendNode(antinodes, node, size)
            }
        }
        // displacement = head - tail
        // head <----- tail
        // antinodes are head + displacement, tail - displacement
        
    }
	fmt.Println("Done Part One", len(antinodes), "if test", partOneExpected)
}

func generateNaiveAntinodes(nodeA vec2, list []vec2) (antinodes []vec2){
    antinodes = make([]vec2, 0, len(list) * 2)
    for _, nodeB := range list {
        displacement := sub(nodeA, nodeB)
        antinodes = append(antinodes, add(displacement, nodeA), sub(nodeB, displacement))
    }
    return
}

func appendNode(list []vec2, node vec2, size vec2) []vec2 {
    // do not add if it already exists
    if slices.Contains(list, node){
        return list
    }
    if !isInBounds(node, size) {
        return list
    }
    return append(list, node)
}


func runPartTwo(size vec2, antennae map[rune][]vec2) {
	fmt.Println("Hit Part Two")
    antinodes := make([]vec2, 0, 32)
    for _, nodes := range antennae {
        for i := range nodes {
            if i + 1 >= len(nodes) {
                break
            }
            for _, node := range generateFullAntinodes(nodes[i], nodes[i+1:], size) {
                antinodes = appendNode(antinodes, node, size)
            }
        }
    }
	fmt.Println("Done Part Two", len(antinodes), "if test", partTwoExpected)
}

func generateFullAntinodes(nodeA vec2, list []vec2, size vec2) (antinodes []vec2){
    antinodes = make([]vec2, 0, len(list) * 2)
    for _, nodeB := range list {
        displacement := sub(nodeA, nodeB)
        i := 0
        for {
            newNodeA := add(nodeA, mult (displacement, i))
            if isInBounds(newNodeA, size) {
                antinodes = append(antinodes, newNodeA)
            } else {
                break
            }
            i += 1
        }
        i = 0
        for {
            newNodeB := sub(nodeB, mult (displacement, i))
            if isInBounds(newNodeB, size) {
                antinodes = append(antinodes, newNodeB)
            } else {
                break
            }
            i += 1
        }
    }
    return
}

func isInBounds(src vec2, bounds vec2) bool {
    return src.x >= 0 && src.y >= 0 && src.x < bounds.x && src.y < bounds.y
}

type vec2 struct {
    x int
    y int
}

func newVec2(x int, y int) (vec vec2) {
    return vec2{
        x:x,
        y:y,
    }
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

func sub(a vec2, b vec2) vec2 {
	return vec2 {
		x: a.x - b.x,
		y: a.y - b.y,
	}
}


func mapSlice[T any, U any](arr []T, mapFunc func(T) U) (mappedArr []U) {
	mappedArr = make([]U, 0, len(arr))
	for _, a := range arr {
		mappedArr = append(mappedArr, mapFunc(a))
	}
	return
}
