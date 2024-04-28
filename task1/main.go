package main

import "fmt"

func formatString(num int) string {
	var result string

	if num%10 == 1 && num%100 != 11 {
		result = fmt.Sprintf("%d компьютер", num)
	} else if num%10 == 2 || num%10 == 3 || num%10 == 4 {
		result = fmt.Sprintf("%d компьютера", num)
	} else {
		result = fmt.Sprintf("%d компьютеров", num)
	}
	return result
}

func main() {
	fmt.Println(formatString(11))
}
