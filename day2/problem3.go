package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
type BankAccount struct {
	fund int
}
func (customer *BankAccount) WithDraw(amount int, mu *sync.Mutex) {
	mu.Lock()
	if customer.fund >= amount {
		customer.fund -= amount
	}else{
		panic("Not, Sufficient Amount")
	}
	mu.Unlock()
}
func (customer *BankAccount) Deposit(amount int, mu *sync.Mutex) {
	mu.Lock()
	customer.fund += amount
	mu.Unlock()
}
func main() {
	var mu sync.Mutex
	customer := BankAccount{fund: 500}
	for i := 0; i < 500; i++ {
		if rand.Intn(2) == 1 {
			go customer.WithDraw(rand.Intn(100), &mu)
		}else{
			go customer.Deposit(rand.Intn(1000), &mu)
		}
	}

	time.Sleep(time.Second * 2)
	fmt.Println(customer.fund)
}
