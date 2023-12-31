package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	valid bool
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

func parseData(rows []string) (parsed []game) {
	for _, item := range rows {
		var current game
		var err error

		current.valid = true

		splitId := strings.Split(item, ":")

		re := regexp.MustCompile(`\d+`)
		idPart := re.FindString(splitId[0])

		current.id, err = strconv.Atoi(idPart)
		if err != nil {
			fmt.Println("Conversion error", err)
			return
		}

		splitRolls := strings.Split(splitId[1], ";")
		for _, rolls := range splitRolls {
			var currentRoll roll
			red_re := regexp.MustCompile(`red`)
			green_re := regexp.MustCompile(`green`)
			blue_re := regexp.MustCompile(`blue`)

			for _, color := range strings.Split(rolls, ",") {
				if red_re.MatchString(color) {
					num, _ := strconv.Atoi(re.FindString(color))
					if num > 12 {
						current.valid = false
					}
					currentRoll.red = num
				}
				if green_re.MatchString(color) {
					num, _ := strconv.Atoi(re.FindString(color))
					if num > 13 {
						current.valid = false
					}
					currentRoll.green = num
				}
				if blue_re.MatchString(color) {
					num, _ := strconv.Atoi(re.FindString(color))
					if num > 14 {
						current.valid = false
					}
					currentRoll.blue = num
				}
				current.rolls = append(current.rolls, currentRoll)
			}
		}
		// if current.valid {
		parsed = append(parsed, current)
		// }
	}
	return parsed
}

func main() {

	data := readFile("file.txt")
	rows := textToRows(data)

	parsedGames := parseData(rows)
	var tot int

	for _, game := range parsedGames {
		fmt.Println(game)
		if game.valid {
			tot = tot + game.id
		}
	}

	fmt.Println(tot)
}
