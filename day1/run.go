package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	numbers := [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	dat, error := os.ReadFile("./newCalibrationDocument.txt")
	if error != nil {
		panic(error)
	}
	text := string(dat)
	lines := strings.Split(text, "\n")

	var sum int64 = 0
	for _, s := range lines {
		first := ""
		last := ""

		// if digit update first and last
		for ci, c := range s {
			// digit detection
			if unicode.IsDigit(c) {
				if len(first) == 0 {
					first = string(c)
					last = first
				} else {
					last = string(c)
				}
			} else {
				for ni, num := range numbers {
					numlen := len(num)
					if len(s) < ci+numlen {
						continue
					}
					if s[ci:ci+numlen] == num {
						d := fmt.Sprint(ni + 1)
						if len(first) == 0 {
							first = d
							last = d
						} else {
							last = d
						}
						break
					}
				}
			}
		}

		if len(first) == 0 {
			continue
		}
		int, err := strconv.ParseInt(first+last, 10, 64)
		if err != nil {
			panic(err)
		}
		sum += int
	}
	fmt.Printf("sum: %v\n", sum)
}
