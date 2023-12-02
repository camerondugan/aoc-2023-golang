package main

import (
	"bufio"
	"fmt"
	"math"
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
	integer, err := strconv.Atoi(numThenColor[0])
	if err != nil {
		return true
	}
	fmt.Printf("numThenColor: %v\n", numThenColor)
	return integer <= validColors[numThenColor[1]]
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
		fmt.Printf("fullLine: %v\n", fullLine)

		games := strings.Split(fullLine, "; ")
		maxRed := math.MinInt16
		maxBlue := math.MinInt16
		maxGreen := math.MinInt16
		for _, game := range games {
			rounds := strings.Split(game, ", ")
			for _, round := range rounds {
				numThenColor := strings.Split(round, " ")
				numBlocks, err := strconv.Atoi(numThenColor[0])
				check(err)
				switch numThenColor[1] {
				case "red":
					if numBlocks > maxRed {
						maxRed = numBlocks
					}
				case "blue":
					if numBlocks > maxBlue {
						maxBlue = numBlocks
					}
				case "green":
					if numBlocks > maxGreen {
						maxGreen = numBlocks
					}
				}
			}
		}
		fmt.Printf("maxRed: %v\n", maxRed)
		fmt.Printf("maxBlue: %v\n", maxBlue)
		fmt.Printf("maxGreen: %v\n", maxGreen)
		sum += maxRed * maxGreen * maxBlue

		fullLine = ""
		line, truncated, err = reader.ReadLine()
		fmt.Printf("sum: %v\n", sum)
	}
}
