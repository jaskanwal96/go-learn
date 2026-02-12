package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			fmt.Println("Sending now")
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func Select() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Recieving now")
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func ProducerConsumer() {
	numbers := make(chan int)

	// Producer: sends numbers
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("About to send:", i)
			numbers <- i // This blocks until received
			fmt.Println("Sent:", i)
		}
		close(numbers) // Close channel when done
	}()

	// Consumer: receives numbers
	for num := range numbers {
		fmt.Println("Got:", num)
	}
}
