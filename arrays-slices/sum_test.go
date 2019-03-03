package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("sum array using simple for", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		output := Sum(numbers)
		expected_output := 15

		if output != expected_output {
			t.Errorf("got '%d' want '%d' given '%v'", output, expected_output, numbers)
		}
	})
}
