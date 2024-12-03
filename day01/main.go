package main

import (
	"aoc-2024/common/mathutils"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		nums := strings.Fields(lines[i])
		l, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatal(err)
		}
		left[i] = l
		r, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatal(err)
		}
		right[i] = r
	}
	fmt.Println("Distance:", solveDistance(left, right))
	fmt.Println("Similarity:", solveSimilarity(left, right))
}

func solveDistance(left []int, right []int) int {
	sort.Ints(left)
	sort.Ints(right)
	distance := 0
	for i := 0; i < len(left); i++ {
		distance += mathutils.Abs(left[i] - right[i])
	}
	return distance
}

func solveSimilarity(left []int, right []int) int {
	similarity := 0
	freqMap := getFrequencyMap(right)
	for _, value := range left {
		similarity += freqMap[value] * value
	}
	return similarity
}

func getFrequencyMap(values []int) map[int]int {
	freq := make(map[int]int)
	for _, value := range values {
		freq[value] += 1
	}
	return freq
}
