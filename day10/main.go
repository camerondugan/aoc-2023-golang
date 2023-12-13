package main

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	x, y int
}

func findStart(lines []string) (int, int) {
	startx, starty := 0, 0
SearchStart:
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == 'S' {
				startx = j
				starty = i
				break SearchStart
			}
		}
	}
	return startx, starty
}

func Adjacents(c byte, p Position) []Position {
	switch c {
	case 'S':
		{
			return []Position{{p.x - 1, p.y}, {p.x + 1, p.y}, {p.x, p.y - 1}, {p.x, p.y + 1}}
		}
	case '-':
		return []Position{{p.x - 1, p.y}, {p.x + 1, p.y}}
	case '|':
		return []Position{{p.x, p.y + 1}, {p.x, p.y - 1}}
	case 'L':
		return []Position{{p.x + 1, p.y}, {p.x, p.y - 1}}
	case '7':
		return []Position{{p.x - 1, p.y}, {p.x, p.y + 1}}
	case 'J':
		return []Position{{p.x - 1, p.y}, {p.x, p.y - 1}}
	case 'F':
		return []Position{{p.x + 1, p.y}, {p.x, p.y + 1}}
	default:
		return []Position{}
	}
}

func InBounds(lines *[]string, pos Position) bool {
	if pos.x > len((*lines)[0]) || pos.x < 0 {
		return false
	}
	if pos.y > len((*lines)) || pos.y < 0 {
		return false
	}
	return true
}

func p1(lines []string) int {
	distances := [][]int{}
	for range lines {
		distances = append(distances, make([]int, len(lines[0])))
	}
	startx, starty := findStart(lines)
	start := Position{startx, starty}
	distances[starty][startx] = 1
	startAdjs := Adjacents('S', start)
	valids := 0
	path1, path2 := Position{}, Position{}
	// remove invalid start paths
	for i := range startAdjs {
		if InBounds(&lines, startAdjs[i]) {
			adjs := Adjacents(lines[startAdjs[i].y][startAdjs[i].x], startAdjs[i])
			for _, adj := range adjs {
				if adj.x == start.x && adj.y == start.y {
					if valids == 0 {
						path1 = startAdjs[i]
						valids++
					} else {
						path2 = startAdjs[i]
					}
				}
			}
		}
	}
	paths := []Position{path1, path2}
	for _, p := range paths {
		distances[p.y][p.x] = 1
	}
	// -1 for start for counting later
	distances[start.y][start.x] = -1

	maxDist := 1
	for len(paths) != 0 {
		newPaths := []Position{}
		maxDist++
		for _, pos := range paths {
			for _, adj := range Adjacents(lines[pos.y][pos.x], pos) {
				if !InBounds(&lines, adj) {
					continue
				}
				// if new path
				if lines[adj.y][adj.x] != '.' && distances[adj.y][adj.x] == 0 {
					newPaths = append(newPaths, adj)
					distances[adj.y][adj.x] = maxDist
				}
			}
		}
		paths = newPaths
	}
	distances[starty][startx] = 0 // optional
	maxDist--
	for _, ds := range lines {
		for _, d := range ds {
			fmt.Print(string(d))
		}
		fmt.Println()
	}
	fmt.Printf("maxDist: %v\n", maxDist)
	p2(distances, lines)
	return maxDist
}

func p2(distances [][]int, lines []string) {
	flat := [][]bool{}
	for i := range distances {
		flat = append(flat, make([]bool, len(distances[i])))
	}

	vert := [][]bool{}
	for i := range distances {
		vert = append(vert, make([]bool, len(distances[i])))
	}

	ins := false // inside
	wall := false
	for i := range distances {
		for j := range distances[i] {
			nwall := distances[i][j] != 0
			if nwall {
				if !wall {
					ins = !ins
				}
				continue
			}
			flat[i][j] = ins
		}
	}
	fmt.Println("Horizontal finished")

	ins = true // inside
	wall = false
	for i := range distances[0] {
		for j := range distances {
			nwall := distances[j][i] != 0
			if nwall {
				if !wall {
					ins = !ins
				}
				continue
			}
			vert[j][i] = ins
		}
	}
	fmt.Println("Horizontal finished")

	inside := 0
	for i := range distances {
		for j := range distances[i] {
			if flat[i][j] && vert[i][j] {
				inside++
			}
		}
	}
	fmt.Printf("inside: %v\n", inside)
}

func main() {
	file, err := os.ReadFile("example.txt")
	// file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	lines = lines[0 : len(lines)-1]
	p1(lines)
}
