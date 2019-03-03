package main

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	output := Sum(numbers)
	expected_output := 15

	if output != expected_output {
		t.Errorf("got '%d' want '%d' given '%v'", output, expected_output, numbers)
	}
}

func TestSumWithRange(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	output := SumWithRange(numbers)
	expected_output := 15

	if output != expected_output {
		t.Errorf("got '%d' want '%d' given '%v'", output, expected_output, numbers)
	}
}
