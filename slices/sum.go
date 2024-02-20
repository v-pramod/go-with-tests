package slices

func Sum(num []int) int {
	sum := 0
	for _, number := range num {
		sum += number
	}
	return sum
}

func SumAll(numToSum ...[]int) []int {
	var sumSlice []int
	sum := 0
	for _, num := range numToSum {
		for _, num1 := range num {
			sum += num1
		}
		sumSlice = append(sumSlice, sum)
		sum = 0
	}
	return sumSlice
}

func SumTails(numToSum ...[]int) []int {
	var sumSlice []int

	for _, num := range numToSum {
		if len(num) == 0 {
			sumSlice = append(sumSlice, 0)
		} else {
			tail := num[1:]
			sumSlice = append(sumSlice, Sum(tail))
		}
	}

	return sumSlice
}
