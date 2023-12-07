package main

import (
	"fmt"
	"os"
	"strings"
)

type roll struct {
	red   int
	green int
	blue  int
}

type game struct {
	id    int
	rolls []roll
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

func main() {

	data := readFile("file.txt")
	rows := textToRows(data)

	var output []game

	for _, item := range rows {
		var current game
		current.id = strings.Split(item, ":")
	}

	fmt.Println(rows[0])
}
