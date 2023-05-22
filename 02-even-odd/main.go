package main

import (
	"fmt"
)

func main() {
	numbers := numSlice()
	evenOdd(numbers)
}

func numSlice() []int {
	var numbers []int
	for i := range [11]int{} {
		numbers = append(numbers, i)
	}
	return numbers
}

func evenOdd(numbers []int) {
	for _, value := range numbers {
		if value%2 == 0 {
			fmt.Println("Value", value, "is even")
		} else {
			fmt.Println("Value", value, "is odd")
		}
	}
}
