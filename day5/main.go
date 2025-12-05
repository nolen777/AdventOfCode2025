package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ingredientRange struct {
	first int
	last  int
}

func readData(filename string) ([]ingredientRange, []int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	freshIngredients := make([]ingredientRange, 0, 0)
	ingredientList := make([]int, 0, 0)

	scanner := bufio.NewScanner(file)

	// first the ranges
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		vals := strings.Split(line, "-")
		first, _ := strconv.Atoi(vals[0])
		second, _ := strconv.Atoi(vals[1])

		freshIngredients = append(freshIngredients, ingredientRange{first: first, last: second})
	}

	for scanner.Scan() {
		line := scanner.Text()

		ing, _ := strconv.Atoi(line)

		ingredientList = append(ingredientList, ing)
	}

	return freshIngredients, ingredientList
}

func contained(r ingredientRange, a int) bool {
	return a >= r.first && a <= r.last
}

func part1(ranges []ingredientRange, available []int) {
	count := 0

	for _, ing := range available {
		for _, r := range ranges {
			if contained(r, ing) {
				count++
				break
			}
		}
	}

	fmt.Println("Count: ", count)
}

func part2(ranges []ingredientRange, available []int) {
}

func main() {
	path := "day5/input.txt"
	ranges, available := readData(path)

	fmt.Println(ranges)
	fmt.Println(available)

	fmt.Println("Part 1:")
	part1(ranges, available)

	fmt.Println("Part 2:")
	part2(ranges, available)
}
