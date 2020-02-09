package main

import "fmt"

func createAndInit() {
	//定义一个User结构类型
	//首字母大写代表是public的
	type User struct {
		id   int
		name string
		age  int
	}

	//声明一个结构体变量，并初始化为零值,string 类型的零值为""空字符串
	var user User
	fmt.Println("空结构体：", user) //{0  0}

	//声明一个结构体，并用字面量初始化
	var user1 User = User{
		id:   1,
		name: "kilric",
		age:  23,
	}

	fmt.Println(user1)
	//输出具体字段信息
	fmt.Printf("%#v\n", user1)

	//不使用字段值，创建字面量
	lisa := User{2, "lisa", 20}

	fmt.Printf("%#v\n", lisa)

	//嵌入其他结构体类型，作为匿名字段
	//定义类型Admin
	type Admin struct {
		User
		level string //管理权限
	}

	//初始化
	root := Admin{
		User: User{
			id:   1000,
			name: "root",
			age:  20,
		},
		level: "SuperAdministrator",
	}
	fmt.Println("===================间接访问===================")
	//可以这样访问
	fmt.Println("id:", root.User.id, "name:", root.User.name, "age:", root.User.age)
	fmt.Println("===================直接访问===================")
	//也可以直接访问匿名字段的值，这就是嵌入的好处，间接的实现了继承
	fmt.Println("id:", root.id, "name:", root.name, "age:", root.age)

	fmt.Println("===================利用系统打印================")
	fmt.Println(root)
	fmt.Println("====================打印出具体信息=============")
	fmt.Printf("%#v", root)
}
