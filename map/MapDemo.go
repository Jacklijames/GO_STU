package main

import "fmt"

///* 声明变量，默认 map 是 nil */
//var map_variable map[key_data_type]value_data_type
//
///* 使用 make 函数 */
//map_variable := make(map[key_data_type]value_data_type)

func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["里斯"] = 80
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["张三"])
	fmt.Printf("type of a:%T\n", scoreMap)

	userInfo := map[string]string{
		"中国":     "冠军",
		"yuzhou": "牛逼",
	}
	fmt.Println(userInfo)
	//判断某个键是否存在  ok 这个是判断结果，v对应的值，不存在v为0
	value, ok := scoreMap["hah"]
	ok1 := scoreMap["hah"]
	fmt.Println(ok1)
	fmt.Println(ok, value)
	//	遍历map
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	//只遍历key
	for k := range scoreMap {
		fmt.Println("key", k)
	}
}

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*Node
	head, tail *Node
}

type Node struct {
	prev, next *Node
	key, value int
}

func initNode(key, value int) *Node {
	return &Node{
		key:   key,
		value: value,
	}
}

//初始化函数
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    map[int]*Node{},
		head:     initNode(0, 0),
		tail:     initNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.move2Head(node)
	return node.value
}

func (this *LRUCache) Put(key, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initNode(key, value)
		this.add(node)
		this.size++
		if this.size > this.capacity {
			node := this.removeLast()
			delete(this.cache, node.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.move2Head(node)
	}

}

func (this *Node) init() {
	this.next = new(Node)
	this.prev = new(Node)
}

func (this *LRUCache) add(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) move2Head(node *Node) {
	this.remove(node)
	this.add(node)
}

func (this *LRUCache) removeLast() *Node {
	node := this.tail.prev
	this.remove(node)
	return node
}
