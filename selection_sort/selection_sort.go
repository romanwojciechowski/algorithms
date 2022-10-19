package main

import "fmt"

func main() {
	list := []int{5,1,3,4,0,-2}
	result := findSmallest(list)
	fmt.Println(result)
}

func findSmallest(numbers []int) int {
	length := len(numbers)
	smallest := numbers[0]

	for i := 1; i < length; i++ {
		if numbers[i] < smallest {
			smallest = numbers[i]
		}
	}
	return smallest
}
