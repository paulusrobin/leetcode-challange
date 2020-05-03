package main

import "fmt"

func firstUniqChar(s string) int {
	var existCharacter = make(map[string]bool, 0)
	for i := 0; i < len(s); i++ {
		var exist = false
		if _, ok := existCharacter[string(s[i])]; ok {
			continue
		}
		for j := i+1; j < len(s); j++ {
			if s[i] == s[j] {
				existCharacter[string(s[i])] = true
				exist = true
				break
			}
		}
		if !exist {
			return i
		}
	}
	return -1
}

func main() {
	var inputs = []string{
		"leetcode",
		"loveleetcode",
		"lolo",
	}
	for _, input := range inputs {
		fmt.Println(firstUniqChar(input))
	}
}
