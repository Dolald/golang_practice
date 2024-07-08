package main

import "fmt"

func main() {
	// Пример использования функции removeDuplicates
	inputs := []string{"a", "b", "w", "c", "c", "d", "d"}
	inputStream := make(chan string, len(inputs))
	outputStream := make(chan string)

	go func() {
		defer close(inputStream)

		for _, input := range inputs {
			fmt.Println("Input: %s to inputStream", input)
			inputStream <- input
		}
	}()

	go func() {
		defer close(outputStream)

		var prev string

		for val := range inputStream {
			if val != prev {
				outputStream <- val
				fmt.Println("Input: %s to outputStream", val)
				prev = val
			}
		}
	}()

	for val := range outputStream {
		fmt.Print(val)
	}

}
