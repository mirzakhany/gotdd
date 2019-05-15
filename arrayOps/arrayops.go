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

	count := len(numbersToSum)
	sums := make([]int, count)

	for i, number := range numbersToSum {
		sums[i] = ArraySum(number)
	}
	return sums
}
