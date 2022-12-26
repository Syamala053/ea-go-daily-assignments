package main

import (
	"fmt"
	"runtime"
	"time"
)

func greet() {
	fmt.Println("From Greet func")
	for{}
}

func main() {
	fmt.Println("Hello!")
	runtime.StartTrace()
	go greet()
	go greet()
	time.Sleep(200 * time.Second)
}
