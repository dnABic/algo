package main

import "fmt"

func binarySearch(sortedList []int, lookingFor int) int {
	first := 0
	last := len(sortedList) - 1

	for first <= last {
		middleIndex := first + (last-first)/2
		middleElement := sortedList[middleIndex]
		fmt.Println("Middle element is:", middleElement)

		if middleElement == lookingFor {
			return middleIndex
		} else if middleElement > lookingFor {
			last = middleIndex - 1
		} else {
			first = middleIndex + 1
		}
	}
	return -1
}

func main() {
	var lookingFor = 2
	var sortedList []int = []int{1, 3, 4, 6, 7, 8, 10, 11, 14}
	fmt.Println("Looking for", lookingFor, "in the sorted list:", sortedList)

	index := binarySearch(sortedList, lookingFor)

	if index >= 0 {
		fmt.Println("Found number", lookingFor, "at possition", index)
	} else {
		fmt.Println("Didn't find the number", lookingFor)
	}
}
