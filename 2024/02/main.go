package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var isReal bool

func main() {
	flag.BoolVar(&isReal, "real", false, "Use real dataset")
	flag.Parse()
	filePath := "/home/dev/advent-of-code/2024/02/test"
	if isReal {
		filePath = "/home/dev/advent-of-code/2024/02/real"
	}
	fmt.Println(isReal)
	var scanner bufio.Scanner
	if file, err := os.Open(filePath); err != nil {
		panic(err)
	} else {
		scanner = *bufio.NewScanner(file)
	}

	if !scanner.Scan() {
		panic("failed to read")
	}
	testValue := -1
	if !isReal {
		val, err := strconv.Atoi(strings.Split(scanner.Text(), ":")[1])
		if err != nil {
			panic(err)
		}
		testValue = val
	}

	realValue := 0
	// ScanLoop:
	for scanner.Scan() {
		nums := mapSlice(strings.Split(scanner.Text(), " "), func(str string) int {
			if ret, err := strconv.Atoi(str); err != nil {
				panic(err)
			} else {
				return ret
			}
		})
		// last := nums[0]
		// lastDelta := 0
		// numFailures := 0
		// NumLoop:
		// 	for _, num := range nums[1:] {
		// 		delta := num - last
		// 		if delta*lastDelta < 0 || math.Abs(float64(delta)) < 1 || math.Abs(float64(delta)) > 3 {
		// 			if numFailures > 0 {
		// 				continue ScanLoop
		// 			}
		// 			numFailures += 1
		// 			continue NumLoop
		// 		}
		// 		lastDelta = delta
		// 		last = num
		// 	}
		// 	realValue += 1
		if testLine(nums) {
			realValue += 1
		}
	}
	if !isReal {
		fmt.Println("ExpectedValue:", testValue)
	}
	fmt.Println("FoundValue:", realValue)
}


/* 

   -3  -6  -3
  1  -4  -2  -1
7   8   4   2   1

    6   6   2
  1   5   1   1
1   2   7   8   9

   -3  -5  -5
 -2  -1  -4  -1
9   7   6   2   1

    1   2   3
  2  -1   2   0
1   3   2   4   5

   -4  -2  -3
 -2  -2   0  -3
8   6   4   4   1


TLDR is, form a quick graph w/ edge weights

*/
type edge struct {
	to *node
	weight int
}

type node struct {
	val int
	short edge
	long edge
}

func testLine(nums []int) bool {
	for i := range nums[:len(nums)-1] {
		e1 := edge{
// Create edges for from i to i+1, i+2
// Create special node for the last section
// Traverse nodes down, if all short edges are fine in length
// or if we can make a good path with ONLY one long edge
// then it's fine.
		}
		node := node{
			val: nums[i],

		}
	}
}

func checkDelta(a int, b int) bool {
	delta := a - b
	return delta >= 1 && delta <= 3
}

func mapSlice[T any, U any](arr []T, mapFunc func(T) U) (mappedArr []U) {
	mappedArr = make([]U, 0, len(arr))
	for _, a := range arr {
		mappedArr = append(mappedArr, mapFunc(a))
	}
	return
}
