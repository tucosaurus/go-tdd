package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numberSlices ...[]int) []int {
	numberOfSlices := len(numberSlices)
	sums := make([]int, numberOfSlices)

	for i, numbers := range numberSlices {
		sums[i] += Sum(numbers)
	}

	return sums
}
