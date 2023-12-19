package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// hot springs
func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// parse
		springs, configStr, found := strings.Cut(scanner.Text(), " ")
		if !found {
			continue
		}
		// config
		confStrs := strings.Split(configStr, ",")
		config := []int{}
		for _, str := range confStrs {
			i, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			config = append(config, i)
		}
		// remove extra .
		springs = strings.ReplaceAll(springs, "..", ".")
		springs = strings.ReplaceAll(springs, "..", ".")
		fmt.Printf("config: %v\n", config)
		fmt.Printf("springs: %v\n", springs)
		// define num # and .
		totalChars := len(springs)
		numWorking, numBroken := 0, 0
		for _, num := range config {
			numBroken += num
		}

		// remove existing
		numWorking = totalChars - numBroken
		for _, c := range springs {
			if c == '.' {
				numWorking--
			} else if c == '#' {
				numBroken--
			}
		}
		// status update
		numUnknown := strings.Count(springs, "?")
		fmt.Printf("numWorking: %v\n", numWorking)
		fmt.Printf("numBroken: %v\n", numBroken)
		fmt.Println(numUnknown)
		if (numWorking == 0 && numBroken == numUnknown) ||
			(numBroken == 0 && numWorking == numUnknown) {
			fmt.Printf("answer: %v\n", 1)
		}
	}
}
