package main

import "fmt"

func main() {
	res := climbStairs(5)
	fmt.Println(res)
}

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	first, two, temp := 1, 2, 0
	for i := 3; i <= n; i++ {
		temp = first + two
		first = two
		two = temp
	}
	return temp
}
