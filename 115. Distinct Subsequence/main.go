package main

import "fmt"

func search(s, t string, lenS, lenT, idx1, idx2 int, memo [][]int) int {
	if idx2 == lenT {
		return 1
	}
	if idx1 == lenS {
		return 0
	}

	if memo[idx1][idx2] != -1 {
		return memo[idx1][idx2]
	}

	var count = 0
	if s[idx1] == t[idx2] {
		count = search(s, t, lenS, lenT, idx1+1, idx2+1, memo) + search(s, t, lenS, lenT, idx1+1, idx2, memo)
	} else {
		count = search(s, t, lenS, lenT, idx1+1, idx2, memo)
	}

	memo[idx1][idx2] = count
	return count
}

func numDistinct(s string, t string) int {
	var lenS = len(s)
	var lenT = len(t)

	var memo = make([][]int, lenS)
	for i := 0; i < lenS; i++ {
		memo[i] = make([]int, lenT)
		for j := 0; j < lenT; j++ {
			memo[i][j] = -1
		}
	}

	return search(s, t, lenS, lenT, 0, 0, memo)
}

func main() {
	var inputs = map[string]string{
		"rabbbit": "rabbit",
		"babgbag": "bag",
	}

	for s, t := range inputs {
		fmt.Println(numDistinct(s, t))
	}
}
