package main

import "fmt"

func multiplicationTable(n int) {
	fmt.Printf("  ")
	for i := 1; i <= n; i++ {
		fmt.Printf("%4d", i)
	}
	fmt.Println()

	for i := 1; i <= n; i++ {
		fmt.Printf("%2d", i)
		for j := 1; j <= n; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}

func main() {
	multiplicationTable(6)
}
