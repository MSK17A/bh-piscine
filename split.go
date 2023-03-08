package piscine

/*
	func Split(s string, sep rune) []string {
		var strArray []string
		var str_temp []rune

		for _, char := range s {
			if char != sep {
				str_temp = append(str_temp, char)
			} else if str_temp != nil {
				strArray = append(strArray, string(str_temp))
				str_temp = nil
			}
		}

		if str_temp != nil {
			strArray = append(strArray, string(str_temp))
		}
		return strArray
	}
*/
func Split(s string, sep string) []string {
	str_len := len(s)             // Input string length
	sep_len := len(sep)           // Seperator string length
	sep_counts := 0               // Seperator counter
	start_str_counts_flag := true // reset start index flag
	str_Start := 0                // The index of which the string starts
	str_end := 0                  // The index of which the string ends
	result_arr_index := 0         // The index of the result array

	/* This routine will count how many seps are in the string */
	for idx := 0; idx < str_len; idx++ {
		// Check if not end of the string and search for the seperator string
		if idx+sep_len < str_len && s[idx:idx+sep_len] == sep {
			// Increase the count + 1
			sep_counts++
		}
	}

	result_arr := make([]string, sep_counts+1) // Allocate space for the results

	/* Now loop again */
	for idx := 0; idx <= str_len; idx++ {
		// Save the start index for the string
		if start_str_counts_flag {
			str_Start = idx
			start_str_counts_flag = false
		}
		// Stop at the seprator
		if idx+sep_len < str_len && s[idx:idx+sep_len] == sep {
			str_end = idx                                       // Get the string end index
			result_arr[result_arr_index] = s[str_Start:str_end] // Put the string slice from the start to end into results array at location result_arr_index
			start_str_counts_flag = true                        // Set the flag to true to restart the counter
			result_arr_index++                                  // Go to the next index
			idx += sep_len - 1                                  // Skip the string index by length of seperation to prevent insert it into results array
		}
		// If we reach the end just insert the last word into the results array
		if idx == str_len {
			str_end = idx
			result_arr[result_arr_index] = s[str_Start:str_end]
		}
	}

	return result_arr
}
