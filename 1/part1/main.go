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

func textToRows(data string) (rows []string) {
	reg := regexp.MustCompile(`[A-Za-z]`)
	cleanData := reg.ReplaceAllString(data, "${1}")

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

	// fmt.Println("Adding", startValue, "to", result)
	return startValue + result
}

func main() {

	data := readFile("file.txt")
	rows := textToRows(data)
	currentValue := 0

	for i := 0; i < len(rows); i++ {
		currentValue = calculate(rows[i], currentValue)
		// fmt.Println("Current Value is:", currentValue)
	}

	fmt.Println("Result is:", currentValue)
}
