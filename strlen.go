package piscine

func StrLen(s string) int {
	length := 0
	for i := range s {
		length = i
	}
	return length + 1
}
