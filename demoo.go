package main

import (
	"fmt"
	_ "github.com/gogf/gf"
)
import _ "fmt"
//var a,b,c,d string
func main() {
	//fmt.Println("hello word!!!")
	a,b,c,d := getStr(string(1), string(2), string(3), string(4))
	fmt.Println(a,b,c,d)

	g,h,j,l := getStrT()
	fmt.Printf("g=%v , h=%v, j=%v, l=%v\n",g,h,j,l)

	bitCal()

	itor()
}


func getStr(firstName, middleName, lastName, nickName string) (string, string, string, string){
	return "1","2","3","5"
}
func getStrT()(firstName, middleName, lastName, nickName string){
	firstName = "1"
	middleName = "2"
	lastName = "3"
	nickName = "4"
	return "1","2","2","4"
}

func bitCal(){
	fmt.Println(1<<3)
}

func itor(){
	arr := [5]int{1,2,3,4,5}
	for i ,v := range arr{
		fmt.Println("遍历数组",i,v)
	}

}