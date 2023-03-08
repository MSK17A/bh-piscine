package piscine

func MakeRange(min, max int) []int {
	var nil_array []int
	if min > max {
		return nil_array
	}
	size := max - min
	result_array := make([]int, size)

	for i := 0; i < size; i++ {
		result_array[i] = min + i
	}
	return result_array
}
