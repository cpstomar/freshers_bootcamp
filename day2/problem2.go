package main

import (
	"fmt"
	"math/rand"
	"sync"
)
func teacher(wg *sync.WaitGroup, sum *int) {
	defer wg.Done()
	//time.Sleep(time.Second * time.Duration(rand.Intn(1)))
	*sum += rand.Intn(100)
}
func main() {
	var wg sync.WaitGroup
	var sum int = 0
	for i := 1; i <= 200; i++ {
		wg.Add(1)
		go teacher(&wg, &sum)
	}
	wg.Wait()
	fmt.Println((sum / 200))
}
