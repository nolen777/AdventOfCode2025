package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type IdRange struct {
	first int
	last  int
}

func Ids(filename string) []IdRange {
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

	rangeStrings := strings.Split(lines[0], ",")
	ranges := make([]IdRange, 0, len(rangeStrings))

	for _, elt := range strings.Split(lines[0], ",") {
		components := strings.Split(elt, "-")
		first, _ := strconv.Atoi(components[0])
		last, _ := strconv.Atoi(components[1])

		ranges = append(ranges, IdRange{first, last})
	}
	return ranges
}

func isInvalid(id int) bool {
	idStr := strconv.Itoa(id)
	l := len(idStr)
	if l%2 == 1 {
		return false
	}

	firstHalf := idStr[0 : l/2]
	secondHalf := idStr[l/2:]

	if firstHalf == secondHalf {
		return true
	}
	return false
}

func Part1(idRanges []IdRange) {
	invalidCount := 0
	invalidSum := 0

	for _, idRange := range idRanges {
		for id := idRange.first; id <= idRange.last; id++ {
			if isInvalid(id) {
				invalidCount++
				invalidSum += id
			}
		}
	}

	fmt.Println("Total invalid: ", invalidCount)
	fmt.Println("Sum: ", invalidSum)
}

func isInvalidPart2(id int) bool {
	idStr := strconv.Itoa(id)
	l := len(idStr)

	for repeatSize := 1; repeatSize <= l/2; repeatSize++ {
		if l%repeatSize != 0 {
			continue
		}
		potentialRepeat := idStr[:repeatSize]

		couldMatch := true
		for pos := repeatSize; pos < l; pos += repeatSize {
			if idStr[pos:pos+repeatSize] != potentialRepeat {
				couldMatch = false
				break
			}
		}

		if couldMatch {
			fmt.Println(idStr, " is invalid with repeating ", potentialRepeat)
			return true
		}
	}

	return false
}

func Part2(idRanges []IdRange) {
	invalidCount := 0
	invalidSum := 0

	for _, idRange := range idRanges {
		for id := idRange.first; id <= idRange.last; id++ {
			if isInvalidPart2(id) {
				invalidCount++
				invalidSum += id
			}
		}
	}

	fmt.Println("Total invalid: ", invalidCount)
	fmt.Println("Sum: ", invalidSum)
}

func main() {
	path := "day2/input.txt"
	ids := Ids(path)

	fmt.Println(ids)

	fmt.Println("Part 1:")
	Part1(ids)

	fmt.Println("Part 2:")
	Part2(ids)
}
