package piscine

func ConcatParams(args []string) string {
	var concatedString string
	for counter, word := range args {
		concatedString += word
		if counter == len(args)-1 {
			break
		}
		concatedString += "\n"
	}

	return concatedString
}
