package main

import (
	"fmt"
	"sync"
	"time"
)

const numWorkers = 3
const numTasks = 5

func main() {
	var wg sync.WaitGroup
	result := make(chan int, numTasks)
	tasksChan := make(chan int, numTasks)

	for idWorker := 1; idWorker <= numWorkers; idWorker++ {
		wg.Add(1)
		go func() {
			defer wg.Done() // Этот вызов должен быть в конце горутины
			for taskId := range tasksChan {
				fmt.Printf("Рабочий %d начал выполнение задачи %d\n", idWorker, taskId)
				time.Sleep(2 * time.Second) // имитация выполнения задачи
				fmt.Printf("Рабочий %d завершил выполнение задачи %d\n", idWorker, taskId)
				result <- taskId * 2
			}
		}() // Передаем idWorker как аргумент в анонимную функцию
	}

	go func() {
		defer close(tasksChan)
		for taskId := 1; taskId <= numTasks; taskId++ {
			tasksChan <- taskId
		}
	}()

	wg.Wait()
	close(result)

	for res := range result {
		fmt.Println("Результат работы:", res)
	}
}
