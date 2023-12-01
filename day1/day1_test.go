package day1_test

import (
	"advent-of-code-23/day1"
	"testing"
)

// TestParseTwoNumbers calls day1.Parse with a string, checking
// for a valid return value.
func TestParseTwoNumbers(t *testing.T) {
	input := "1abc2"
	want := 12

	num, err := day1.Parse(input)
	if want != num || err != nil {
		t.Fatalf(`Parse("%s") = %d, %v, want match for %d, nil`, input, num, err, want)
	}
}
