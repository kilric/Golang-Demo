package main

import (
	"fmt"

	"method/util"
)

func main() {
	user := util.User{}

	ok1 := user.SetId(1)
	ok2 := user.SetName("lisa")
	if ok1 && ok2 {
		fmt.Println(user)
	}
	//已经设置成功了，说明User类型的实例可以调用*User类型的方法
}
