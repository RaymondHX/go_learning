package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	count int32
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func lock() {
	wg.Add(2)
	go intCountLock()
	go intCountLock()
	wg.Wait()
	fmt.Println(count)
}

func intCountLock() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mutex.Lock()
		value := count
		runtime.Gosched()
		value++
		count = value
		mutex.Unlock()
	}
}

func main() {
	lock()
}
