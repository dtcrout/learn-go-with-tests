package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	// numbersToSum = [[1, 2] [0, 9]]
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	// tail = [2]
	// sums = [sum([2])] = [2]
	// tail = [9]
	// sums = [2, sum([9])] = [2, 9]

	return sums
}
