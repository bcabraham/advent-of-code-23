package day1_test

import (
	"advent-of-code-23/day1"
	"regexp"
	"testing"
)

var (
	rg, _ = regexp.Compile(`one|two|three|four|five|six|seven|eight|nine|\d`)
)

// TestParseTwoNumbers calls day1.Parse with a string, checking
// for a valid return value.
func TestParseTwoNumbers(t *testing.T) {
	input := "1abc2"
	want := 12

	num, err := day1.Parse(rg, input)
	if want != num || err != nil {
		t.Fatalf(`Parse("%s") = %d, %v, want match for %d, nil`, input, num, err, want)
	}
}

// TestParseOneNumber calls day1.Parse with a string, checking
// for a valid return value.
func TestParseOneNumber(t *testing.T) {
	input := "treb7uchet"
	want := 77

	num, err := day1.Parse(rg, input)
	if want != num || err != nil {
		t.Fatalf(`Parse("%s") = %d, %v, want match for %d, nil`, input, num, err, want)
	}
}

// TestParseAndSumNumbers calls day1.ParseAndSumNumbers with a slice of strings, checking
// for a valid return value.
func TestParseAndSumNumbers(t *testing.T) {
	input := []string{"1abc2", "treb7uchet"}
	want := 89

	num, err := day1.ParseAndSumNumbers(input)
	if want != num || err != nil {
		t.Fatalf(`ParseAndSumNumbers("%s") = %d, %v, want match for %d, nil`, input, num, err, want)
	}
}

// TestParseMultipleNumbers calls day1.Parse with a string, checking
// for a valid return value.
func TestParseMultipleNumbers(t *testing.T) {
	input := "a1b2c3d4e5f"
	want := 15

	num, err := day1.Parse(rg, input)
	if want != num || err != nil {
		t.Fatalf(`Parse("%s") = %d, %v, want match for %d, nil`, input, num, err, want)
	}
}
