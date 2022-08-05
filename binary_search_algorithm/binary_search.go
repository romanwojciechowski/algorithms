package main

import "fmt"

func main() {
	var sortedList []int = []int{1,3,5,8,9,10,14}
	var lookingFor int = 1
	result := binarySearch(sortedList, lookingFor)
	if result >= 0 {
		fmt.Println(result)
	} else {
		fmt.Println("not found")
	}
}

func binarySearch(array []int, out int) int {
	var low int = 0
	var high int = len(array)

	for low <= high {
		var mid int = low + (high - low)/2
		var value int = array[mid]
		if value == out {
			return mid
		} else if value < out {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
