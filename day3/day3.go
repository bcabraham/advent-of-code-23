package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	testFile    = "/home/babraham/projects/personal/advent-of-code-23/day3/schematic-test.txt"
	problemFile = "/home/babraham/projects/personal/advent-of-code-23/day3/schematic.txt"
	tableWidth  int
)

func Run() {
	lines, err := readLines(problemFile)
	if err != nil {
		log.Fatal(err)
	}

	var sum int
	sum = sumAllPartNumbers(lines)
	fmt.Println("Sum of all part numbers:", sum)

	sum = calcGearRatios(lines)
	fmt.Println("Sum of gear ratios:", sum)
}

func calcGearRatios(input []string) int {
	fmt.Println("#### Calculating Gear Ratios ####")
	tableWidth = len(input[0])
	fmt.Println("tableWidth:", tableWidth)

	data := strings.Join(input, "")
	fmt.Println("input length:", len(data))

	symbolRegex := regexp.MustCompile(`[^0-9\.\n]`)
	symbolIndexes := symbolRegex.FindAllStringIndex(data, -1)
	fmt.Printf("Symbol indexes: %d\n", len(symbolIndexes))

	numericRegex := regexp.MustCompile(`(\d+)`)
	numericIndexes := numericRegex.FindAllStringIndex(data, -1)
	fmt.Printf("Numeric indexes: %d\n", len(numericIndexes))

	sum := 0
	for _, symbol := range symbolIndexes {
		symbolIndex := symbol[0]
		fmt.Printf("Searching adjacents to index: %d\n", symbolIndex)

		searchIndexes := getAdjacentIndexes(symbolIndex, tableWidth)
		fmt.Printf("Adjacent indexes: %+v\n", searchIndexes)

		partNumbers := []int{}
		for _, searchIndex := range searchIndexes {
			// fmt.Println("Checking index:", searchIndex)
			for i, numericIndex := range numericIndexes {
				start, end := numericIndex[0], numericIndex[1]
				if searchIndex >= start && searchIndex < end {
					fmt.Printf("Match found:%d =< %d < %d ~ value: %s\n", start, searchIndex, end, data[start:end])
					numericIndexes = removeAt(numericIndexes, i)

					partNum, _ := strconv.Atoi(data[start:end])
					partNumbers = append(partNumbers, partNum)
				}
			}
		}

		if len(partNumbers) < 2 {
			fmt.Println("Not enough adjacent part numbers found")
		} else if len(partNumbers) > 2 {
			fmt.Println("Too many adjacent part numbers found")
		} else {
			product := partNumbers[0] * partNumbers[1]
			sum += product
			fmt.Printf("Calculating gear ratio: %d * %d = %d sum: %d\n", partNumbers[0], partNumbers[1], product, sum)
		}
	}

	return sum
}

func sumAllPartNumbers(input []string) int {
	tableWidth = len(input[0])
	fmt.Println("tableWidth:", tableWidth)

	data := strings.Join(input, "")
	fmt.Println("input length:", len(data))

	symbolRegex := regexp.MustCompile(`[*]`)
	symbolIndexes := symbolRegex.FindAllStringIndex(data, -1)
	fmt.Printf("Symbol indexes: %d\n", len(symbolIndexes))

	numericRegex := regexp.MustCompile(`(\d+)`)
	numericIndexes := numericRegex.FindAllStringIndex(data, -1)
	fmt.Printf("Numeric indexes: %d\n", len(numericIndexes))

	partNumbers := []int{}
	for _, symbol := range symbolIndexes {
		symbolIndex := symbol[0]
		fmt.Printf("Searching adjacents to index: %d\n", symbolIndex)

		searchIndexes := getAdjacentIndexes(symbolIndex, tableWidth)
		fmt.Printf("Adjacent indexes: %+v\n", searchIndexes)
		for _, searchIndex := range searchIndexes {
			// fmt.Println("Checking index:", searchIndex)
			for i, numericIndex := range numericIndexes {
				start, end := numericIndex[0], numericIndex[1]
				if searchIndex >= start && searchIndex < end {
					fmt.Printf("Match found:%d =< %d < %d ~ value: %s\n", start, searchIndex, end, data[start:end])
					numericIndexes = removeAt(numericIndexes, i)

					partNum, _ := strconv.Atoi(data[start:end])
					partNumbers = append(partNumbers, partNum)
				}
			}
		}
	}

	sum := 0
	for _, num := range partNumbers {
		sum += num
	}

	return sum
}

func getAdjacentIndexes(index, width int) []int {
	indexes := []int{}
	indexes = append(indexes, index-1)       // left
	indexes = append(indexes, index-width-1) // upper left
	indexes = append(indexes, index-width)   // up
	indexes = append(indexes, index-width+1) // upper right
	indexes = append(indexes, index+1)       // right
	indexes = append(indexes, index+width+1) // lower right
	indexes = append(indexes, index+width)   // down
	indexes = append(indexes, index+width-1) // lower left

	return indexes
}

func removeAt(list [][]int, index int) [][]int {

	if index == 0 {
		return list[1:]
	}

	if index >= len(list) {
		return list
	}

	if index == len(list)-1 {
		return list[:len(list)-1]
	}

	var newList [][]int
	newList = list[:index]
	newList = append(newList, list[index+1:]...)

	return newList
}

func readLines(path string) ([]string, error) {
	lines := []string{}

	file, err := os.Open(path)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return lines, err
	}

	return lines, nil
}
