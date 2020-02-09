package main

import "fmt"

func create(){
	//声明一个固定长度的数组
	var array [5]int
	fmt.Println(array)

	//声明一个数组，并初始化
	var array1 = [5]int{ 5,4,3,2,1 }
	array2 := [5]int{1,2,3,4,5}
	fmt.Println(array1)
	fmt.Println(array2)

	//让go自动计算数组长度
	array3 := [...]int{11,22,33,44,55,66}
	fmt.Println(array3,"长度：", len(array3))

	//声明数组，并制定特殊元素的值,如果用自动推断长度的方式，就会按照能满足最大下标的长度创建
	array4 := [...]int{ 0: 123,12: 321}
	fmt.Println(array4)
}
