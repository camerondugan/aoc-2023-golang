package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func abssum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += int(math.Abs(float64(n)))
	}
	return sum
}

func parse2() {
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	answer := 0
	for scanner.Scan() {
		numberStings := strings.Split(scanner.Text(), " ")
		numbers := []int{}
		for i := range numberStings {
			num, _ := strconv.Atoi(numberStings[i])
			numbers = append(numbers, num)
		}

		// reverse numbers for p2
		for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}

		if len(numbers) == 0 {
			continue
		}
		total := 0
		stackNums := [][]int{}
		stackNums = append(stackNums, numbers)
		for abssum(stackNums[len(stackNums)-1]) != 0 {
			tmpNumbers := []int{}
			prevNumbers := stackNums[len(stackNums)-1]
			lastNum := prevNumbers[0]
			for j := range prevNumbers[1:] {
				j++ // i should start at 1
				tmpNumbers = append(tmpNumbers, prevNumbers[j]-lastNum)
				lastNum = prevNumbers[j]
			}
			total += lastNum
			stackNums = append(stackNums, tmpNumbers)
		}
		numbers = append(numbers, total)
		// fmt.Printf("stackNums: %v\n", stackNums)
		fmt.Printf("total: %v\n", total)
		answer += total
	}
	fmt.Printf("answer: %v\n", answer)
}

func parse1() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	answer := 0
	for scanner.Scan() {
		numberStings := strings.Split(scanner.Text(), " ")
		numbers := []int{}
		for i := range numberStings {
			num, _ := strconv.Atoi(numberStings[i])
			numbers = append(numbers, num)
		}

		if len(numbers) == 0 {
			continue
		}
		total := 0
		stackNums := [][]int{}
		stackNums = append(stackNums, numbers)
		for abssum(stackNums[len(stackNums)-1]) != 0 {
			tmpNumbers := []int{}
			prevNumbers := stackNums[len(stackNums)-1]
			lastNum := prevNumbers[0]
			for j := range prevNumbers[1:] {
				j++ // i should start at 1
				tmpNumbers = append(tmpNumbers, prevNumbers[j]-lastNum)
				lastNum = prevNumbers[j]
			}
			total += lastNum
			stackNums = append(stackNums, tmpNumbers)
		}
		numbers = append(numbers, total)
		// fmt.Printf("stackNums: %v\n", stackNums)
		// fmt.Printf("total: %v\n", total)
		answer += total
	}
	fmt.Printf("answer: %v\n", answer)
}

func main() {
	parse1()
	parse2()
}
