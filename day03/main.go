package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	solvePartOne(lines)
	solvePartTwo(lines)
}

func parseAndMultiply(match []string) int {
	l, err := strconv.Atoi(match[1])
	if err != nil {
		log.Fatal(err)
	}
	r, err := strconv.Atoi(match[2])
	if err != nil {
		log.Fatal(err)
	}
	return l * r
}

func solvePartOne(lines []string) {
	sum := 0
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			sum += parseAndMultiply(match)
		}
	}
	fmt.Println("Total part one:", sum)
}

func solvePartTwo(lines []string) {
	sum := 0
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|(don't\(\))|(do\(\))`)
	enabled := true
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if strings.HasPrefix(match[0], "mul") && enabled {
				sum += parseAndMultiply(match)
			} else {
				enabled = match[0] == "do()"
			}
		}
	}
	fmt.Println("Total part two:", sum)
}
