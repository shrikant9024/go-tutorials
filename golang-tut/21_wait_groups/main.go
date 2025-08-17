package main

import (
	"fmt"
	"sync"
)

func counter(count int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Doing Task", count)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go counter(i, &wg)
	}

	wg.Wait()
}
