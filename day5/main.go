package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// reverse search
func p2() int {
	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)

	readSeeds := false
	seeds := []int{}
	starts := []int{}
	destinations := []int{}
	ranges := []int{}
	startsD := [][]int{}
	destinationsD := [][]int{}
	rangesD := [][]int{}

	for scanner.Scan() {
		text := scanner.Text()
		if !readSeeds {
			seedsAsStr := strings.Split(strings.Split(text, ": ")[1], " ")
			for _, seedStr := range seedsAsStr {
				seed, err := strconv.Atoi(seedStr)
				check(err)
				seeds = append(seeds, seed)
			}
			readSeeds = true
			// fmt.Printf("seeds: %v\n", seeds)
			continue
		}

		// Record maps
		if len(text) == 0 {
			if len(destinations) == 0 {
				continue
			}
			destinationsD = append(destinationsD, destinations)
			startsD = append(startsD, starts)
			rangesD = append(rangesD, ranges)
			destinations = []int{}
			starts = []int{}
			ranges = []int{}
			continue
		} else if text[len(text)-1] == ':' {
			fmt.Printf("text: %v\n", text)
			continue
		}

		// Fill maps
		numbers := strings.Split(text, " ")

		dest, err := strconv.Atoi(numbers[0])
		check(err)
		start, err := strconv.Atoi(numbers[1])
		check(err)
		r, err := strconv.Atoi(numbers[2])
		check(err)
		destinations = append(destinations, dest)
		starts = append(starts, start)
		ranges = append(ranges, r)
	}

	destinationsD = append(destinationsD, destinations)
	startsD = append(startsD, starts)
	rangesD = append(rangesD, ranges)

	found := false
	smallest := 0
	for !found {
		// fmt.Printf("smallest: %v\n", smallest)
		// find seed for matching loc
		location := smallest
		// fmt.Printf("%v\n", "startSearch")
		// fmt.Printf("location: %v\n", location)
		for depth := len(startsD) - 1; depth >= 0; depth-- {
			for i := range startsD[depth] {
				if location >= destinationsD[depth][i] && location < destinationsD[depth][i]+rangesD[depth][i] {
					location += startsD[depth][i] - destinationsD[depth][i]
					break
				}
			}
			// fmt.Printf("location: %v\n", location)
		}
		// check seeds
		seed := -1
		for _, seedRange := range seeds {
			if seed == -1 {
				seed = seedRange
			} else {
				// fmt.Printf("seedRange: %v\n", seedRange)
				if location >= seed && location < seed+seedRange {
					fmt.Printf("seed: %v\n", seed)
					fmt.Printf("seedRange: %v\n", seedRange)
					fmt.Printf("location: %v\n", location)
					fmt.Printf("smallest: %v\n", smallest)
					return smallest
				}
				seed = -1
			}
		}
		smallest++
	}

	return -1
}

// forwards search
func p1() int {
	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)

	readSeeds := false
	seeds := []int{}
	starts := []int{}
	destinations := []int{}
	ranges := []int{}

	for scanner.Scan() {
		text := scanner.Text()
		if !readSeeds {
			seedsAsStr := strings.Split(strings.Split(text, ": ")[1], " ")
			for _, seedStr := range seedsAsStr {
				seed, err := strconv.Atoi(seedStr)
				check(err)
				seeds = append(seeds, seed)
			}
			readSeeds = true
			// fmt.Printf("seeds: %v\n", seeds)
			continue
		}

		// Use then clear maps
		if len(text) == 0 {
			// Use
			newSeeds := []int{}
			for _, seed := range seeds {
				newSeed := seed
				for i := range starts {
					if seed >= starts[i] && seed < starts[i]+ranges[i] {
						newSeed += destinations[i] - starts[i]
						break
					}
				}
				newSeeds = append(newSeeds, newSeed)
			}
			seeds = newSeeds
			// fmt.Printf("seeds: %v\n", seeds)
			// Clear
			starts = []int{}
			destinations = []int{}
			ranges = []int{}
			continue
		} else if text[len(text)-1] == ':' {
			fmt.Printf("text: %v\n", text)
			continue
		}

		// Fill maps
		numbers := strings.Split(text, " ")

		dest, err := strconv.Atoi(numbers[0])
		check(err)
		start, err := strconv.Atoi(numbers[1])
		check(err)
		r, err := strconv.Atoi(numbers[2])
		check(err)
		destinations = append(destinations, dest)
		starts = append(starts, start)
		ranges = append(ranges, r)
	}

	// run once again after loop
	newSeeds := []int{}
	for _, seed := range seeds {
		newSeed := seed
		for i := range starts {
			if seed >= starts[i] && seed < starts[i]+ranges[i] {
				newSeed += destinations[i] - starts[i]
				break
			}
		}
		newSeeds = append(newSeeds, newSeed)
	}
	seeds = newSeeds
	// fmt.Printf("seeds: %v\n", seeds)
	smallest := math.MaxInt
	for _, seed := range seeds {
		if seed < smallest {
			smallest = seed
		}
	}
	// fmt.Printf("seeds: %v\n", seeds)
	fmt.Printf("smallest: %v\n", smallest)
	return smallest
}

func main() {
	p1()
	p2()
}
