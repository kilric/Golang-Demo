package main

import "fmt"

func use(){
	//修改具体索引的值
	array := [8]int{1,2,3,4,5,6}
	//修改索引为2的元素
	array[2] = 123
	fmt.Println(array)

	//访问指针数组的元素，即该数组的元素全部是指针
	array1 := [5]*int{1:new(int),4:new(int)}
	//为其赋值
	//*array1[2] = 43  错误示例，没有初始化为其分配空间
	*array1[1] = 44
	for i,v := range array1{
		if v != nil {
			fmt.Println(i,":",*v,"  ")
		}
	}

	//数组之间的赋值，只能在同一种类型之间，即类型和长度都相同
	arr := [5]string{"red","blue","green","black","white"}
	arr1 := [5]string{"1","2"}
	fmt.Println("arr:",arr,"arr1",arr1)
	arr1 = arr
	fmt.Println("arr1:",arr1)

	//指针数组的赋值，赋值后两个数组元素指向同一个值
	arr2 := [2]*string{0:new(string),1:new(string)}
	*arr2[0] = "one"
	*arr2[1] = "two"
	var arr3 [2]*string
	arr3 = arr2
	for i,v := range arr2{
		if v != nil {
		fmt.Println(i,":",*v,"  ")
		}
	}
	for i,v := range arr3{
		if v != nil {
		fmt.Println(i,":",*v,"  ")
		}
	}

	//改变一个另一个也会改变
	*arr3[1] = "modified"
	for i,v := range arr2{
		if v != nil {
		fmt.Println(i,":",*v,"  ")
	}
	}
	for i,v := range arr3{
		if v != nil {
			fmt.Println(i,":",*v,"  ")
		}
	}
	//但是直接改变指针就不会影响另一个
	arr3[1] = new(string)
	*arr3[1] = "three"
	for i,v := range arr2{
		if v != nil {
			fmt.Println(i,":",*v,"  ")
		}
	}
	for i,v := range arr3{
		if v != nil {
			fmt.Println(i,":",*v,"  ")
		}
	}

	//传递数组
	deliverArray()
}

func deliverArray(){
	var array = [4]int{1,2,3,4}
	foo(array)
}

func foo(arr [4]int){
	fmt.Println("传递的数组：" ,arr)
}
