package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestName(t *testing.T) {
	done := make(chan struct{})
	tasks := make(chan int, 2)
	go func() {
		for i := range tasks {// for range chan要等关闭才结束
			logrus.Info(i)
			//fmt.Println(i)
		}
		done <- struct{}{}
	}()
	tasks <- 1
	tasks <- 1
	tasks <- 1
	close(tasks)
	<-done
	fmt.Println("all done")
}
