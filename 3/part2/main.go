package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type part struct {
	row   int
	coord []int
	value int
	valid bool
}

type gear struct {
	row    int
	coord  []int
	values []int
}

func readFile(path string) (data string) {
	input, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	data = string(input)
	return
}

func textToRows(data string) (rows []string) {
	rows = strings.Split(data, "\n")
	return rows
}

func parseNumbers(rows []string) (parts []part) {
	for i, row := range rows {
		re := regexp.MustCompile(`\d+`)
		for _, match := range re.FindAllStringIndex(row, -1) {
			var current part
			current.coord = match
			current.value, _ = strconv.Atoi(row[match[0]:match[1]])
			current.row = i

			parts = append(parts, current)
		}
	}
	return parts
}

func parseGears(rows []string) (gears []gear) {
	for i, row := range rows {
		re := regexp.MustCompile(`\*`)
		for _, match := range re.FindAllStringIndex(row, -1) {
			var current gear
			current.coord = match
			current.row = i

			gears = append(gears, current)
		}
	}
	return gears
}

func matchParts(rows []string, parts []part) (matchedParts []part) {
	for _, part := range parts {
		var rowAbove int
		var rowBelow int
		var firstChar int
		var lastChar int
		var currentRows []string

		if part.coord[0] == 0 {
			firstChar = 0
		} else {
			firstChar = part.coord[0] - 1
		}

		if part.coord[1] == len(rows[part.row]) {
			lastChar = len(rows[part.row])
		} else {
			lastChar = part.coord[1] + 1
		}

		if part.row > 0 {
			rowAbove = part.row - 1
			currentRows = append(currentRows, rows[rowAbove][firstChar:lastChar])
		}

		currentRows = append(currentRows, rows[part.row][firstChar:lastChar])

		if part.row+1 < len(rows) {
			rowBelow = part.row + 1
			currentRows = append(currentRows, rows[rowBelow][firstChar:lastChar])
		}

		for _, rowData := range currentRows {
			re := regexp.MustCompile(`[^0-9\.]`)
			if re.MatchString(rowData) {
				part.valid = true
			}
		}
		matchedParts = append(matchedParts, part)
	}
	return
}

func main() {
	var totalResult int

	data := readFile("file.txt")
	rows := textToRows(data)
	allParts := parseNumbers(rows)
	matchedParts := matchParts(rows, allParts)

	for _, gear := range parseGears(rows) {
		var gearMatch []int
		var searchRows []int
		if gear.row > 0 {
			searchRows = append(searchRows, gear.row-1)
		}

		searchRows = append(searchRows, gear.row)

		if gear.row+1 < len(rows) {
			searchRows = append(searchRows, gear.row+1)
		}

		for _, part := range matchedParts {
			for _, row := range searchRows {
				if part.row == row {
					startCoord := gear.coord[0]
					endCoord := gear.coord[1]
					if startCoord >= part.coord[0] && startCoord <= part.coord[1] {
						gearMatch = append(gearMatch, part.value)
					} else if endCoord >= part.coord[0] && endCoord <= part.coord[1] {
						gearMatch = append(gearMatch, part.value)
					}
				}
			}
		}
		if len(gearMatch) == 2 {
			gearMultiply := gearMatch[0] * gearMatch[1]
			totalResult = totalResult + gearMultiply
		}
		gear.values = gearMatch
	}

	fmt.Println("Total is:", totalResult)
}
