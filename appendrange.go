package piscine

func AppendRange(min, max int) []int {
	var result_array []int
	if min < max {
		for i := min; i < max; i++ {
			result_array = append(result_array, i)
		}
	}
	return result_array
}
