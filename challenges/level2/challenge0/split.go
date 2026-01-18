package main

import (
	"fmt"
	"time"
)

func sumAsync(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println("About to send:", sum) // This WILL print
	c <- sum                           // This BLOCKS FOREVER (no receiver)
	fmt.Println("Sent successfully!")  // This NEVER prints!
}

func sumSync(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func Split() {
	s := make([]int, 10000000)
	for i := range s {
		s[i] = i % 100 // or some other pattern
	}

	start := time.Now()
	c := make(chan int)
	go sumAsync(s[:len(s)/2], c)
	go sumAsync(s[len(s)/2:], c)
	x, y := <-c, <-c            // receive from c
	time.Sleep(2 * time.Second) // Wait 2 seconds
	fmt.Println("Done", x, y, time.Since(start))
	start = time.Now()
	fmt.Println(sumSync(s[:]), time.Since(start))
}
