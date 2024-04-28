package main

import "fmt"

func Divisors(nums []int) []int {
	var commonDivisors []int
	if len(nums) == 0 {
		return nil
	}
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	for i := 2; i <= min; i++ {
		isCommDiv := true
		for _, num := range nums {
			if num%i != 0 {
				isCommDiv = false
				break
			}
		}
		if isCommDiv {
			commonDivisors = append(commonDivisors, i)
		}
	}
	return commonDivisors
}

func main() {
	numbers := []int{42, 12, 18}
	fmt.Println(Divisors(numbers))
}
