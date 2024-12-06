package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var isReal bool = true

func main() {
	filePath := "/home/dev/advent-of-code/2024/02/test"
	if isReal {
		filePath = "/home/dev/advent-of-code/2024/02/real"
	}
	var scanner bufio.Scanner
	if file, err := os.Open(filePath); err != nil {
		panic(err)
	} else {
		scanner = *bufio.NewScanner(file)
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
		fmt.Println(nums)
		if testLine(nums, true) || testLine(nums[1:], false) {
			fmt.Println("Passed")
			realValue += 1
		} else {
			fmt.Println("Failed")
		}
	}
	fmt.Println("FoundValue:", realValue)
}

func testLine(nums []int, allowSkip bool) bool {
	fmt.Print("\t")
	if allowSkip {
		fmt.Print("0")
	} else {
		fmt.Print("1")
	}
	fmt.Println(" Ascending")
	if testBounds(nums, 1, 3, allowSkip) {
		return true
	}
	fmt.Print("\t")
	if allowSkip {
		fmt.Print("0")
	} else {
		fmt.Print("1")
	}
	fmt.Println(" Descending")
	if testBounds(nums, -3, -1, allowSkip) {
		return true
	}
	return false
}

func testBounds(nums []int, low int, high int, allowSkip bool) bool {
	prev := nums[0]
	for _, curr := range nums[1:] {
		delta := curr - prev
		fmt.Println("\t\t", prev, curr, low, delta, high, allowSkip)
		if delta >= low && delta <= high {
			fmt.Println(" next")
			prev = curr
		} else if allowSkip {
			fmt.Println(" skipping")
			allowSkip = false
		} else {
			fmt.Println(" Failed")
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
