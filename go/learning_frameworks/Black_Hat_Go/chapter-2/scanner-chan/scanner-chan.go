package main

import (
	"fmt"
	"net"
	"sort"
)

//Use 2 channels, one to send request to worker and the second to receive the values
//You need to send to the workers in a separate goroutine because the result-gathering
//loop needs to start before more than 100 items of work can continue !
//syncronization with the channels

//Worker read the channel and send processed values to another channel
func Worker(ports, results chan int) {
	for port := range ports {
		// fmt.Println(port)
		// results <- port
		address := fmt.Sprintf("scanme.nmap.org:%d", port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- port

	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openPorts []int

	for i := 0; i < cap(ports); i++ {
		go Worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port == 0 {
			continue
		}
		openPorts = append(openPorts, port)
	}

	close(ports)
	close(results)

	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}
}
