package main

import "fmt"

func main() {
//	声明切片
	var sli []int
	if sli == nil {
		fmt.Println("是空")
	}else {
		fmt.Println("不是空")
	}
//	2
	slli := []int{}
//	3 make
	var s3 []int = make([]int,0)
	fmt.Println(sli,slli,s3)
//	初始化赋值
	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)
	s5 := []int{1,2,3}
	fmt.Println(s5)
//	从数组切片
arr := [5]int{1,2,3,4,5}
var s6 []int
//前包后不包
s6 = arr[0:4]
fmt.Println(s6)

//切片初始化
var arr1 = [...]int{1,2,3,4,5,6,7,8,9,0}
fmt.Println(arr1)

 ss1  := arr1[3:5:10]
fmt.Println(ss1)

	
}
