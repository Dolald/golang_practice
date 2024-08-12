package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var mp sync.Map

	for i := 0; i <= 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mp.Store(i, i)

		}()
	}

	wg.Wait()

	mp.Range(func(key, value any) bool {
		fmt.Printf("%d: %d\n", key, value)
		return true
	})
}
