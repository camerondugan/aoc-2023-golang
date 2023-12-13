package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// hand Enum
const (
	HighCard     = iota
	OnePair      = iota
	TwoPair      = iota
	ThreeOfAKind = iota
	FullHouse    = iota
	FourOfAKind  = iota
	FiveOfAKind  = iota
)

type play struct {
	hand string
	bet  int
}

func p1(scanner *bufio.Scanner) {
	plays := []play{}

	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")

		bet, err := strconv.Atoi(words[1])
		check(err)
		plays = append(plays, play{
			words[0],
			bet,
		})
	}

	// find rank
	sort.Slice(plays, func(i, j int) bool {
		h1 := getHandType(plays[i].hand)
		h2 := getHandType(plays[j].hand)
		if h1 == h2 {
			for z := 0; z < len(plays[i].hand); z++ {
				if plays[i].hand[z] != plays[j].hand[z] {
					return cardValue(plays[i].hand[z]) < cardValue(plays[j].hand[z])
				}
			}
		} else {
			return h1 < h2
		}
		return true // shouldn't happen but might?
	})

	// return winnings
	winnings := 0
	for rank, p := range plays {
		winnings += p.bet * (rank + 1)
	}
	fmt.Printf("winnings: %v\n", winnings)
}

func getHandType(hand string) int {
	mainNum := 0
	secondaryNum := 0
	// part2 only vvv
	jokers := strings.Count(hand, "J")
	hand = strings.ReplaceAll(hand, "J", "")
	// part2 only ^^^
	for i, c := range hand {
		counted := false
		for j := range hand {
			if i >= j {
				continue
			}
			if c == rune(hand[j]) {
				counted = true
				break
			}
		}
		if counted {
			continue
		}
		count := strings.Count(hand, string(c))
		if count > mainNum {
			secondaryNum = mainNum
			mainNum = count
		} else if count > secondaryNum {
			secondaryNum = count
		}
	}

	// part2 only
	mainNum += jokers

	switch mainNum {
	case 5:
		return FiveOfAKind
	case 4:
		return FourOfAKind
	case 3:
		if secondaryNum == 2 {
			return FullHouse
		} else {
			return ThreeOfAKind
		}
	case 2:
		if secondaryNum == 2 {
			return TwoPair
		} else {
			return OnePair
		}
	default:
		return HighCard
	}
}

func cardValue(card byte) int {
	if card >= '0' && card <= '9' {
		return int(card) - '0'
	} else {
		switch card {
		case 'A':
			return 14
		case 'K':
			return 13
		case 'Q':
			return 12
		case 'J':
			return 1 // joker for p2
		default: // T
			return 10
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	p1(scanner)
}
