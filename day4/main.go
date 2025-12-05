package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readMap(filename string) [][]bool {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var m [][]bool
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		mapLine := make([]bool, len(line))
		for idx, c := range line {
			if c == '@' {
				mapLine[idx] = true
			}
		}
		m = append(m, mapLine)
	}

	return m
}

func validPosition(m [][]bool, r int, c int) bool {
	if r < 0 {
		return false
	}
	if c < 0 {
		return false
	}
	if r >= len(m) {
		return false
	}
	if c >= len(m[r]) {
		return false
	}
	return true
}

func adjacentCount(m [][]bool, r int, c int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		if r+i < 0 {
			continue
		}
		if r+i >= len(m) {
			continue
		}
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if c+j < 0 {
				continue
			}
			if c+j >= len(m[r+i]) {
				continue
			}
			if m[r+i][c+j] {
				count++
			}
		}
	}
	return count
}

func part1(m [][]bool) {
	count := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if !m[i][j] {
				fmt.Print(".")
				continue
			}
			adjCount := adjacentCount(m, i, j)

			if adjCount < 4 {
				count++
				fmt.Print("x")
			} else {
				fmt.Print("@")
			}
		}
		fmt.Println("")
	}

	fmt.Println("Total: ", count)
}

func part2(m [][]bool) {
	count := 0
	keepGoing := true
	for keepGoing {
		fmt.Println("Next iteration")
		keepGoing = false
		for i := 0; i < len(m); i++ {
			for j := 0; j < len(m[i]); j++ {
				if !m[i][j] {
					fmt.Print(".")
					continue
				}
				adjCount := adjacentCount(m, i, j)

				if adjCount < 4 {
					count++
					keepGoing = true
					m[i][j] = false
					fmt.Print("x")
				} else {
					fmt.Print("@")
				}
			}
			fmt.Println("")
		}
	}

	fmt.Println("Total: ", count)

}

func main() {
	path := "day4/input.txt"
	m := readMap(path)

	fmt.Println("Part 1:")
	part1(m)

	fmt.Println("Part 2:")
	part2(m)
}
