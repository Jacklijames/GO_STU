package main

import "fmt"

func main() {
	a := 10
	b := &a                            //1、对变量进行取地址（&）操作，可以获得这个变量的指针变量。    2.指针变量的值是指针地址。
	fmt.Printf("a:%d ptr:%p\n", a, &a) //a:10 ptr:0xc0000ac058
	fmt.Printf("b:%p type:%T\n", b, b) //b:0xc0000ac058 type:*int
	fmt.Println(&b)                    //0xc0000d8018
	fmt.Println(*b)                    // 10
	c := *b                            // 指针取值   对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。
	fmt.Println(c)
	fmt.Printf("type of c:%T\n", c)  //type of c:int
	fmt.Printf("value of c:%v\n", c) //value of c:10

}
