package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var count int64
	c := make(chan int64)
	var wg sync.WaitGroup

	// bees
	for i := 0; i < 5000; i++ {
		//Add方法向内部计数加上delta，delta可以是负数；
		//如果内部计数器变为0，Wait方法阻塞等待的所有线程都会释放，
		//如果计数器小于0，方法panic。注意Add加上正数的调用应在Wait之前，
		//否则Wait可能只会等待很少的线程。
		//一般来说本方法应在创建新的线程或者其他应等待的事件之前调用。
		wg.Add(1)
		go func(in chan int64) {
			defer wg.Done()
			time.Sleep(100)
			in <- 2
		}(c)
	}

	// hive
	go func() {
		for out := range c {
			count += out
		}
	}()

	wg.Wait()
	// bang! but why?
	fmt.Println(count)
}
