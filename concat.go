package piscine

func Concat(str1 string, str2 string) string {

	if str1 == "" || str2 == "" {
		return "0"
	}
	return str1 + str2
}
