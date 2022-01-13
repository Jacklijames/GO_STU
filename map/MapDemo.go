package main

import "fmt"

func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["里斯"] = 80
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["张三"])
	fmt.Printf("type of a:%T\n",scoreMap)
	
	userInfo := map[string]string{
		"中国":"冠军",
		"yuzhou":"牛逼",
	}
	fmt.Println(userInfo)
	//判断某个键是否存在  ok 这个是判断结果，v对应的值，不存在v为0
	value,ok := scoreMap["hah"]
	fmt.Println(ok,value)
//	遍历map
	for k,v := range scoreMap{
		fmt.Println(k,v)
	}
	//只遍历key
	for k := range scoreMap{
		fmt.Println("key",k)
	}


}
