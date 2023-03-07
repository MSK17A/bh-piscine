package main

import (
	"fmt"
	"os"
)

func main() {
	params := os.Args[1:]
	order_flag := false
	var str_param string
	var inserted_string string
	print_help_flag := true

	for _, param := range params {
		// Search for --order or -o
		if param == "-o" || param == "--order" {
			order_flag = true
		} else if param == "--help" || param == "-h" {
			fmt.Printf("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.")
			print_help_flag = false
		} else if param != "" && param[:1] == "-" {
			if param[:3] == "-i=" {
				inserted_string += param[3:]
			} else if param[:9] == "--insert=" {
				inserted_string += param[9:]
			}
		} else {
			str_param += param
		}
		print_help_flag = false
	}
	if !print_help_flag {
		str_param += inserted_string
		if order_flag {
			str_param = SortString(str_param)
		}
		fmt.Println(str_param)
	} else {
		fmt.Printf("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.")
	}
}

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
