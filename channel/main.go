package main

import (
	"fmt"
)

func test(i interface{}){
    if v,ok := i.(int);ok {
        fmt.Println(v)
    }else{
        fmt.Println(i)
    }
    
    switch i.(type) {
    case int:
        fmt.Println("int")
    case float64:
        fmt.Println("float64")
    default:
        fmt.Println("其他")
    }
    
}

func main() {
	//channel 用于在两个goroutine之间发送数据
	// <- 通过这个符号向指定方向发送

	var c1 chan int // 存储int值的双向管道，初始值为nil
	fmt.Println("ci==nil?", c1 == nil)

	c1 = make(chan int, 4) //c1是长度为4的管道

	fmt.Println("ci==nil?", c1 == nil)

	//c2 := make(<- chan int,4)  //发送，一般放在参数上,带缓冲
	c3 := make(chan int) //接收 ，一般放在参数上

	go func() {
		for i := 0; i < 4; i++ {
			c3 <- i + 10*i
		}
		close(c3) //用完后要关闭，否则造成死锁
	}()
    
	//for value:=range c3 {
	//    fmt.Printf("%v ",value)
	//}

	//同上
loop:
	for {
		select {
		case x, ok := <-c3:
			{
				if ok {
					fmt.Print(x, "  ")
				} else {
					break loop
				}
			}
		}
	}
	
	//var x = 45
	test("\nString")

}
