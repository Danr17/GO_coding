package main

import (
	"fmt"
	"sync"
)

//Worker read the channel and print the values
func Worker(ports chan int, wg *sync.WaitGroup) {
	for port := range ports {
		fmt.Println(port)
		wg.Done()
	}
}

func main() {
	var wg sync.WaitGroup
	ports := make(chan int, 100)

	for i := 0; i < cap(ports); i++ {
		go Worker(ports, &wg)
	}

	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}

	wg.Wait()
	close(ports)
}
