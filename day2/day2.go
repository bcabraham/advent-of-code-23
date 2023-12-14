package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	testFile    = "/home/babraham/projects/personal/advent-of-code-23/day2/cube-test.txt"
	problemFile = "/home/babraham/projects/personal/advent-of-code-23/day2/cube-game.txt"
)

func Run() {
	lines, err := readLines(problemFile)
	if err != nil {
		log.Fatal(err)
	}

	// referenceCubes := map[string]int{
	// 	"red":   12,
	// 	"green": 13,
	// 	"blue":  14,
	// }

	games := []Game{}
	for _, line := range lines {
		game := loadGame(line)
		fmt.Printf("%+v\n", game)
		games = append(games, game)
	}

	sum := 0
	for _, game := range games {
		// isPossible := isGamePossible(game, referenceCubes)
		// fmt.Printf("Game %d IsGamePossible ? %t\n", game.ID, isPossible)

		minCubes := getMinCubes(game)
		sum += getPowerSet(minCubes)
	}

	fmt.Println("Product of the sets:", sum)
}

type Game struct {
	ID      int
	Subsets [][]Subset
}

type Subset struct {
	Count int
	Color string
}

func getMinCubes(game Game) map[string]int {
	minCubes := map[string]int{
		"red":   -1,
		"blue":  -1,
		"green": -1,
	}

	for _, subsets := range game.Subsets {
		for _, subset := range subsets {
			if subset.Count > minCubes[subset.Color] {
				minCubes[subset.Color] = subset.Count
			}
		}
	}

	fmt.Printf("Game %d: %+v\n", game.ID, minCubes)

	return minCubes
}

func getPowerSet(set map[string]int) int {
	product := 1

	for _, count := range set {
		product *= count
	}

	fmt.Printf("PowerSet %+v = %d\n", set, product)

	return product
}

func isGamePossible(game Game, refCubes map[string]int) bool {
	for i, subsets := range game.Subsets {
		for _, subset := range subsets {
			if subset.Count > refCubes[subset.Color] {
				fmt.Printf("Game %d, subset %d has too many %s cubes: %d > %d\n",
					game.ID,
					i,
					subset.Color,
					subset.Count,
					refCubes[subset.Color],
				)

				return false
			}
		}
	}

	return true
}

func loadGame(input string) Game {
	parts := strings.Split(input, ": ")
	id, _ := strconv.Atoi(strings.Replace(parts[0], "Game ", "", 1))

	subsets := loadSubsets(parts[1])

	return Game{id, subsets}
}

func loadSubsets(input string) [][]Subset {
	var subsets [][]Subset

	// 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	for _, subset := range strings.Split(input, ";") {
		sb := []Subset{}

		subset = strings.TrimSpace(subset)

		for _, s := range strings.Split(subset, ",") {
			s = strings.TrimSpace(s)
			reveal := strings.SplitN(s, " ", 2)
			count, _ := strconv.Atoi(reveal[0])
			color := reveal[1]

			sb = append(sb, Subset{count, color})
		}

		subsets = append(subsets, sb)
	}

	return subsets
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
