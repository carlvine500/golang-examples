package main

import (
	"sync"
	"testing"
)

func TestName(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 执行一些任务
	}()
	wg.Wait() // 等待所有任务完成
}
func Test2(t *testing.T) {
	// done chan 能异常处理
	done := make(chan struct{})
	go func() {
		select {//无default阻塞
		case <-done:
			// 执行清理工作并退出
		}
	}()
	close(done) // 通知 goroutine 退出
}