package main

import (
	"aoc-2024/common/arrayutils"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	reports := strings.Split(strings.TrimSpace(string(data)), "\n")
	numSafe := 0
	numSafeWithDampener := 0
	for _, report := range reports {
		sequence := arrayutils.Map(strings.Fields(report), func(numStr string) int {
			num, _ := strconv.Atoi(numStr)
			return num
		})
		if isSafe(sequence) {
			numSafe++
		}
		if isSafeWithDampener(sequence) {
			numSafeWithDampener++
		}
	}
	fmt.Println("Safe reports:", numSafe)
	fmt.Println("Safe reports with dampener:", numSafeWithDampener)
}

func isSafe(sequence []int) bool {
	decrementing := sequence[0] > sequence[1]
	for i := 1; i < len(sequence); i++ {
		diff := sequence[i] - sequence[i-1]
		if decrementing {
			if diff > -1 || diff < -3 {
				return false
			}
		} else {
			if diff < 1 || diff > 3 {
				return false
			}
		}
	}
	return true
}

func isSafeWithDampener(sequence []int) bool {
	if isSafe(sequence) {
		return true
	}
	for i, _ := range sequence {
		clone := slices.Clone(sequence)
		dampened := append(clone[:i], clone[i+1:]...)
		if isSafe(dampened) {
			return true
		}
	}
	return false
}
