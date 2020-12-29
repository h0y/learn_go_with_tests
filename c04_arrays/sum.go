package main

func Sum(numbers [5]int) (sum int) {
	sum = 0

	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }
	
	for _, value := range numbers {
		sum += value
	}

	return sum
}

func Sum2(numbers []int) (sum int) {
	sum = 0
	
	for _, value := range numbers {
		sum += value
	}

	return sum
}

func SumAll(numbers ...[]int) (sums []int) {

	/*
	length := len(numbers)
	sums = make([]int, length)

	for i, value := range numbers {
		sums[i] = Sum2(value)
	}
	*/

	for _, value := range numbers {
		sums = append(sums, Sum2(value))
	}

	return
}

func SumAllTails(numbers ...[]int) (sums []int) {

	for _, value := range numbers {
		if len(value) == 0 {
			sums = append(sums, 0)
		} else {
			number := value[1:]
			sums = append(sums, Sum2(number))
		}
	}

	return 
}