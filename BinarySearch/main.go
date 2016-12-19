package main

import "fmt"

func binarySearch(sortedList []int, lookingFor int) int {
	var middleIndex int = len(sortedList) / 2
	var middleElement = sortedList[middleIndex]
	fmt.Println("Middle element is:", middleElement)

	if middleElement == lookingFor {
		return middleIndex
	} else if middleElement > lookingFor {
		// search for lower part
	} else {
		// upper part
	}

	return -1

}

func main() {
	var lookingFor = 6
	var sortedList []int = []int{1, 3, 4, 6, 7, 8, 10, 11, 14}
	fmt.Println("Looking for", lookingFor, "in the sorted list:", sortedList)

	index := binarySearch(sortedList, lookingFor)

	if index >= 0 {
		fmt.Println("Found number", lookingFor, "at possition", index)
	} else {
		fmt.Println("Didn't find the number", lookingFor)
	}
}
