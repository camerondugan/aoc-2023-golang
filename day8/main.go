package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type branch struct {
	center string
	left   string
	right  string
}

func parse(s *bufio.Scanner) (string, []branch) {
	instructions := ""
	branches := []branch{}
	for s.Scan() {
		if instructions == "" {
			instructions = s.Text()
			continue
		} else if len(s.Text()) == 0 {
			continue
		}
		center, text, _ := strings.Cut(s.Text(), " = (")
		left, text, _ := strings.Cut(text, ", ")
		branches = append(branches, branch{
			center,
			left,
			text[:len(text)-1],
		})
	}
	return instructions, branches
}

func p1(s *bufio.Scanner) {
	instructions, branches := parse(s)
	start := "AAA"
	cur := &start
	steps := 0
	for *cur != "ZZZ" {
		fmt.Printf("cur: %v\n", *cur)
		for i := range branches {
			if *cur == branches[i].center {
				s := instructions[steps%len(instructions)]
				if s == 'L' {
					cur = &branches[i].left
					break
				} else {
					cur = &branches[i].right
					break
				}
			}
		}
		steps++
	}
	fmt.Printf("branches: %v\n", branches)
	fmt.Printf("steps: %v\n", steps)
}

func p2(s *bufio.Scanner) {
	instructions, branches := parse(s)
	centers := []*string{}
	for i := range branches {
		if branches[i].center[len(branches[i].center)-1] == 'A' {
			fmt.Printf("b: %v\n", branches[i].center)
			centers = append(centers, &branches[i].center)
		}
	}
	steps := 0
	complete := false
	for !complete {
		complete = true
		for i := range centers {
			for b := range branches {
				if *centers[i] == branches[b].center {
					s := instructions[steps%len(instructions)]
					if s == 'L' {
						centers[i] = &branches[b].left
						break
					} else {
						centers[i] = &branches[b].right
						break
					}
				}
			}
			if (*centers[i])[len(*centers[i])-1] != 'Z' {
				complete = false
			}
		}
		steps++
		if steps%100000 == 0 {
			fmt.Printf("steps: %v\n", steps)
		}
	}
	fmt.Printf("steps: %v\n", steps)
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	s := bufio.NewScanner(file)
	// p1(s)
	file, err = os.Open("input.txt")
	check(err)
	s = bufio.NewScanner(file)
	p2(s)
}
