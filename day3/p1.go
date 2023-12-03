package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func detect(c byte) (bool, bool) {
	number := false
	symbol := false

	if '0' <= c && '9' >= c {
		number = true
	} else if c != '.' {
		symbol = true
	}

	return number, symbol
}

func grabIntRange(line []byte, startPos int) (int, int) {
	maxPos := startPos + 1
	num, _ := detect(line[maxPos])
	for num {
		maxPos++
		if maxPos+1 > len(line) {
			break
		}
		num, _ = detect(line[maxPos])
	}
	minPos := startPos - 1
	num, _ = detect(line[minPos])
	for num {
		minPos--
		if minPos < 0 {
			break
		}
		num, _ = detect(line[minPos])
	}
	minPos++
	return minPos, maxPos
}

func validate(prev []byte, cur []byte, next []byte) int {
	fmt.Printf("prevLine: %v\n", string(prev))
	fmt.Printf("currLine: %v\n", string(cur))
	fmt.Printf("nextLine: %v\n", string(next))

	sum := 0

	if len(cur) == 0 {
		return 0
	}

	for i, char := range cur {
		num, symb := detect(char)

		if symb {
			for _, str := range [][]byte{prev, cur, next} {
				minI := i - 1
				if minI < 0 {
					minI = 0
				}

				maxI := i + 2
				if maxI > len(str) {
					maxI = len(str)
				}

				for j := range str[minI:maxI] {
					pos := minI + j
					num, _ = detect(str[pos])

					if num {
						minPos, maxPos := grabIntRange(str, pos)
						numVal, err := strconv.Atoi(string(str[minPos:maxPos]))
						if err != nil {
							panic(err)
						}
						for z := minPos; z < maxPos; z++ {
							str[z] = '.'
						}
						sum += numVal
					}
				}
			}
		}
	}
	return sum
}

func p1(filename string) int {
	file, _ := os.Open(filename)
	reader := bufio.NewReader(file)
	line, prefix, err := reader.ReadLine()
	nextLine := []byte("")
	curLine := []byte("")
	prevLine := []byte("")

	sum := 0

	for err == nil {
		for _, c := range line {
			nextLine = append(nextLine, c)
		}
		if prefix {
			continue
		}

		sum += validate(prevLine, curLine, nextLine)

		prevLine = curLine
		curLine = nextLine
		nextLine = []byte("")
		line, prefix, err = reader.ReadLine()
	}

	return sum + validate(prevLine, curLine, nextLine)
}

func main() {
	// fmt.Printf("p1(): %v\n", p1("input.txt"))
	fmt.Printf("p2(): %v\n", p2("input.txt"))
}
