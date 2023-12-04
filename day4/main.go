package main

import (
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func p1() {
	lines, err := os.ReadFile("input.txt")
	check(err)
	sum := 0
	for _, line := range strings.Split(string(lines), "\n") {
		if len(line) == 0 {
			break
		}
		cardPoints := 0
		a := strings.Split(string(line), ": ")
		numbers := strings.ReplaceAll(a[1], "  ", " ")
		b := strings.Split(numbers, " | ")
		winners := strings.Split(b[0], " ")
		ourNums := strings.Split(b[1], " ")

		for _, num := range ourNums {
			winning := false
			for _, win := range winners {
				if num == win {
					winning = true
					break
				}
			}
			if winning {
				if cardPoints == 0 {
					cardPoints++
				} else {
					cardPoints *= 2
				}
			}
		}
		fmt.Printf("cardPoints: %v\n", cardPoints)
		sum += cardPoints
	}
	fmt.Printf("sum: %v\n", sum)
}

func p2() {
	lines, err := os.ReadFile("input.txt")
	check(err)
	numCards := 0
	copies := []int{}
	for _, line := range strings.Split(string(lines), "\n") {
		if len(line) == 0 {
			break
		}
		copiesWon := 0

		removeName := strings.Split(string(line), ": ")
		numbers := strings.ReplaceAll(removeName[1], "  ", " ")
		splitNums := strings.Split(numbers, " | ")
		winners := strings.Split(splitNums[0], " ")
		ourNums := strings.Split(splitNums[1], " ")

		for _, num := range ourNums {
			winning := false

			for _, win := range winners {
				if num == win {
					winning = true
					break
				}
			}

			if winning {
				copiesWon++
			}
		}

		numCopies := 1
		if len(copies) > 0 {
			numCopies = copies[0] + 1
			copies = copies[1:]
		}

		for i := 0; i < copiesWon; i++ {
			if i >= len(copies) {
				copies = append(copies, numCopies)
			} else {
				copies[i] += numCopies
			}
		}

		// fmt.Printf("copyTimes: %v\n", copies)
		// fmt.Printf("numCopies: %v\n", numCopies)
		numCards += numCopies
	}
	fmt.Printf("sum: %v\n", numCards)
}

func main() {
	// p1()
	p2()
}
