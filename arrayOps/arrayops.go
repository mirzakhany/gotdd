package arrayops

// ArraySum return sum of an array elements
func ArraySum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return sum
}

// SumAll return sum of n number of arrays
func SumAll(numbersToSum ...[]int) []int {

	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, ArraySum(numbers))
	}
	return sums
}

// SumAllTail ...
func SumAllTail(numberToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numberToSum {
		tail := numbers[1:]
		sums = append(sums, ArraySum(tail))
	}
	return sums
}
