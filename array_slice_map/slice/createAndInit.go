/*
 * Copyright (c) 2019  kilric.
 */

package main

import "fmt"

func create(){
	//nil切片
	var slice1 []int
	fmt.Println(slice1,slice1 == nil)  //[]  true
	//初始化1，长度为4,容量为4
	var slice2 = []int{1,2,3,4}
	fmt.Println(slice2,slice2 == nil,"cap:" ,cap(slice2),"len:",len(slice2))

	//创建长度和容量都是10的切片
	slice3 := []int{0:1,9:10}
	fmt.Println(slice3,slice3 == nil,"cap:" ,cap(slice3),"len:",len(slice3))

	//使用make创建
	var sli []int = make([]int,2,5)  //长度2,容量5
	sli[0] = 2
	sli[1] = 3
	//sli[2] = 4  panic，因为长度是2
	fmt.Println(sli)

	//创建空切片，不是nil切片
	s1 := make([]int,0)
	s2 := []int{}
	fmt.Println("s1:",s1,"s2:",s2)
	fmt.Println("s1 == nil?",s1 == nil,"s2 == nil?",s2 == nil)

}
