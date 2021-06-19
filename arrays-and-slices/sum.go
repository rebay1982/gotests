package main

func Sum(numbers []int) (sum int) {

	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {

		var sum int
		if len(numbers) > 0 {
			sum = Sum(numbers[1:])
		} else {
			sum = 0
		}

		sums = append(sums, sum)
	}
	return
}
