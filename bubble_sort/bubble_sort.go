package main

import "fmt"

func main() {
	var list []int = []int{5,1,3,4,2,6}
	bubbleSort(list)
	fmt.Println(list)
}

func sweep(numbers []int) {
	length := len(numbers)
	first := 0
	second := 1
	

	for second < length {
		firstNumber := numbers[first]
		secondNumber := numbers[second]

		if firstNumber > secondNumber {
			numbers[first] = secondNumber
			numbers[second] = firstNumber
		}

		first++
		second++
	}
}

func bubbleSort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		sweep(numbers)
	}
}
