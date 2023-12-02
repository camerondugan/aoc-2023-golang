package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var validColors = make(map[string]int)

func valid(numThenColor []string) bool {
	// fmt.Printf("numThenColor: %v\n", numThenColor)
	integer, err := strconv.ParseInt(numThenColor[0], 10, 64)
	if err != nil {
		return true
	}
	// fmt.Printf("numThenColor: %v\n", numThenColor)
	return int(integer) <= validColors[numThenColor[1]]
}

func main() {
	// init map
	validColors["red"] = 12
	validColors["green"] = 13
	validColors["blue"] = 14

	file, fileerr := os.Open("input1.txt")
	check(fileerr)

	reader := bufio.NewReader(file)

	fullLine := ""

	line, truncated, err := reader.ReadLine()
	check(err)

	sum := 0
	lineNum := 0
	for err == nil {
		fullLine += string(line)
		if truncated {
			continue
		}
		lineNum++

		fullLine = strings.Split(fullLine, ": ")[1]
		// fmt.Printf("fullLine: %v\n", fullLine)

		games := strings.Split(fullLine, "; ")
		allValid := true
		for _, game := range games {
			rounds := strings.Split(game, ", ")
			for _, round := range rounds {
				numThenColor := strings.Split(round, " ")
				if !valid(numThenColor) {
					allValid = false
					break
				}
			}
			if !allValid {
				break
			}
		}
		if allValid {
			sum += lineNum
		}

		fullLine = ""
		line, truncated, err = reader.ReadLine()
	}
	fmt.Printf("sum: %v\n", sum)
}
