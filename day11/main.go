package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func p1(file *os.File) {
	scanner := bufio.NewScanner(file)
	observation := []string{}
	var emptyColumns []bool
	for scanner.Scan() {
		// read in line
		observation = append(observation, scanner.Text())
		// dup if empty line
		if !strings.ContainsRune(scanner.Text(), '#') {
			observation = append(observation, scanner.Text())
		}
		// detect col
		if emptyColumns == nil {
			emptyColumns = make([]bool, len(scanner.Text()))
			for i := range emptyColumns {
				emptyColumns[i] = true
			}
		}
		for i, c := range scanner.Text() {
			if c == '#' {
				emptyColumns[i] = false
			}
		}
	}
	// expand column
	for i := range observation {
		newObs := strings.Builder{}
		for j := range observation[i] {
			if emptyColumns[j] {
				err := newObs.WriteByte(observation[i][j])
				if err != nil {
					panic(err)
				}
			}
			err := newObs.WriteByte(observation[i][j])
			if err != nil {
				panic(err)
			}
		}
		observation[i] = newObs.String()
	}
	for _, obs := range observation {
		fmt.Println(obs)
	}
	sumPairs(observation)
}

type galaxy struct {
	x, y int
}

func sumPairs(observation []string) {
	galaxies := []galaxy{}
	for i, obs := range observation {
		for j, o := range obs {
			if o == '#' {
				galaxies = append(galaxies, galaxy{j, i})
			}
		}
	}

	sum := 0
	for i, g1 := range galaxies {
		for j, g2 := range galaxies {
			if i >= j { // only unique combinations
				continue
			}
			dx := math.Abs((float64)(g1.x - g2.x))
			dy := math.Abs((float64)(g1.y - g2.y))
			sum += (int)(dx + dy)
		}
	}
	fmt.Printf("sum: %v\n", sum)
}

func main() {
	file, _ := os.Open("input.txt")
	p1(file)
}
