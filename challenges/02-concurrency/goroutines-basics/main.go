package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	say("hello")
	fmt.Println(runtime.NumCPU())
	go say("world")
	time.Sleep(100 * time.Millisecond)
	ProducerConsumer()
}
