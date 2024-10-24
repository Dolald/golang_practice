package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var prodWg sync.WaitGroup
	var consWg sync.WaitGroup

	prodWg.Add(1)
	go produce(ch, &prodWg)

	go func() {
		prodWg.Wait()
		close(ch)
		fmt.Println("Channel was closed")
	}()

	consWg.Add(1)
	go consume(ch, &consWg)
	consWg.Wait()

	fmt.Println("done")
}

func produce(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 20; i++ {
		ch <- i
		fmt.Printf("%d was sent\n", i)
	}
	fmt.Println("producer done")
}

func consume(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Printf("%d was get\n", v)
	}
	fmt.Println("consumer done")
}
