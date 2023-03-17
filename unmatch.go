package piscine

func Unmatch(a []int) int {
	Size := len(a)                 // Size of the array
	PairFlag := make([]bool, Size) // Flags the matched

	/* This will loop throgh the array and whenever a match found it will mark it in the PairFlag */
	for i := 0; i < Size; i++ {
		for j := i + 1; j < Size; j++ {
			if a[i] == a[j] && !PairFlag[j] && !PairFlag[i] {
				PairFlag[j] = true
				PairFlag[i] = true
			}
		}
	}

	/* Output the number that don't have a match (false flag) */
	for i := 0; i < Size; i++ {
		if !PairFlag[i] {
			return a[i]
		}
	}
	return -1
}
