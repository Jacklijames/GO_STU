package main

import (
	Jianzhi "MY_DEMO/algorithm/jianzhi"
	"fmt"
)

func main() {
	arr :=[][]int{{1,3,5,6},{2,4,6,7},{4,5,7,8},{8,10,35,89}}
	fmt.Println(Jianzhi.Find(90, arr))
}