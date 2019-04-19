package main

import (
	"fmt"
	"time"
)

func main() {
	// routine
	go say("hello")
	say("world")
}

/**
goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。
*/
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
