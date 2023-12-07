package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, e := os.Open("input.txt")
	check(e)
	scanner := bufio.NewScanner(file)

	times := []int{}
	distances := []int{}
	addToDistance := false

	// populate times + distances
	p1 := false
	for scanner.Scan() {
		line := scanner.Text()
		_, line, _ = strings.Cut(line, ":")
		if p1 {
			numbers := strings.Split(line, " ")
			for _, num := range numbers {
				if len(num) > 0 {
					i, e := strconv.Atoi(num)
					check(e)
					if addToDistance {
						distances = append(distances, i)
					} else {
						times = append(times, i)
					}
				}
			}
		} else {
			number := strings.ReplaceAll(line, " ", "")
			i, e := strconv.Atoi(number)
			check(e)
			if addToDistance {
				distances = append(distances, i)
			} else {
				times = append(times, i)
			}
		}
		addToDistance = true
	}

	// calc answer
	ans := 1
	for i := range times {
		wins := 0
		time := times[i]
		dist := distances[i]
		wasWinning := false

		for t := 0; t < time; t++ {
			if t*(time-t) > dist {
				wins++
				wasWinning = true
			} else if wasWinning { // optimize by stop early
				break
			}
		}
		ans *= wins
	}
	fmt.Printf("ans: %v\n", ans)
}
