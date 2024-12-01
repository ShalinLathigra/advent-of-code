package main

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

// alrighty, so what is the order of operations here?

var isReal bool
var partOneValueText string
var partTwoValueText string

// Goal is to create an ORDERED set
// Requirements:
// Values stored in a zero indexed order
// Number of copies stored with each value
// Values ascending from low to high
/*
	OrdRefSet
		counts[]int	// Array ordered by when each is first seen
		keys[]int
		indexOrder[]int // What order to process above elements in

		Add(X) {
			index := Contains(keys, X)
			index == -1 {
				append entry to both maps
				insert new entry into indexOrder with this new value

				add new indexToCount entry (e)
			}
		}

		OrdRefSet.Get(0) ->
*/

type CountSet struct {
	name   string
	keys   []int64
	counts []int
}

func NewCountSet(name string, len int64) CountSet {
	return CountSet{
		name:   name,
		keys:   make([]int64, 0, len),
		counts: make([]int, 0, len),
	}
}

func (set *CountSet) Add(val int64) {

	if index := slices.Index(set.keys, val); index < 0 {
		set.keys = append(set.keys, val)
		set.counts = append(set.counts, 1)
	} else {
		set.counts[index] += 1
	}
}

func (set *CountSet) GetOrder() (ret []int) {
	ret = make([]int, 0, len(set.keys))
KeyLoop:
	for j, key := range set.keys {
		for i, index := range ret {
			// if key < set.keys[index]
			if key < set.keys[index] {
				ret = slices.Insert(ret, i, j)
				continue KeyLoop
			}
			// if this key is smaller than the one at this index,
		}
		ret = append(ret, j)
	}
	return
}

func main() {
	flag.BoolVar(&isReal, "real", false, "Enable real test")
	flag.Parse()

	filepath := "./2024/01/test"
	var scanner bufio.Scanner
	fmt.Println(isReal)
	if isReal {
		filepath = "./2024/01/real"
	}
	if f, err := os.Open(filepath); err != nil {
		panic(err)
	} else {
		scanner = *bufio.NewScanner(f)
	}

	// if test, read the first line
	if !isReal {
		scanner.Scan()
		partOneValueText = scanner.Text()
		scanner.Scan()
		partTwoValueText = scanner.Text()
	}
	start := time.Now()
	leftSet := NewCountSet("left", 7)
	rightSet := NewCountSet("right", 7)

	for scanner.Scan() {
		chunks := strings.SplitAfter(scanner.Text(), "   ")
		if len(chunks) != 2 {
			panic(fmt.Errorf("not enough chunks in line:%s", scanner.Text()))
		}
		if left, err := strconv.ParseInt(strings.TrimSpace(chunks[0]), 10, 32); err != nil {
			panic(err)
		} else {
			leftSet.Add(left)
		}
		if right, err := strconv.ParseInt(strings.TrimSpace(chunks[1]), 10, 32); err != nil {
			panic(err)
		} else {
			rightSet.Add(right)
		}
	}
	fmt.Println("Parsed contents", time.Since(start).Milliseconds())
	// for each element in this set,
	leftIndex, leftOrder, leftCounts := 0, leftSet.GetOrder(), slices.Clone(leftSet.counts)
	rightIndex, rightOrder, rightCounts := 0, rightSet.GetOrder(), slices.Clone(rightSet.counts)
	overallDistance := 0
	for leftIndex < len(leftOrder) && rightIndex < len(rightOrder) {
		overlap := min(leftCounts[leftOrder[leftIndex]], rightCounts[rightOrder[rightIndex]])
		distance := int(math.Abs(float64(leftSet.keys[leftOrder[leftIndex]] - rightSet.keys[rightOrder[rightIndex]])))
		overallDistance += overlap * distance
		leftCounts[leftOrder[leftIndex]] -= overlap
		rightCounts[rightOrder[rightIndex]] -= overlap
		if leftCounts[leftOrder[leftIndex]] == 0 {
			leftIndex += 1
		}
		if rightCounts[rightOrder[rightIndex]] == 0 {
			rightIndex += 1
		}

	}
	fmt.Println("Processed Part One", time.Since(start).Milliseconds())
	fmt.Println("Test Value", partOneValueText, "Part One Value", overallDistance)
	var similarityScore int64 = 0
	for i, key := range leftSet.keys {
		if index := slices.Index(rightSet.keys, key); index == -1 {
			continue
		} else {
			similarityScore += key * int64(leftSet.counts[i]*rightSet.counts[index])
		}
	}
	fmt.Println("Processed Part Two", time.Since(start).Milliseconds())
	fmt.Println("Test Value", partTwoValueText, "Part Two Value", similarityScore)
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

// alrighty, so it'd be annoying to have to create each
