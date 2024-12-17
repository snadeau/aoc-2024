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

type Equation struct {
	Total    int
	Operands []int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open file %s: %v", f.Name(), err)
	}

	scanner := bufio.NewScanner(f)
	equations := parseInput(scanner)
	fmt.Println("Part 1:", solve(equations, []rune{'+', '*'}))
	fmt.Println("Part 2:", solve(equations, []rune{'+', '*', '|'}))
}

func evaluate(equation Equation, operators []rune) int {
	curValue := equation.Operands[0]
	operandIdx := 1
	operatorIdx := 0
	for operandIdx < len(equation.Operands) {
		if operators[operatorIdx] == '*' {
			curValue *= equation.Operands[operandIdx]
		} else if operators[operatorIdx] == '+' {
			curValue += equation.Operands[operandIdx]
		} else if operators[operatorIdx] == '|' {
			curValStr := strconv.Itoa(curValue)
			thisValStr := strconv.Itoa(equation.Operands[operandIdx])
			curValue = parseInt(curValStr + thisValStr)
		}
		operandIdx++
		operatorIdx++
	}
	return curValue
}

func recursePossibleOperators(equation Equation, possibleOperators []rune, length int, current []rune, found *bool) {
	if *found {
		return
	}
	if len(current) == length {
		*found = evaluate(equation, current) == equation.Total
		return
	}
	for _, op := range possibleOperators {
		current = append(current, op)
		recursePossibleOperators(equation, possibleOperators, length, current, found)
		current = current[:len(current)-1]
	}
}

func canBeSolved(equation Equation, possibleOperators []rune) bool {
	length := len(equation.Operands) - 1
	found := false
	recursePossibleOperators(equation, possibleOperators, length, []rune{}, &found)
	return found
}

func solve(equations []Equation, possibleOperators []rune) int {
	sum := 0
	for _, equation := range equations {
		if canBeSolved(equation, possibleOperators) {
			sum += equation.Total
		}
	}
	return sum
}

func parseInt(s string) int {
	parsed, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Unable to parse int")
	}
	return parsed
}

func parseInput(scanner *bufio.Scanner) []Equation {
	input := make([]Equation, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(strings.TrimSpace(line), ":")
		total := parseInt(parts[0])
		operands := arrayutils.Map(strings.Fields(parts[1]), func(s string) int {
			return parseInt(s)
		})
		equation := Equation{
			Total:    total,
			Operands: operands,
		}
		input = append(input, equation)
	}
	return input
}
