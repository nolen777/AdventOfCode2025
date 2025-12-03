package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func readBanks(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var banks [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var newBank []int
		for _, c := range line {
			batt, _ := strconv.Atoi(string(c))
			newBank = append(newBank, batt)
		}
		banks = append(banks, newBank)
	}

	return banks
}

func maxJoltage(bank []int, count int) int {
	var inner func([]int, int, int) int

	inner = func(remBank []int, remCount int, acc int) int {
		firstDigit := slices.Max(remBank[:len(remBank)-remCount+1])
		pos := slices.Index(remBank, firstDigit)

		if remCount <= 1 {
			return 10*acc + firstDigit
		}
		return inner(remBank[pos+1:], remCount-1, 10*acc+firstDigit)
	}

	return inner(bank, count, 0)
}

func part1(banks [][]int) {
	total := 0
	for _, bank := range banks {
		total += maxJoltage(bank, 2)
	}

	fmt.Println("Total: ", total)
}

func part2(banks [][]int) {
	total := 0
	for _, bank := range banks {
		jolt := maxJoltage(bank, 12)
		total += jolt
	}

	fmt.Println("Total: ", total)
}

func main() {
	path := "day3/input.txt"
	banks := readBanks(path)

	fmt.Println("Part 1:")
	part1(banks)

	fmt.Println("Part 2:")
	part2(banks)
}
