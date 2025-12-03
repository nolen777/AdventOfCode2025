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

func maxJoltage(bank []int) int {
	pos := 0
	firstDigit := 0

	for idx, batt := range bank[:len(bank)-1] {
		if batt > firstDigit {
			pos = idx
			firstDigit = batt
		}
	}

	secondDigit := slices.Max(bank[pos+1:])

	return 10*firstDigit + secondDigit
}

func part1(banks [][]int) {
	total := 0
	for _, bank := range banks {
		fmt.Println("max joltage: ", maxJoltage(bank))
		total += maxJoltage(bank)
	}

	fmt.Println("Total: ", total)
}

func part2MaxJoltage(bank []int, count int, acc int) int {
	firstDigit := slices.Max(bank[:len(bank)-count+1])
	pos := slices.Index(bank, firstDigit)

	if count <= 1 {
		return 10*acc + firstDigit
	}
	return part2MaxJoltage(bank[pos+1:], count-1, 10*acc+firstDigit)
}

func part2(banks [][]int) {
	total := 0
	for _, bank := range banks {
		jolt := part2MaxJoltage(bank, 12, 0)
		fmt.Println(jolt)
		total += jolt
	}

	fmt.Println("Total: ", total)
}

func main() {
	path := "day3/input.txt"
	banks := readBanks(path)

	fmt.Println(banks)

	fmt.Println("Part 1:")
	part1(banks)

	fmt.Println("Part 2:")
	part2(banks)
}
