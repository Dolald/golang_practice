package main

import (
	"fmt"
	"time"
)

type Order struct {
	Id     int
	Amount int
}

func main() {
	notification := make(chan Order)
	payment := make(chan Order)
	validation := make(chan Order)

	go notify(notification)
	go pay(payment, notification)
	go validate(validation, payment)

	go func() {
		defer close(validation)
		orders := []Order{
			{Id: 1, Amount: 200},
			{Id: 2, Amount: 200},
			{Id: 3, Amount: 0},
			{Id: 4, Amount: 50},
		}
		for _, v := range orders {
			validation <- v
		}
	}()

	time.Sleep(7 * time.Second)
}

func validate(validation chan Order, payment chan Order) {
	defer close(payment)
	for v := range validation {
		if v.Amount > 0 {
			fmt.Println("Validation order:", v.Id)
			payment <- v
		} else {
			fmt.Println("Rejected order:", v.Id)
		}
	}
}

func pay(payment chan Order, notification chan Order) {
	defer close(notification)
	for v := range payment {
		fmt.Println("The order is being processed")
		time.Sleep(1 * time.Second)
		notification <- v
	}
}

func notify(notification chan Order) {
	for v := range notification {
		fmt.Println("Sending notification for order:", v.Id)
		time.Sleep(1 * time.Second)
		fmt.Println("Completed processing order:", v.Id)
	}

	for v := range notification {
		fmt.Println(v)
	}

}
