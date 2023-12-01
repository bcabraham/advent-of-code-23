package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

func Run() {
	lines, err := ReadLines("/home/babraham/projects/personal/advent-of-code-23/day1/calibration-test.txt")
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
	num := ""

	for _, c := range str {
		if unicode.IsNumber(c) {
			num += string(c)
		}
	}

	if len(num) == 1 {
		num += num
	}

	val, err := strconv.Atoi(num)

	log.Printf("Parse('%s') = %d\n", str, val)

	return val, err
}

func ParseAndSumNumbers(lines []string) (int, error) {
	sum := 0
	for _, line := range lines {
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
