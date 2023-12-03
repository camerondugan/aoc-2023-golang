package main

import (
	"bufio"
	"os"
	"strconv"
)

func validate2(prev []byte, cur []byte, next []byte) int {
	// fmt.Printf("prevLine: %v\n", string(prev))
	// fmt.Printf("currLine: %v\n", string(cur))
	// fmt.Printf("nextLine: %v\n", string(next))

	sum := 0

	if len(cur) == 0 {
		return 0
	}

	for i, char := range cur {
		num, symb := detect(char)

		if symb {
			numValues := 0
			gearRatio := 0
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
						numValues++
						minPos, maxPos := grabIntRange(str, pos)
						numVal, err := strconv.Atoi(string(str[minPos:maxPos]))
						if err != nil {
							panic(err)
						}
						for z := minPos; z < maxPos; z++ {
							str[z] = '.'
						}
						if gearRatio == 0 {
							gearRatio = numVal
						} else {
							gearRatio *= numVal
						}
					}
				}
			}
			if char == '*' && numValues == 2 {
				sum += gearRatio
			}
		}
	}
	return sum
}

func p2(filename string) int {
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

		sum += validate2(prevLine, curLine, nextLine)

		prevLine = curLine
		curLine = nextLine
		nextLine = []byte("")
		line, prefix, err = reader.ReadLine()
	}

	return sum + validate2(prevLine, curLine, nextLine)
}
