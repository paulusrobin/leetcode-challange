package main

import "fmt"

func fourSumCount(A []int, B []int, C []int, D []int) int {
	tupleAB := getTuple(A,B)
	tupleCD := getTuple(C,D)

	result := 0
	for key, value := range tupleAB {
		if count, ok := tupleCD[-key]; ok {
			result += value * count
		}
	}
	return result
}

func getTuple(A []int, B []int) map[int]int {
	tuple := make(map[int]int)
	for _, a := range A {
		for _, b := range B {
			tuple[a+b]++
		}
	}
	return tuple
}

func main() {
	var inputs = [][][]int{
		{
			{1, 2},
			{-2,-1},
			{-1, 2},
			{ 0, 2},
		},
	}

	for _, input := range inputs {
		fmt.Println(fourSumCount(input[0],input[1],input[2],input[3]))
	}
}
