package day1

import (
	"strconv"
	"unicode"
)

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

	return val, err
}
