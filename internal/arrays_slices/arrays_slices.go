package arraysslices

func Sum(nums ...int) int {
	var result int
	for _, num := range nums {
		result += num
	}

	return result
}

func SumList(list []int) int {
	return Sum(list...)
}
