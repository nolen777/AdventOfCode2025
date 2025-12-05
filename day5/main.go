package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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

func consolidate(a ingredientRange, b ingredientRange) ingredientRange {
	if a.first > b.first {
		panic("bad")
	}

	return ingredientRange{first: a.first, last: max(a.last, b.last)}
}

func part2(ranges []ingredientRange) {
	slices.SortFunc(ranges, func(a ingredientRange, b ingredientRange) int {
		if a.first == b.first {
			return a.last - b.last
		}
		return a.first - b.first
	})

	consolidated := make([]ingredientRange, 0)
	consolidated = append(consolidated, ranges[0])
	curSize := 1

	for _, r := range ranges[1:] {
		if r.first-1 <= consolidated[curSize-1].last {
			consolidated[curSize-1] = consolidate(consolidated[curSize-1], r)
		} else {
			consolidated = append(consolidated, r)
			curSize++
		}
	}

	totalCount := 0
	for _, r := range consolidated {
		fmt.Println(r, (r.last - r.first + 1))
		totalCount += (r.last - r.first + 1)
	}

	fmt.Println("Fresh count: ", totalCount)
}

func main() {
	path := "day5/input.txt"
	ranges, available := readData(path)

	fmt.Println("Part 1:")
	part1(ranges, available)

	fmt.Println("Part 2:")
	part2(ranges)
}
