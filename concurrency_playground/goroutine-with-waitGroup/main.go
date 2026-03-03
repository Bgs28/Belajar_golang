package main

import (
	"fmt"
	"sync"
	"time"
)

func processTask(id int, wg *sync.WaitGroup) {
	defer wg.Done() // kasih tahu bahwa task selesai

	fmt.Println("Task", id, "started")
	time.Sleep(2 * time.Second)
	fmt.Println("Task", id, "finished")
}

func main(){
	var wg sync.WaitGroup

	for i := 1; i <= 100; i++{
		wg.Add(1)
		go processTask(i, &wg)
	}
	wg.Wait() // tunggu selesai semua

	fmt.Println("All Task Dispatched")
}