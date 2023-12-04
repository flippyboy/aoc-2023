package main

import (
	"fmt"
	"os"
	"regexp"
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

func textToNumbers(input string) string {
	numbers := [9][2]string{
		{"one", "1"},
		{"two", "2"},
		{"three", "3"},
		{"four", "4"},
		{"five", "5"},
		{"six", "6"},
		{"seven", "7"},
		{"eight", "8"},
		{"nine", "9"},
	}
	replaced := input
	for i := 0; i < 9; i++ {
		replaced = strings.Replace(replaced, numbers[i][0], numbers[i][0]+numbers[i][1]+numbers[i][0], -1)
	}
	return replaced
}

func textToRows(data string) (rows []string) {
	dataWithNumbers := textToNumbers(data)
	reg := regexp.MustCompile(`[A-Za-z]`)
	cleanData := reg.ReplaceAllString(dataWithNumbers, "${1}")

	rows = strings.Fields(cleanData)
	return rows
}

func calculate(row string, startValue int) (endValue int) {
	length := len(row) - 1
	values := strings.Split(row, "")

	a := values[0]
	z := values[length]

	result, err := strconv.Atoi(a + z)
	if err != nil {
		fmt.Println("Conversion error", err)
		return
	}

	fmt.Println("Adding", startValue, "to", result)
	return startValue + result
}

func main() {

	data := readFile("file.txt")
	// fmt.Println(textToNumbers(data))
	rows := textToRows(data)
	currentValue := 0

	for i := 0; i < len(rows); i++ {
		currentValue = calculate(rows[i], currentValue)
		fmt.Println("Current Value is:", currentValue)
	}

	fmt.Println("Result is:", currentValue)
}
