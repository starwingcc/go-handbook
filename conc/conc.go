// Package conc 演示并发：goroutine、channel、select、Mutex、Once、atomic。
package conc

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"sync/atomic"
	"time"
)

func goroutineDemo() {
	// wg.Go（Go 1.25+）= wg.Add(1)+go func(){ wg.Done() }()
	var wg sync.WaitGroup
	for range 5 {
		wg.Go(func() { fmt.Println("goroutine:", rand.IntN(100)) })
	}
	wg.Wait()
}

func channelDemo() {
	// 带缓冲 channel
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch) // range 前需 close
	for v := range ch {
		fmt.Println("chan:", v)
	}

	// 方向 channel：chan<- 只发、<-chan 只收（API 约束）
	produce := func(out chan<- int) { out <- 42; close(out) }
	consume := func(in <-chan int) { fmt.Println("dir chan:", <-in) }
	c := make(chan int, 1)
	produce(c)
	consume(c)
}

func selectDemo() {
	// select 多路复用 + 超时（time.After）
	ch := make(chan int)
	go func() { time.Sleep(5 * time.Millisecond); ch <- 1 }()
	select {
	case v := <-ch:
		fmt.Println("select got:", v)
	case <-time.After(50 * time.Millisecond):
		fmt.Println("select timeout")
	}

	// nil channel 在 select 中阻塞：用于动态启停分支
	var nilCh chan int // nil
	select {
	case <-nilCh: // 永不就绪，此分支被禁用
	case <-time.After(1 * time.Millisecond):
		fmt.Println("nil-chan branch disabled")
	}
}

func syncDemo() {
	// Mutex 保护共享状态
	var mu sync.Mutex
	cnt := 0
	var wg sync.WaitGroup
	for range 100 {
		wg.Go(func() {
			mu.Lock()
			cnt++
			mu.Unlock()
		})
	}
	wg.Wait()
	fmt.Println("mutex count:", cnt)

	// sync.Once 恰好一次
	var once sync.Once
	for range 3 {
		once.Do(func() { fmt.Println("once: 仅打印一次") })
	}

	// atomic 无锁计数器
	var a atomic.Int64
	for range 50 {
		wg.Go(func() { a.Add(1) })
	}
	wg.Wait()
	fmt.Println("atomic:", a.Load())
}

func Demo() {
	goroutineDemo()
	channelDemo()
	selectDemo()
	syncDemo()
	fmt.Println("build-tag platform:", platformName()) // 来自 build-tag 文件
}
