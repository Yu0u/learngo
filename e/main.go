package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func(ch chan int) {
		for i := 1; i <= 100; i++ {
			ch <- i
			if i%2 == 1 {
				fmt.Println("1 goroutine ", i)
			}
		}
		wg.Done()
	}(ch)
	go func(ch chan int) {
		for i := 1; i <= 100; i++ {
			<-ch
			if i%2 == 0 {
				fmt.Println("2 goroutine ", i)
			}
		}
		wg.Done()
	}(ch)
	wg.Wait()
}
