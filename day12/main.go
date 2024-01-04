package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	prmt "github.com/gitchander/permutation"
)

// hot springs
func main() {
	// file, _ := os.Open("example.txt")
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	sum := 0
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
		springs = strings.Trim(springs, ".")
		// define num # and .
		totalChars := len(springs)

		// find anwser
		// exactlyOneLen calc
		exactlyOneLen := 0
		for _, conf := range config {
			exactlyOneLen++
			exactlyOneLen += conf
		}
		if exactlyOneLen > 0 {
			exactlyOneLen--
		}

		if totalChars == exactlyOneLen {
			sum++
			continue
		} else if totalChars < exactlyOneLen {
			continue
		}
		// changeable string
		springsOpt := []rune(springs)

		// brute force
		numWorking, numBroken := 0, 0
		for _, num := range config {
			numBroken += num
		}
		numWorking = len(springsOpt) - numBroken

		for _, c := range springsOpt {
			if c == '.' {
				numWorking--
			} else if c == '#' {
				numBroken--
			}
		}
		arrangement := []string{}
		for i := 0; i < numWorking; i++ {
			arrangement = append(arrangement, ".")
		}
		for _, c := range config {
			arrangement = append(arrangement, strings.Repeat("#", c))
		}
		p := prmt.New(prmt.StringSlice(arrangement))
		arrangementsChecked := []string{}
		fmt.Printf("springsOpt: %v\n", string(springsOpt))
		fmt.Printf("config: %v\n", config)
		// fmt.Printf("arrangement: %v\n", arrangement)

	nextArr:
		for p.Next() {
			// skip stupid arrangements
			noOneA := []int{}
			for _, a := range arrangement {
				if len(a) != 1 || a == "#" {
					noOneA = append(noOneA, len(a))
				}
			}
			// ensure non-1 len arrangements have good order
			minLen := len(noOneA)
			if len(config) < minLen {
				minLen = len(config)
			}
			for i := 0; i < minLen; i++ {
				if config[i] != noOneA[i] {
					continue nextArr
				}
			}

			thisStr := make([]rune, len(springsOpt)) // copy
			copy(thisStr, springsOpt)
			// fmt.Printf("arrangement: %v\n", arrangement)
			// replace arangement into optimized string
			broken := []int{}
			fixed := []int{}
			for i := range thisStr {
				if thisStr[i] == '#' {
					thisStr[i] = '?'
					broken = append(broken, i)
				} else if thisStr[i] == '.' {
					fixed = append(fixed, i)
				}
			}
			for _, s := range arrangement {
				for _, sc := range s {
					for i := range thisStr {
						if thisStr[i] == '?' {
							thisStr[i] = sc
							break
						}
					}
				}
			}
			for i := range thisStr {
				if thisStr[i] == '?' {
					thisStr[i] = '.'
				}
			}
			// fmt.Printf("thisStr: %v\n", string(thisStr))

			// fmt.Printf("thisStr: %v\n", string(thisStr))
			finalString := string(thisStr)
			// skip if seen
			for _, s := range arrangementsChecked {
				if strings.Compare(s, finalString) == 0 {
					continue nextArr
				}
			}
			// check new
			if valid(finalString, config, broken, fixed) {
				fmt.Printf("arrangement: %v\n", finalString)
				sum++
			}
			// remember
			arrangementsChecked = append(arrangementsChecked, finalString)
		}
	}
	fmt.Printf("sum: %v\n", sum)
}

func valid(arrangement string, config []int, broken []int, fixed []int) bool {
	count := 0
	for _, c := range arrangement {
		if c == '?' {
			panic("Didn't replace all ? properly")
		}
		if c == '#' {
			count++
		} else {
			if count != 0 {
				if len(config) == 0 {
					return false
				}
				if config[0] != count {
					return false
				}
				config = config[1:]
				count = 0
			}
		}
	}
	// one last iter bc might end in #
	if count != 0 {
		if len(config) == 0 {
			return false
		}
		if config[0] != count {
			return false
		}
		config = config[1:]
	}
	for _, i := range broken {
		if arrangement[i] != '#' {
			fmt.Printf("i: %v\n", i)
			fmt.Printf("arrangement: %v\n", arrangement)
			return false
		}
	}
	for _, i := range fixed {
		if arrangement[i] != '.' {
			fmt.Printf("i: %v\n", i)
			fmt.Printf("arrangement: %v\n", arrangement)
			return false
		}
	}
	return len(config) == 0
}
