package main

import "sync"
import "time"

func main() {
	var mu sync.Mutex

	var wg sync.WaitGroup

	foo := func() {
		mu.Lock()
		time.Sleep(5 * time.Second)
		mu.Unlock()
		//wg.Done()
	}

	go foo()
	go foo()

	wg.Add(2)
	wg.Wait()
}