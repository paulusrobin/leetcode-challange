package main

import "fmt"

func findDuplicate(nums []int) int {
	data := make(map[int]bool)
	for _, n := range nums {
		if _, ok := data[n]; ok {
			return n
		}
		data[n] = true
	}
	return -1
}

func main() {
	var inputs = [][]int{
		{1,3,4,2,2},
		{3,1,3,4,2},
	}

	for _, input := range inputs {
		fmt.Println(findDuplicate(input))
	}
}
