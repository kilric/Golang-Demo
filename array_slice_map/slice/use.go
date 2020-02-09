/*
 * Copyright (c) 2019  kilric.
 */

package main

import "fmt"

func use(){
	reslice()
}

//reslice，利用切片创建切片

func reslice(){
	sli1 := []int{1,2,3,4,5,6,7}
	fmt.Println("sli1的长度：",len(sli1),"sli1的容量:",cap(sli1))
	fmt.Println(sli1)
	//sli2由sli1reslice而来，和sli1共享底层数组
	sli2 := sli1[1:3]  //从下标1开始的3个元素
	//newSlice := slice[i:j:k]  容量=k-i  长度=j-i
	fmt.Println("sli2的长度：",len(sli2),"sli2的容量:",cap(sli2))
	fmt.Println(sli2)

	/*
	当新切片的长度小于容量时，和原切片共享底层数组，当长度大于容量后，会复制到新的数组
	因此，一般reslice的时候，让长度等于容量
	 */
	//For example
	slice := make([]int,4,10)
	slice = append(slice,1,2,3,4,5,5,6,7)
	fmt.Println(slice)

	newSlice1 := slice[1:5]
	fmt.Println(newSlice1)

	newSlice1[2] = 333
	fmt.Println("slice:",slice,"newSlice1:",newSlice1)

	newSlice2 := slice[2:5:5]
	fmt.Println("slice:",slice,"newSlice2:",newSlice2)

	newSlice2[1] = 444444

	fmt.Println("slice:",slice,"newSlice:",newSlice2)





}
