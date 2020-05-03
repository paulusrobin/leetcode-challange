package main

import "fmt"

func findNumbers(nums []int) int {
	var count = 0;
	for _, num := range nums {
		str := fmt.Sprintf("%d", num)
		if len(str)%2 == 0 {
			count = count + 1
		}
	}
	return count
}

func main() {
	var inputs = [][]int{
		{12,345,2,6,7896},
		{555,901,482,1771},
	}

	for _, input := range inputs {
		fmt.Println(findNumbers(input))
	}
}
