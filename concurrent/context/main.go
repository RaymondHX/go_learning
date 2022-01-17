package main

import (
	"context"
	"fmt"
	"time"
)

var key string = "name"

func monitorwithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止了")
				return
			default:
				fmt.Println("gorountine监控中")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)
	time.Sleep(time.Second * 10)
	fmt.Println("通知监控程序停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func monitorMultiWithContext() {
	ctx, cacel := context.WithCancel(context.Background())
	go watch(ctx, "监控1")
	go watch(ctx, "监控2")
	go watch(ctx, "监控3")
	time.Sleep(time.Second * 10)
	cacel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func monitorMultiWithValueContext() {
	ctx, cacel := context.WithCancel(context.Background())
	valueCtx := context.WithValue(ctx, key, "监控1")
	go watchWithMetaValue(valueCtx)
	time.Sleep(time.Second * 10)
	cacel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

func watchWithMetaValue(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//取出值
			fmt.Println(ctx.Value(key), "监控退出，停止了...")
			return
		default:
			//取出值
			fmt.Println(ctx.Value(key), "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	monitorMultiWithValueContext()
}
