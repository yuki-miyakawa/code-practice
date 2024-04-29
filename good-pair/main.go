package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 1, 1, 3}
	fmt.Println(goodPair(arr))
}

func goodPair(array []int) int {
	var goodPairs int
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] == array[j] {
				goodPairs++
			}
		}
	}
	return goodPairs
}
