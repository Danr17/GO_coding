//channels
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// This will block
	/*
		<-ch
		fmt.Println("Here")
	*/
	go func() {
		//Send the number of the channel
		ch <- 353
	}()

	//receive from the channel
	val := <-ch
	fmt.Printf("got %d\n", val)

	fmt.Println("--------------")

	//Send multiple
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("Sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 3; i++ {
		val := <-ch
		fmt.Printf("received %d\n", val)
	}

	fmt.Println("--------------")

	//close to signal we're done
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("Sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("received %d\n", i)
	}

}
