package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup){
	defer wg.Done()

	fmt.Println("Worker", id, "done")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go worker(i, &wg)
	}

	wg.Wait()

	fmt.Println("All Worker are finished")
}