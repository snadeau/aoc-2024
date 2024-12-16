package main

import (
	"aoc-2024/common/arrayutils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open file %s: %v", f.Name(), err)
	}

	scanner := bufio.NewScanner(f)
	rulesMap := generateRulesMap(scanner)
	pages := getPages(scanner)
	partOne := solvePartOne(rulesMap, pages)
	fmt.Println("Part 1:", partOne)
	partTwo := solvePartTwo(rulesMap, pages)
	fmt.Println("Part 2:", partTwo-partOne)
}

func solvePartOne(rulesMap map[int]map[int]bool, pages [][]int) int {
	sum := 0
OuterLoop:
	for _, page := range pages {
		seen := make(map[int]bool)
		for _, i := range page {
			if rulesMap[i] != nil {
				for k := range rulesMap[i] {
					if seen[k] {
						continue OuterLoop
					}
				}
			}
			seen[i] = true
		}
		sum += page[(len(page)-1)/2]
	}
	return sum

}

func solvePartTwo(rulesMap map[int]map[int]bool, pages [][]int) int {
	sum := 0
	for _, page := range pages {
		for i := range page {
			j := i
			for j > 0 {
				if rulesMap[page[j]] != nil {
					for k := range rulesMap[page[j]] {
						if page[j-1] == k {
							temp := page[j-1]
							page[j-1] = page[j]
							page[j] = temp
							break
						}
					}
				}
				j--
			}
		}
		fmt.Println(page)
		sum += page[(len(page)-1)/2]
	}
	return sum
}

func generateRulesMap(scanner *bufio.Scanner) map[int]map[int]bool {
	rulesMap := make(map[int]map[int]bool)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		nums := arrayutils.Map(strings.Split(strings.TrimSpace(line), "|"), func(s string) int {
			num, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal("Unable to parse integer")
			}
			return num
		})
		if rulesMap[nums[0]] == nil {
			rulesMap[nums[0]] = make(map[int]bool)
		}
		rulesMap[nums[0]][nums[1]] = true
	}
	return rulesMap
}

func getPages(scanner *bufio.Scanner) [][]int {
	pages := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		page := arrayutils.Map(strings.Split(strings.TrimSpace(line), ","), func(s string) int {
			num, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal("Unable to parse integer")
			}
			return num
		})
		pages = append(pages, page)
	}
	return pages
}
