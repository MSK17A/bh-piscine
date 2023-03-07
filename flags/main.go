package main

import (
	"fmt"
	"os"
)

func main() {
	params := os.Args[1:]      // We don't want the program name to be printed
	order_flag := false        // Is order required? well, look for -o or --order args
	var str_param string       // This will hold the string that is going to be printed
	var inserted_string string // If you inserted a string using --insert="string" or -i="string" it will be stored here
	print_help_flag := true    // Just another flag that decides wether to print help or not

	for _, param := range params {
		// Search for --order or -o
		if param == "-o" || param == "--order" {
			order_flag = true
			// Search for --help or -h
		} else if param == "--help" || param == "-h" {
			fmt.Printf("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.")
			print_help_flag = false
		} else if param != "" && param[:1] == "-" {
			// Search for insert args
			if param[:3] == "-i=" {
				inserted_string += param[3:]
			} else if param[:9] == "--insert=" {
				inserted_string += param[9:]
			}
		} else {
			// Accumulate the strings here
			str_param += param
		}
		// Khalaaaas we printed everything, why do you need help? (set print_help_flag to)
		print_help_flag = false
	}
	// if the help flag is false continue with accumulating str_param with the inserted_string
	if !print_help_flag {
		str_param += inserted_string
		// Also check if the order is flagged as well
		if order_flag {
			str_param = SortString(str_param)
		}
		// Print or else print help
		fmt.Println(str_param)
	} else {
		fmt.Printf("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n")
	}
}

// This function will sort the string with ASCII order with increasing order
func SortString(str string) string {
	str_rune := []rune(str)
	temp := rune('0') // Temporary variable to be used fo swapping
	min_idx := 0
	for table_idx_i := 0; table_idx_i < len(str_rune)-1; table_idx_i++ {
		min_idx = table_idx_i
		for table_idx_j := table_idx_i + 1; table_idx_j < len(str_rune); table_idx_j++ {
			if str_rune[table_idx_j] < str_rune[min_idx] {
				min_idx = table_idx_j
			}
		}

		if min_idx != table_idx_i {
			temp = str_rune[table_idx_i]
			str_rune[table_idx_i] = str_rune[min_idx]
			str_rune[min_idx] = temp
		}
	}
	return string(str_rune)
}
