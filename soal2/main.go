package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	data1 := []string{"bisa1", "bisa2", "bisa3"}
	data2 := []string{"coba1", "coba2", "coba3"}

	for i := 0; i < 1; i++ {
		wg.Add(2)
		go func() {
			mu.Lock()
			defer mu.Unlock()
			for key, _ := range data1 {
				fmt.Println(data1, key)
			}
			wg.Done()
		}()

		go func() {
			mu.Lock()
			defer mu.Unlock()
			for key, _ := range data2 {
				fmt.Println(data2, key)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
