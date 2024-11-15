package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func parseSeeds(rows []string) (seeds []int) {
	splitNumbers := strings.Split(rows[0], ":")

	numbers := strings.Split(splitNumbers[1], " ")

	// re := regexp.MustCompile(`\d+`)
	// idPart := re.FindString(rows[0])

	// numbers := strings.Fields(idPart)
	for _, n := range numbers {
		nInt, _ := strconv.Atoi(n)
		seeds = append(seeds, nInt)
	}
	return seeds
}

func main() {
	data := readFile("input.txt")
	rows := textToRows(data)
	seeds := parseSeeds(rows)

	fmt.Println(seeds)
}
