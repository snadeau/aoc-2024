package main

import (
	"aoc-2024/common/arrayutils"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := arrayutils.Map(strings.Split(strings.TrimSpace(string(data)), "\n"), func(s string) string {
		return strings.TrimSpace(s)
	})
	xmas := 0
	xmasAlt := 0
	for row, line := range lines {
		for column := range line {
			xmas += searchUp(lines, row, column)
			xmas += searchUpRight(lines, row, column)
			xmas += searchRight(lines, row, column)
			xmas += searchDownRight(lines, row, column)
			xmas += searchDown(lines, row, column)
			xmas += searchDownLeft(lines, row, column)
			xmas += searchLeft(lines, row, column)
			xmas += searchUpLeft(lines, row, column)
			xmasAlt += searchXmas(lines, row, column)
		}
	}

	fmt.Println("XMAS occurences:", xmas)
	fmt.Println("X-MAS occurences:", xmasAlt)
}

func checkBounds(lines []string, row int, column int) bool {
	return row > -1 && row < len(lines) && column > -1 && column < len(lines[0])
}

func shouldContinue(lines []string, row int, column int, current string) bool {
	return checkBounds(lines, row, column) && current != "XMAS" && strings.HasPrefix("XMAS", current)
}

func searchXmas(lines []string, row int, column int) int {
	if row-1 < 0 || row+1 >= len(lines) || column-1 < 0 || column+1 >= len(lines[0]) || string(lines[row][column]) != "A" {
		return 0
	}
	diag1 := string(lines[row-1][column-1]) + "A" + string(lines[row+1][column+1])
	diag2 := string(lines[row-1][column+1]) + "A" + string(lines[row+1][column-1])
	if (diag1 == "MAS" || diag1 == "SAM") && (diag2 == "MAS" || diag2 == "SAM") {
		return 1
	}
	return 0
}

func searchUp(lines []string, row int, column int) int {
	current := ""
	for ok := true; ok; ok = shouldContinue(lines, row, column, current) {
		current += string(lines[row][column])
		row--
	}
	if current == "XMAS" {
		return 1
	}
	return 0
}

func searchUpRight(lines []string, row int, column int) int {
	current := ""
	for ok := true; ok; ok = shouldContinue(lines, row, column, current) {
		current += string(lines[row][column])
		row--
		column++
	}
	if current == "XMAS" {
		return 1
	}
	return 0
}

func searchRight(lines []string, row int, column int) int {
	current := ""
	for ok := true; ok; ok = shouldContinue(lines, row, column, current) {
		current += string(lines[row][column])
		column++
	}
	if current == "XMAS" {
		return 1
	}
	return 0
}

func searchDownRight(lines []string, row int, column int) int {
	current := ""
	for ok := true; ok; ok = shouldContinue(lines, row, column, current) {
		current += string(lines[row][column])
		row++
		column++
	}
	if current == "XMAS" {
		return 1
	}
	return 0
}

func searchDown(lines []string, row int, column int) int {
	current := ""
	for ok := true; ok; ok = shouldContinue(lines, row, column, current) {
		current += string(lines[row][column])
		row++
	}
	if current == "XMAS" {
		return 1
	}
	return 0
}

func searchDownLeft(lines []string, row int, column int) int {
	current := ""
	for ok := true; ok; ok = shouldContinue(lines, row, column, current) {
		current += string(lines[row][column])
		row++
		column--
	}
	if current == "XMAS" {
		return 1
	}
	return 0
}

func searchLeft(lines []string, row int, column int) int {
	current := ""
	for ok := true; ok; ok = shouldContinue(lines, row, column, current) {
		current += string(lines[row][column])
		column--
	}
	if current == "XMAS" {
		return 1
	}
	return 0
}

func searchUpLeft(lines []string, row int, column int) int {
	current := ""
	for ok := true; ok; ok = shouldContinue(lines, row, column, current) {
		current += string(lines[row][column])
		row--
		column--
	}
	if current == "XMAS" {
		return 1
	}
	return 0
}
