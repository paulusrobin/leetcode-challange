package main

import "fmt"

func fizzBuzz(n int) []string {
	result := make([]string,0)
	for num := 1 ; num <= n ; num++ {
		str := ""
		if num%3 == 0 {
			str += "Fizz"
		}
		if num%5 == 0 {
			str += "Buzz"
		}
		if str == "" {
			str = fmt.Sprintf("%d", num)
		}
		result = append(result, str)
	}
	return result
}

func main(){
	var inputs = []int{
		15,
	}

	for _, input := range inputs {
		fmt.Println(fizzBuzz(input))
	}
}
