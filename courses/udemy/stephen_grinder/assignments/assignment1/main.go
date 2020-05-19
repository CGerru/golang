package main

import "fmt"

func main() {
	numbers := make([]int, 10)

	for i := 0; i < len(numbers); i++ {
		numbers[i] = i
		if numbers[i]%2 == 0 {
			fmt.Println(numbers[i], "is even")
		} else {
			fmt.Println(numbers[i], "is odd")
		}
	}
}
