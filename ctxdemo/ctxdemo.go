// Package ctxdemo 演示 context 的取消、超时、传值。
package ctxdemo

import (
	"context"
	"fmt"
	"time"
)

func Demo() {
	// WithTimeout：到点自动取消
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	<-ctx.Done()
	fmt.Println("timeout cause:", ctx.Err()) // context.DeadlineExceeded

	// WithCancel：手动取消
	cctx, cancel2 := context.WithCancel(context.Background())
	go func() {
		time.Sleep(2 * time.Millisecond)
		cancel2()
	}()
	<-cctx.Done()
	fmt.Println("cancel cause:", cctx.Err()) // context.Canceled

	// WithValue：沿调用链传值（仅传请求级元数据，勿传业务数据）
	type key struct{}
	vctx := context.WithValue(context.Background(), key{}, "req-42")
	if v, ok := vctx.Value(key{}).(string); ok {
		fmt.Println("ctx value:", v)
	}
}
