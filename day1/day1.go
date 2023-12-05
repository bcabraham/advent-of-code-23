package day1

import (
	"bufio"
	"log"
	"os"
	"regexp"
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

func Parse(regex *regexp.Regexp, str string) (int, error) {
	var first, last int

	log.Println("Parsing", str)
	matches := regex.FindAllString(str, -1)

	numMatches := len(matches)
	log.Println("Matches found:", numMatches)
	for _, match := range matches {
		log.Println("\t", match)
	}

	if numMatches > 0 {
		first = numMap[matches[0]]
	}

	if numMatches == 1 {
		last = numMap[matches[0]]
	} else {
		last = numMap[matches[numMatches-1]]
	}

	num := first*10 + last

	log.Printf("Parse('%s') = %d\n", str, num)

	return num, nil
}

func ParseAndSumNumbers(lines []string) (int, error) {
	rg, _ := regexp.Compile(`one|two|three|four|five|six|seven|eight|nine|\d`)

	sum := 0
	for _, line := range lines {
		num, err := Parse(rg, line)
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
