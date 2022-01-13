package main

import "fmt"

func main() {
	var a Integer = 1
	var b Integer = 1
	a.Add(3)
	b.AddF(3)
	fmt.Println(a,b)
}
//type Integer struct {
//	val int
//}

/**
  这是什么用法
 */
type Integer int

//类型基于值传递，如果要修改值需要传递指针
func (a *Integer) Add(b Integer){
	 *a += b   //通过指针传递来改变值
}
//是因为Go语言和C语言一样，类型都是基于值传递的。
//要想修改变量的值，只能传递指针
func (a Integer) AddF(b Integer){
	a += b
}
