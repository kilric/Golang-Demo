package main

import "interface/util"

func main() {
	user := util.User{
		Name:  "lisa",
		Email: "lisa@qq.com",
	}

	//f1(user)  不能调用值为*user的方法集
	f1(&user)
	f2(user)
	f2(&user)

	//方法集的规则
	/*
	   values               methods Receivers
	   T                     (t T)
	   *T                    (t T) & (t *T)
	*/
	/*
	   Methods Receivers    Values
	   (t T)                T   *T
	   (t *T)                *T
	*/

}

func f1(n util.Notify) {
	n.Notify()
}

func f2(o util.Out) {
	o.Out()
}
