package arrays_and_slices

func SumWithArray(numbers [5]int) int {
	var result int
	for _,number := range numbers {
		result += number
	}
	return result
}

func SumWithSlice(numbers []int) int {
	var result int
	for _,number := range numbers {
		result += number
	}
	return result
}

func SumAll(numbersToSum ...[]int) (sum []int) {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, SumWithSlice(numbers))
	}

	return sums
}
