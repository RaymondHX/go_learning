package main

import "fmt"

// 无缓冲通道，通道不会保存任何值，必须要两边都准备好,如果通道没有值，那么从通道里取值的操作就要一直等待
func unbufferedChannel() {
	ch := make(chan int)
	go func() {
		var sum int
		for i := 0; i < 10; i++ {
			sum += i
		}
		ch <- sum
	}()
	fmt.Println(<-ch)
}

//利用两个通道，实现一个管道的效果
func pipeline() {
	one := make(chan int)
	two := make(chan int)
	go func() {
		one <- 100
	}()

	go func() {
		value := <-one
		two <- value
	}()

	fmt.Println(<-two)
}

//有缓冲通道，通道其实相当于一个队列，发送数据就是向队列尾部插入，接收就是从队列头部取出
func bufferedChannel() {
	respones := make(chan string, 3)
	go func() { respones <- "1" }()
	go func() { respones <- "2" }()
	go func() { respones <- "3" }()
	fmt.Println(<-respones)
}

func main() {
	bufferedChannel()
}
