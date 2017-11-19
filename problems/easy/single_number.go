package easy

// SingleNumber finds single integer.
// Given an array of integers, every element appears twice except for one.
// Note:Your algorithm should have a linear runtime complexity.
// Could you implement it without using extra memory?
func SingleNumber(numbers []int) int {
	var result int
	for _, number := range numbers {
		result = result ^ number
	}

	return result
}

// SingleNumberDummy is unoptimised dummy version to find single number
// in array
func SingleNumberDummy(numbers []int) int {
	numbersMap := make(map[int]int)
	for _, number := range numbers {
		numbersMap[number]++
	}

	for number, count := range numbersMap {
		if count == 1 {
			return number
		}
	}

	return 0
}
