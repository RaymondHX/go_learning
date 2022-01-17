package main

import (
	"fmt"
	"sync"
)

//sync.WaitGroup的使用也非常简单，先是使用Add 方法设设置计算器为2，每一个goroutine的函数执行完之后，就调用Done方法减1。
//Wait方法的意思是如果计数器大于0，就会阻塞，所以main 函数会一直等待2个goroutine完成后，再结束。
func waitGruop() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("A")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("B")
	}()
	wg.Wait()
	fmt.Println("finish")
}

func main() {
	waitGruop()
}
