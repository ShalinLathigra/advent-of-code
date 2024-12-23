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
	partOneExpected int = 1930
	partTwoExpected int = 1206
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
	
    grid := parseInput(scanner)
	if isPartTwo {
		runPartTwo(grid)
	} else {
		runPartOne(grid)
	}
}

func parseInput(scanner *bufio.Scanner) (grid [][]byte) {
    grid = make([][]byte, 0, 12)
    for scanner.Scan() {
        grid = append(grid, []byte(scanner.Text()))
    }
	return
}

type vec2 struct {
    x int
    y int
}

type cluster struct {
    points []vec2
    variant byte
    minPoint vec2
    maxPoint vec2
}

func maxVec (a vec2, b vec2) (c vec2) {
    return vec2{max(a.x, b.x), max(a.y, b.y)}
}

func minVec (a vec2, b vec2) (c vec2) {
    return vec2{min(a.x, b.x), min(a.y, b.y)}
}

func (c *cluster) AddPoint (p vec2) {
    if !slices.Contains(c.points, p) {
        c.points = append(c.points, p)
    }
    c.maxPoint = maxVec(c.maxPoint, p)
    c.minPoint = minVec(c.minPoint, p)
}

func (c *cluster) Contains (p vec2) bool {
    return slices.Contains(c.points, p)
}

func merge (a cluster, b cluster) (c cluster) {
    c.variant = a.variant
    c.points = slices.Concat(a.points, b.points)
    c.minPoint = minVec(a.minPoint, b.minPoint)
    c.maxPoint = maxVec(a.maxPoint, b.maxPoint)
    return
}

func add (a vec2, b vec2) (c vec2) {
    return vec2{x: a.x + b.x, y: a.y + b.y}
}

func generateClusters (grid [][]byte) (clusters []cluster) {
    clusters = make([]cluster, 0, 16)
    // define clusters
    for y, g := range grid {
        for x, c := range g {
            point := vec2{x, y}
            // check up and right. Are we adjacent to a "c"
            upC := sampleGrid(grid, add(point, up))
            upI:= getClusterFromPoint(add(point, up), clusters)
            upA:= upC == c
            leftC := sampleGrid(grid, add(point, left))
            leftI:= getClusterFromPoint(add(point, left), clusters)
            leftA:= leftC == c
            if upA{
                clusters[upI].AddPoint(point)
                if leftA && upI != leftI {
                    clusters[upI] = merge(clusters[upI], clusters[leftI])
                    clusters = slices.Delete(clusters, leftI, leftI + 1)
                }
            } else if leftA{
                clusters[leftI].AddPoint(point)
            } else {
                newCluster := cluster{
                    points: []vec2{point},
                    variant: c,
                    maxPoint: point,
                    minPoint: point,
                }
                clusters = append(clusters, newCluster)
            }
        }
    }
    return clusters
}

func runPartOne(grid [][]byte) {
    fmt.Println("Hit Part One")
    // could do all the processing at the same time, doing this to avoid
    // extra math for correcting clusters afterwards
    clusters := generateClusters(grid)
    sum := 0
    perimeter := 0
    for _, cluster := range clusters {
        perimeter = 0
        for _, p := range cluster.points {
            for _, dir := range dirs {
                if !cluster.Contains(add(p, dir)) {
                    perimeter += 1
                }
            }
        }
        sum += perimeter * len(cluster.points)
    }
	fmt.Println("Done Part One", sum, "if test", partOneExpected)
}

func getClusterFromPoint(p vec2, clusters []cluster) (index int) {
    for i, cluster := range clusters {
        if cluster.Contains(p) {
            return i
        }
    }
    return -1
}

var up = vec2{0, -1}
var left = vec2{-1, 0}
var down = vec2{0, 1}
var right = vec2{1, 0}

var dirs = []vec2{
    up,
    left,
    down,
    right,
}

func sampleGrid(grid [][]byte, pos vec2) byte {
    if pos.y < 0 || pos.y >= len(grid) || pos.x < 0 || pos.x >= len(grid[0]) {
        return '-'
    }
    return grid[pos.y][pos.x]
}

func runPartTwo(grid [][]byte) {
	fmt.Println("Hit Part Two")
    // clusters are gathered the same way
    // now, we start from a point. If it is surrounded, do nothing
    // if left
    // only consider if we have two adjacent gaps
    clusters := generateClusters(grid)
    sum := 0
    sides := 0
    for i, cluster := range clusters {
        sides = 0
        // bounding box made up of minPoint -> maxPoint
        // from just below to just past the structure (checking n, n+1 per iter) 
        for yOff := range cluster.maxPoint.y - cluster.minPoint.y + 4 {
            testY := cluster.minPoint.y + yOff - 1
            lastPoint := vec2{-500, -500}
            for xOff := range cluster.maxPoint.x - cluster.minPoint.x + 4 {
                testX := cluster.minPoint.x + xOff - 1
                point := vec2{x:testX, y:testY}
                above := add(point, up)
                if cluster.Contains(above) == cluster.Contains(point) {
                    continue
                }
                if testX - lastPoint.x > 1 {
                    fmt.Println("x", point, lastPoint)
                    sides += 1
                }
                lastPoint.x = testX
            }
        }
        for xOff := range cluster.maxPoint.x - cluster.minPoint.x + 4 {
            testX := cluster.minPoint.x + xOff - 1
            lastPoint := vec2{-500, -500}
            for yOff := range cluster.maxPoint.y - cluster.minPoint.y + 4 {
                testY := cluster.minPoint.y + yOff - 1
                point := vec2{x:testX, y:testY}
                above := add(point, left)
                if cluster.Contains(above) == cluster.Contains(point) {
                    continue
                }
                if testY - lastPoint.y > 1 {
                    fmt.Println("y", point, lastPoint)
                    sides += 1
                }
                lastPoint.y = testY
                // as we're scanning, we always want to 
            }
        }
        sum += sides * len(cluster.points)
        fmt.Println(i, cluster.variant, len(cluster.points), sides, sides * len(cluster.points), sum)
    }
	fmt.Println("Hit Part Two", sum, "if test", partTwoExpected)
}

func mapSlice[T any, U any](arr []T, mapFunc func(T) U) (mappedArr []U) {
	mappedArr = make([]U, 0, len(arr))
	for _, a := range arr {
		mappedArr = append(mappedArr, mapFunc(a))
	}
	return
}
