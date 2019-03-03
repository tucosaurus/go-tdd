package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("sum array using simple for", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		output := Sum(numbers)
		expected_output := 15

		if output != expected_output {
			t.Errorf("got '%d' want '%d' given '%v'", output, expected_output, numbers)
		}
	})

	t.Run("sum all slices and return a new slice containing the result", func(t *testing.T) {
		output := SumAll([]int{1, 2, 3}, []int{4, 5, 6, 7})
		expected_output := []int{6, 22}

		if !reflect.DeepEqual(output, expected_output) {
			t.Errorf("got '%v' want '%v'", output, expected_output)
		}
	})
}
