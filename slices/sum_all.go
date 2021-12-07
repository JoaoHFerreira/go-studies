package main

func SumAll(numbersToSum ...[]int) []int {
	lenghtOfNumbers := len(numbersToSum)
	sums := make([]int, lenghtOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
		// alternatively
		// sum = append(sums, numbers)
	}
	
	return sums
}