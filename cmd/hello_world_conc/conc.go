package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	for i := 0; i < 50000; i++ {
		// go starts a goroutine
		go printHelloWorld(i, ch)
	}

	for {
		msg := <-ch
		fmt.Println(msg)
	}
}

func printHelloWorld(i int, ch chan string) {
	for {
		ch <- fmt.Sprintf("Hello world from goroutine %d!\n", i)
	}
}
