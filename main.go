package main

import (
	"fmt"
)

func main() {
	tick1 := make(chan int, 2)
	tick2 := make(chan int, 2)

	tick1 <- 2
	tick2 <- 6

	select {
	case <-tick1:
		fmt.Println("1 канал")
	case <-tick2:
		fmt.Println("2 канал")
	// Блок default выполнится раньше блока case - 1 секунда слишком много для Go
	default:
		fmt.Println("Действие по умолчанию")
	}
}
