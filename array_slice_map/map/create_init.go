package main

import "fmt"

func create(){
	//map的创建
	//创建一个 nil map，不能向其添加元素，需要被初始化或者接收其他map的引用
	fmt.Println("================map的创建====================")
	var map1 map[string]int
	fmt.Println(map1==nil)  //true

	//创建一个长度为0的map
	var map2 = map[string]int{}
	//或者map2 := map[string]int
	fmt.Println(map2==nil)  //false
	fmt.Println(map2,"map2:len=",len(map2))
	map2["one"] = 1
	map2["two"] = 2
	//cap函数的参数不能是map
	fmt.Println(map2,"len=",len(map2))

	//通过make函数创建
	map3 := make(map[string]int,5)  //长度为5
	map3["one"] = 1
	map3["two"] = 2
	map3["three"] = 3
	fmt.Println(map3,"len=", len(map3))

}