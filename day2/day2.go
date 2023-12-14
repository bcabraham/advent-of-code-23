package day2

import (
	"bufio"
	"log"
	"os"
)

var (
	testFile = "/home/babraham/projects/personal/advent-of-code-23/day2/cube-test.txt"
)

func Run() {
	lines, err := ReadLines(testFile)
	if err != nil {
		log.Fatal(err)
	}

}

func ReadLines(path string) ([]string, error) {
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
