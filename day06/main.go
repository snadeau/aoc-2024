package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Location struct {
	X int
	Y int
	D int
}

var directions = []Location{
	{X: 0, Y: -1}, // UP
	{X: 1, Y: 0},  // RIGHT
	{X: 0, Y: 1},  // DOWN
	{X: -1, Y: 0}, // LEFT
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open file %s: %v", f.Name(), err)
	}

	scanner := bufio.NewScanner(f)
	obstacles, guardLoc, rows, columns := getInitialData(scanner)
	path := solvePartOne(obstacles, guardLoc, rows, columns)
	solvePartTwo(obstacles, path, guardLoc, rows, columns)
}

func getNextDirection(currentDirection int) int {
	nextDirection := currentDirection + 1
	if nextDirection >= len(directions) {
		nextDirection = 0
	}
	return nextDirection
}

func solvePartOne(obstacles map[Location]bool, initialLocation Location, rows int, columns int) map[Location]bool {
	seenLocations := make(map[Location]bool)
	x := initialLocation.X
	y := initialLocation.Y
	currentDirection := 0
	for x < columns && x >= 0 && y < rows && y >= 0 {
		p := Location{X: x, Y: y}
		if obstacles[p] {
			x -= directions[currentDirection].X
			y -= directions[currentDirection].Y
			currentDirection = getNextDirection(currentDirection)
		} else if !seenLocations[p] {
			seenLocations[p] = true
		}
		x += directions[currentDirection].X
		y += directions[currentDirection].Y
	}
	fmt.Println("Part 1:", len(seenLocations))
	return seenLocations
}

func hasLoop(obstacles map[Location]bool, initialLocation Location, rows int, columns int) bool {
	seenLocations := make(map[Location]bool)
	x := initialLocation.X
	y := initialLocation.Y
	currentDirection := 0
	for x < columns && x >= 0 && y < rows && y >= 0 {
		p := Location{X: x, Y: y, D: currentDirection}
		// if we hit the same location going the same direction there is a loop
		if seenLocations[p] {
			return true
		}
		if obstacles[Location{X: x, Y: y}] {
			x -= directions[currentDirection].X
			y -= directions[currentDirection].Y
			currentDirection = getNextDirection(currentDirection)
		} else {
			seenLocations[p] = true
		}
		x += directions[currentDirection].X
		y += directions[currentDirection].Y
	}
	return false
}

func solvePartTwo(obstacles map[Location]bool, originalPath map[Location]bool, initialLocation Location, rows int, columns int) {
	sum := 0
	for o := range originalPath {
		if o.X == initialLocation.X && o.Y == initialLocation.Y {
			continue
		}
		obstacles[o] = true
		if hasLoop(obstacles, initialLocation, rows, columns) {
			sum += 1
		}
		obstacles[o] = false
	}
	fmt.Println("Part 2:", sum)
}

func getInitialData(scanner *bufio.Scanner) (map[Location]bool, Location, int, int) {
	obstacles := make(map[Location]bool)
	var guardLoc Location
	lineNum := 0
	columns := 0
	for scanner.Scan() {
		line := scanner.Text()
		for pos, char := range line {
			if char == '#' {
				p := Location{X: pos, Y: lineNum}
				obstacles[p] = true
			}
			if char == '^' {
				guardLoc = Location{X: pos, Y: lineNum, D: 0}
			}
			if pos+1 > columns {
				columns = pos + 1
			}
		}
		lineNum++
	}
	return obstacles, guardLoc, lineNum, columns
}
