package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Mod(a int, b int) int {
	return (a%b + b) % b
}

func Lines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func Part1(lines []string) {
	pos := 50
	zeroCount := 0
	for _, elt := range lines {
		direction := elt[0]
		magnitude, err := strconv.Atoi(elt[1:])
		if err != nil {
			log.Fatal("Bad magnitude")
		}

		if direction == 'L' {
			pos -= magnitude
		} else if direction == 'R' {
			pos += magnitude
		} else {
			log.Fatal("bad direction " + string(direction))
		}

		pos = Mod(pos, 100)
		if pos == 0 {
			zeroCount++
		}
	}

	fmt.Println(zeroCount)
}

func Part2(lines []string) {
	pos := 50
	zeroCount := 0
	for _, elt := range lines {
		direction := elt[0]
		magnitude, err := strconv.Atoi(elt[1:])
		if err != nil {
			log.Fatal("Bad magnitude")
		}

		newPos := pos
		if direction == 'L' {
			newPos -= magnitude
		} else if direction == 'R' {
			newPos += magnitude
		} else {
			log.Fatal("bad direction " + string(direction))
		}

		newZeroes := 0
		if newPos >= 100 {
			newZeroes = newPos / 100
		}
		if newPos < 0 {
			newZeroes = (-newPos) / 100
			if pos != 0 {
				newZeroes++
			}
		}
		if newPos == 0 {
			newZeroes++
		}

		pos = Mod(newPos, 100)

		if newZeroes != 0 {
			fmt.Println(elt, pos, " crossed! ", newZeroes)
			zeroCount += newZeroes
		}
	}

	fmt.Println(zeroCount)
}

func main() {
	path := "day1/input.txt"
	lines := Lines(path)

	fmt.Println("Part 1:")
	Part1(lines)

	fmt.Println("Part 2:")
	Part2(lines)
}
