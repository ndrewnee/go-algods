package exercises

type MinQueue interface {
	Min() int
	Push(value int)
	Pop(value int)
}

func findMinWindows(minQueue MinQueue, numbers []int, windowSize int) []int {
	length := len(numbers)
	if length == 0 {
		return nil
	}

	if windowSize == 1 {
		return numbers
	}

	if windowSize < 1 || windowSize >= length {
		return []int{findMin(numbers)}
	}

	for i := 0; i < windowSize; i++ {
		minQueue.Push(numbers[i])
	}
	minWindows := []int{minQueue.Min()}

	for i, j := 0, windowSize; j < length; i, j = i+1, j+1 {
		minQueue.Push(numbers[j])
		minQueue.Pop(numbers[i])
		minWindows = append(minWindows, minQueue.Min())
	}

	return minWindows
}

func findMin(numbers []int) int {
	min := numbers[0]
	for _, number := range numbers {
		if min > number {
			min = number
		}
	}

	return min
}
