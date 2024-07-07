package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//start := time.Now()
	//wait()
	//withoutWait()
	//fmt.Println(time.Now().Sub(start).Seconds())

	ch := make(chan int)

	go task(ch, 5)

	result := <-ch
	fmt.Println(result) // Output: 6
}

func task(ch chan<- int, i int) {
	ch <- i + 1
	close(ch)
}

func wait() {
	var wg sync.WaitGroup
	var mu sync.RWMutex
	var result int

	wg.Add(1000)
	for i := 0; i < 1000; i++ {

		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)

			mu.Lock()
			result++
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(result)
}

func withoutWait() {
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Nanosecond)
		fmt.Println(i)
	}
}
