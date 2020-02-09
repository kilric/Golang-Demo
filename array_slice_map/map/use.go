package main

import (
	"fmt"
	"sort"
)

func use(){
	fmt.Println("==============map赋值==================")
	var map1 = make(map[string]int,8)
	//赋值方法1
	map1["one"] = 1
	map1["two"] = 2
	map1["three"] = 3
	map1["four"] = 4
	map1["five"] = 5

	//遍历map
	for k,v := range map1{
		fmt.Println("Key:",k,"Value:",v)
	}

	//赋值方法2
	map2 := map[string]int{
		"one":1,
		"two":2,
		"three":3,
		"four":4,   //最后一行要加逗号
	}

	for k,v := range map2 {
		fmt.Println("Key:",k,"Value:",v)
	}
	fmt.Println(map2)
	//取值
	value := map2["one"]
	fmt.Println(value)
	//如果存在，才打印
	value,exists := map2["two"]
	if exists {
		fmt.Println("value:",value)
	}

	//删除元素
	fmt.Println("===================删除=====================")
	delete(map2,"one")
	fmt.Println(map2)

	//删除不存在的键值对
	delete(map2,"ten")
	fmt.Println(map2)

	//排序
	fmt.Println("=================排序===================")
	map3 := map[int]string{
		69:"金克丝",
		5:"卡密尔",
		76:"布兰德",
		34:"菲兹",
	}
	fmt.Println(map3)
	//思路，对键排序，再按照键的顺序打印
	keys := []int{}
	for k,_ := range map3{
		keys = append(keys,k)
	}
	fmt.Println("未排序的键：",keys)
	sort.Ints(keys)
	fmt.Println("排序后的键",keys)

	fmt.Println("========按键的顺序打印===========")
	for _,v := range keys {
		fmt.Println("key:",v,"value:",map3[v])
	}

}
