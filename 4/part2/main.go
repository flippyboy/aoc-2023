package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type card struct {
	id           int
	numbers      []int
	winning      []int
	points       int
	scratchCards int
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

func parseData(rows []string) (parsed []card) {
	for _, row := range rows {
		var current card

		splitId := strings.Split(row, ":")

		re := regexp.MustCompile(`\d+`)
		idPart := re.FindString(splitId[0])

		current.id, _ = strconv.Atoi(idPart)
		splitNumbers := strings.Split(splitId[1], "|")

		numbers := strings.Fields(splitNumbers[0])
		for _, n := range numbers {
			nInt, _ := strconv.Atoi(n)
			current.numbers = append(current.numbers, nInt)
		}

		winning := strings.Fields(splitNumbers[1])
		for _, n := range winning {
			nInt, _ := strconv.Atoi(n)
			current.winning = append(current.winning, nInt)
		}
		current.scratchCards = 1
		parsed = append(parsed, current)
	}
	return parsed
}

func parseWins(cards []card) (winCards []card) {
	for _, c := range cards {
		var numWins int
		for _, n := range c.numbers {
			for _, w := range c.winning {
				if n == w {
					numWins++
				}
			}
		}
		c.points = numWins
		winCards = append(winCards, c)
	}
	return winCards
}

func main() {
	data := readFile("input.txt")
	rows := textToRows(data)

	parsedGames := parseData(rows)
	parsedWins := parseWins(parsedGames)
	var totCards int

	for c := 0; c < len(parsedWins); c++ {
		x := parsedWins[c]
		for i := x.id; i < x.id+x.points; i++ {
			parsedWins[i].scratchCards = parsedWins[i].scratchCards + parsedWins[c].scratchCards
		}
		totCards = totCards + parsedWins[c].scratchCards
	}

	fmt.Println(totCards)
}
