package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello Jarvis")
}
func main() {
	go sayHello()

	fmt.Println("Hello Stark")

	time.Sleep(1 * time.Second)
}