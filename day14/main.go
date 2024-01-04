package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")
	s := bufio.NewScanner(f)
	var platform [][]rune

	for s.Scan() {
		fmt.Println(s.Text())
		platform = append(platform, []rune(s.Text()))
	}

	for i := range platform {
		for j := range platform[i] {
			if platform[i][j] == '.' {
				// find rolly rock
				for rp := i + 1; rp < len(platform); rp++ {
					if platform[rp][j] == 'O' {
						// pull rock to position
						platform[i][j], platform[rp][j] = platform[rp][j], platform[i][j]
						break
					} else if platform[rp][j] == '#' {
						break
					}
				}
			}
		}
	}
	sum := 0
	for i, row := range platform {
		fmt.Printf("%v\n", string(row))
		load := len(platform) - i
		circles := 0
		for _, r := range row {
			if r == 'O' {
				circles++
			}
		}
		sum += load * circles
	}
	fmt.Printf("sum: %v\n", sum)
}
