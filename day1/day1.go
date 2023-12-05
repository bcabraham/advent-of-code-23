package day1

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var (
	testFile    = "/home/babraham/projects/personal/advent-of-code-23/day1/calibration-test.txt"
	testFile2   = "/home/babraham/projects/personal/advent-of-code-23/day1/calibration-test-2.txt"
	problemFile = "/home/babraham/projects/personal/advent-of-code-23/day1/calibration.txt"
	numMap      = map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	tokens = []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
)

func Run() {
	lines, err := ReadLines(problemFile)
	if err != nil {
		log.Fatal(err)
	}

	total, err := ParseAndSumNumbers(lines)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Total: ", total)
}

func Parse(str string) (int, error) {
	var firstDigit, lastDigit int
	var firstIndex, lastIndex int = 9999, -9999

	for _, token := range tokens {
		i := strings.Index(str, token)
		j := strings.LastIndex(str, token)

		if i >= 0 && i < firstIndex {
			firstIndex = i
			firstDigit = numMap[token]
			log.Printf("\tFound %s at index: %d\n", token, i)
		}

		if j >= 0 && j > lastIndex {
			lastIndex = j
			lastDigit = numMap[token]
			log.Printf("\tFound %s at index: %d\n", token, j)
		}
	}

	num := firstDigit*10 + lastDigit
	log.Printf("Parse('%s') = %d\n", str, num)

	return num, nil
}

func ParseAndSumNumbers(lines []string) (int, error) {
	sum := 0
	for i, line := range lines {
		log.Printf("Parsing %d %s\n", i+1, line)
		num, err := Parse(line)
		if err != nil {
			return sum, err
		}

		sum += num
	}

	return sum, nil
}

func ReadLines(path string) ([]string, error) {
	lines := []string{}

	file, err := os.Open(path)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return lines, err
	}

	return lines, nil
}
