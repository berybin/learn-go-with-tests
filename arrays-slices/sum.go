package main

func Sum(n []int) int {
	sum := 0
	for _, number := range n {
		sum += number
	}

	return sum
}

// func SumAll(numbersToSum ...[]int) []int {
// 	numLen := len(numbersToSum)
// 	sums := make([]int, numLen)

// 	for i, numbers := range numbersToSum {
// 		sums[i] = Sum(numbers)
// 	}
// 	return sums
// }

func SumAll(numbersToSum ...[]int) []int {
	sums := []int{}

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		}

		tails := numbers[1:]
		sums = append(sums, Sum(tails))
	}
	return sums
}
