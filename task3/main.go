package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func rangeOfValues(min, max int) []int {
	var arr []int

	for i := min; i <= max; i++ {
		if isPrime(i) {
			arr = append(arr, i)
		}
	}
	return arr
}

func main() {
	fmt.Println(rangeOfValues(11, 20))
}
